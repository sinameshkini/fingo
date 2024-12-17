package core

import (
	"github.com/go-playground/validator/v10"
	"github.com/sinameshkini/fingo/internal/config"
	"github.com/sinameshkini/fingo/internal/repository"
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/microkit/pkg/clients/cache"
	"log"
	"sync"
)

type Core struct {
	env      *config.Env
	repo     repository.Repository
	lock     *cache.Locker
	validate *validator.Validate
	queue    chan endpoint.TransactionRequest
	wg       sync.WaitGroup
}

func New(env *config.Env, repo repository.Repository, ca cache.Cache, v *validator.Validate) *Core {
	locker, err := cache.NewLocker(ca)
	if err != nil {
		log.Fatal(err)
	}

	c := &Core{
		env:      env,
		repo:     repo,
		lock:     locker,
		validate: v,
		queue:    make(chan endpoint.TransactionRequest, 100),
	}

	c.startWorker()

	return c
}
