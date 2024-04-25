package main

import (
	"sigma-order/cmd"
	"sigma-order/config"
)

func main() {
	env := config.ConfigEnv()

	cmd.Run(env)
}
