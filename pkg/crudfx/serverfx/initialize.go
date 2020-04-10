package serverfx

import "go.uber.org/fx"

var Module = fx.Invoke(initServer)
