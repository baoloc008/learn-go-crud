package serverfx

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"learn-go-crud/pkg/logger"
	"net/http"
	"time"
)

func initServer(lc fx.Lifecycle, r *gin.Engine) {
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("server.port")),
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("Server is starting...")
				if err := srv.ListenAndServe(); err != http.ErrServerClosed {
					logger.Fatalf("Server ListenAndServe Error: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Server is shutting down...")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := srv.Shutdown(ctx)
			if err != nil {
				logger.Fatalf("Server Shutdown Error: %v", err)
			} else {
				logger.Info("Server Shutdown Properly")
			}
			return err
		},
	})

}
