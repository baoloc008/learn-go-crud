package main

import (
	"go.uber.org/fx"
	"learn-go-crud/pkg/crudfx/configfx"
	"learn-go-crud/pkg/crudfx/ginfx"
	"learn-go-crud/pkg/crudfx/serverfx"
)

func main() {
	app := fx.New(
		configfx.InitConfig("crudapi", "crudapi", "configs"),
		ginfx.Module,
		serverfx.Module,
	)
	app.Run()
}
