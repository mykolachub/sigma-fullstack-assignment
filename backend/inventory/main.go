package main

import (
	"sigma-inventory/cmd"
	"sigma-inventory/config"
)

func main() {
	env := config.ConfigEnv()

	cmd.Run(env)
}
