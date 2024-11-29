package core

import (
	"github.com/sinameshkini/fingo/internal/repository"
	"github.com/sinameshkini/fingo/pkg/clients/cache"
	"log"
)

type Core struct {
	repo repository.Repository
	lock *cache.Locker
}

func New(repo repository.Repository, c cache.Cache) *Core {
	locker, err := cache.NewLocker(c)
	if err != nil {
		log.Fatal(err)
	}
	return &Core{
		repo: repo,
		lock: locker,
	}
}
