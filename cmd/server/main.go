package main

import (
	"github.com/yvesromeiro/goapifc/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
