package main

import (
	"github.com/sinameshkini/fingo/service"
	"log"
)

func main() {
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
