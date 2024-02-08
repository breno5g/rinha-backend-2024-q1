package main

import "github.com/breno5g/rinha-backend-2024-q1/config"

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}

}
