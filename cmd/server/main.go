package main

import (
	"github.com/breno5g/rinha-backend-2024-q1/config"
	"github.com/breno5g/rinha-backend-2024-q1/internal/router"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}

	router.Init()
}
