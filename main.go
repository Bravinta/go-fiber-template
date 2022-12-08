package main

import (
	env "github.com/joho/godotenv"
	"go-fiber-template/internal"
	"go-fiber-template/internal/infrastructure/utils/logs"
	"go-fiber-template/internal/infrastructure/utils/vars"
)

func main() {
	err := env.Load()
	if err != nil {
		logs.Send(vars.TypeLogs.Error, "Could not load .env file")
	}

	internal.Run()
}
