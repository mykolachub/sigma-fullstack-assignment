package main

import (
	"sigma-test/config"
	"sigma-test/internal/app"
)

func main() {
	env := config.ConfigEnv()

	app.Run(env)
}
