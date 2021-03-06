package main

import (
	"go.uber.org/fx"
	"learn-go-crud/pkg/crudfx/apifx"
	"learn-go-crud/pkg/crudfx/configfx"
	"learn-go-crud/pkg/crudfx/dbfx"
	"learn-go-crud/pkg/crudfx/ginfx"
	"learn-go-crud/pkg/crudfx/loggerfx"
	"learn-go-crud/pkg/crudfx/serverfx"
)

func main() {
	app := fx.New(
		configfx.InitConfig("crudapi", "crudapi", "configs"),
		loggerfx.Module,
		ginfx.Module,
		dbfx.Module,
		serverfx.Module,
		apifx.Module,
	)
	app.Run()
}
