package dbfx

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"learn-go-crud/pkg/logger"

	// import sqlite dialects
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDatabase(lifecycle fx.Lifecycle) (*gorm.DB, error) {
	db, err := gorm.Open(viper.GetString("database.dialect"), viper.GetString("database.url"))
	if err != nil {
		logger.Errorw("cannot init database connection", "error", err)
		return nil, err
	}
	if viper.GetBool("debug.db") {
		db = db.Debug()
	}

	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			logger.Debug("Closing DB")
			return db.Close()
		},
	})

	return db, nil
}
