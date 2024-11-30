package main

import (
	"github.com/sinameshkini/fingo/service"
	"log"
)

func main() {
	if err := service.Run(service.DefaultConf); err != nil {
		log.Fatal(err)
	}
}
