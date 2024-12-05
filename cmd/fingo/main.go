package main

import (
	"github.com/sinameshkini/fingo/internal/config"
	"github.com/sinameshkini/fingo/service"
	"log"
)

func main() {
	if err := service.Run(config.DefaultConf); err != nil {
		log.Fatal(err)
	}
}
