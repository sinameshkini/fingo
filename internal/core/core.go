package core

import "github.com/sinameshkini/fingo/internal/repository"

type Core struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Core {
	return &Core{repo: repo}
}
