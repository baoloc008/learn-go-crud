package serverfx

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"net/http"
	"time"
)

func initServer(lc fx.Lifecycle, r *gin.Engine) {
	srv := http.Server{
		Addr:    ":9620",
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != http.ErrServerClosed {
					log.Fatalf("Server ListenAndServe Error: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Server is shutting down...")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := srv.Shutdown(ctx)
			if err != nil {
				log.Fatalf("Server Shutdown Error: %v", err)
			} else {
				log.Println("Server Shutdown Properly")
			}
			return err
		},
	})

}
