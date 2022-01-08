package main

import (
	"github.com/adKoch/goacc/goapi/app"
	"github.com/adKoch/goacc/goapi/app/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
