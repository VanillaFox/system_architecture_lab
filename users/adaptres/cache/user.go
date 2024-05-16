package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	"github.com/VanillaFox/system_architecture_lab/users/models"
	"github.com/redis/go-redis/v9"
)

const cacheExpirationTime = 1 * time.Minute

func (c *Cache) GetUser(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}

	c.mu.RLock()
	valBytes, err := c.rdb.Get(ctx, username).Bytes()
	c.mu.RUnlock()

	if err == redis.Nil {
		user, err = c.repository.GetByUsername(ctx, username)

		if err != nil {
			return nil, err
		}

		var userBytes bytes.Buffer

		if err := gob.NewEncoder(&userBytes).Encode(user); err != nil {
			return nil, err
		}

		c.mu.Lock()
		err = c.rdb.Set(ctx, username, userBytes.Bytes(), cacheExpirationTime).Err()
		c.mu.Unlock()

		if err != nil {
			return nil, err
		}

		return user, nil
	} else if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(valBytes)

	if err := gob.NewDecoder(reader).Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (c *Cache) FirstSetUser(ctx context.Context, user *models.User) error {
	err := c.repository.CreateUser(ctx, user)

	if err != nil {
		return err
	}

	var userBytes bytes.Buffer

	userWithoutPassword := &models.User{FullName: user.FullName, Username: user.Username}

	if err := gob.NewEncoder(&userBytes).Encode(userWithoutPassword); err != nil {
		return err
	}

	c.mu.Lock()
	err = c.rdb.Set(ctx, userWithoutPassword.Username, userBytes.Bytes(), cacheExpirationTime).Err()
	c.mu.Unlock()

	if err != nil {
		return err
	}

	return err
}

func (c *Cache) SetUser(ctx context.Context, username string, user *models.User) (*models.User, error) {
	user, err := c.repository.UpdateUser(ctx, username, user)

	if err != nil {
		return nil, err
	}

	c.mu.Lock()
	c.rdb.Del(ctx, username)
	c.mu.Unlock()

	var userBytes bytes.Buffer

	userWithoutPassword := &models.User{FullName: user.FullName, Username: user.Username}

	if err := gob.NewEncoder(&userBytes).Encode(userWithoutPassword); err != nil {
		return nil, err
	}

	c.mu.Lock()
	err = c.rdb.Set(ctx, userWithoutPassword.Username, userBytes.Bytes(), cacheExpirationTime).Err()
	c.mu.Unlock()

	if err != nil {
		return nil, err
	}

	return user, err
}
