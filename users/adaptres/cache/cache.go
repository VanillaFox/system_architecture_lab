package cache

import (
	"sync"

	"github.com/VanillaFox/system_architecture_lab/users/adaptres/postgres"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	rdb        *redis.Client
	repository *postgres.Repository
	mu         *sync.RWMutex
}

func NewCache(rdb *redis.Client, repository *postgres.Repository) *Cache {
	return &Cache{
		rdb:        rdb,
		repository: repository,
		mu:         &sync.RWMutex{},
	}
}
