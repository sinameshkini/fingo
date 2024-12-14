package core

import (
	"github.com/go-playground/validator/v10"
	"github.com/sinameshkini/fingo/internal/config"
	"github.com/sinameshkini/fingo/internal/repository"
	"github.com/sinameshkini/microkit/pkg/clients/cache"
	"log"
)

type Core struct {
	env      *config.Env
	repo     repository.Repository
	lock     *cache.Locker
	validate *validator.Validate
}

func New(env *config.Env, repo repository.Repository, c cache.Cache, v *validator.Validate) *Core {
	locker, err := cache.NewLocker(c)
	if err != nil {
		log.Fatal(err)
	}
	return &Core{
		env:      env,
		repo:     repo,
		lock:     locker,
		validate: v,
	}
}
