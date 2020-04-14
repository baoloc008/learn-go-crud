package apifx

import "go.uber.org/fx"

var Module = fx.Invoke(registerHandler)
