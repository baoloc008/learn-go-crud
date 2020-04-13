package loggerfx

import "go.uber.org/fx"

var Module = fx.Invoke(initLogger)
