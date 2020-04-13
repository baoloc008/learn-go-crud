package loggerfx

import (
	"github.com/spf13/viper"
	"learn-go-crud/pkg/logger"
)

func initLogger() {
	logger.InitLogger(viper.GetBool("debugging.logger"))
}
