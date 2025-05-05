package main

import (
	"backend/config"
	"backend/routes"
	"log"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}

	r := routes.SetupRouter()
	port := config.Env.Port
	r.Run(":" + port)
}