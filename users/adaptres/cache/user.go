package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/VanillaFox/system_architecture_lab/users/models"
	"github.com/redis/go-redis/v9"
)

const cacheExpirationTime = 1 * time.Minute

func (c *Cache) GetUser(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}

	resultChan := make(chan *models.User)
	errChan := make(chan error)

	go func() {
		valBytes, err := c.rdb.Get(ctx, username).Bytes()
		if err == nil {
			err = json.Unmarshal(valBytes, user)
			if err == nil {
				resultChan <- user
				return
			}
		} else if err != redis.Nil {
			errChan <- err
			return
		}

		resultChan <- nil
	}()

	select {
	case user := <-resultChan:
		if user != nil {
			return user, nil
		}
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	user, err := c.repository.GetByUsername(ctx, username)

	if err != nil {
		return nil, err
	}

	go func() {
		userBytes, err := json.Marshal(user)
		if err != nil {
			return
		}

		c.rdb.Set(ctx, username, userBytes, cacheExpirationTime)

	}()

	return user, nil
}

func (c *Cache) FirstSetUser(ctx context.Context, user *models.User) error {
	err := c.repository.CreateUser(ctx, user)

	if err != nil {
		return err
	}

	userWithoutPassword := &models.User{FullName: user.FullName, Username: user.Username}

	userBytes, err := json.Marshal(userWithoutPassword)
	if err != nil {
		return err
	}

	c.mu.Lock()
	err = c.rdb.Set(ctx, userWithoutPassword.Username, userBytes, cacheExpirationTime).Err()
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

	userWithoutPassword := &models.User{FullName: user.FullName, Username: user.Username}

	userBytes, err := json.Marshal(userWithoutPassword)
	if err != nil {
		return nil, err
	}

	c.mu.Lock()
	err = c.rdb.Set(ctx, userWithoutPassword.Username, userBytes, cacheExpirationTime).Err()
	c.mu.Unlock()

	if err != nil {
		return nil, err
	}

	return user, err
}
