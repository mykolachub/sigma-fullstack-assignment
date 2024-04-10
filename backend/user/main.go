package main

import (
	"sigma-user/config"
	"sigma-user/internal/app"
)

func main() {
	env := config.ConfigEnv()

	app.Run(env)
}
