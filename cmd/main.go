package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Nezent/mig-test/internal/application/service"
	"github.com/Nezent/mig-test/internal/infrastructure/database"
	"github.com/Nezent/mig-test/internal/infrastructure/repository"
	"github.com/Nezent/mig-test/internal/interface/handler"
	"github.com/Nezent/mig-test/internal/interface/routes"
	"github.com/Nezent/mig-test/pkg/router"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		router.Module,
		routes.Module,
		database.Module,
		handler.Module,
		service.Module,
		repository.Module,
		fx.Invoke(func(
			router *chi.Mux,
			routes *routes.APIV1Routes,
			lc fx.Lifecycle,
		) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					// Register routes
					routes.Register()
					log.Printf("Server started on port %v", 8080)
					go func() {
						if err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), router); err != nil {
							log.Fatalf("failed to start server: %v", err)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return nil
				},
			})
		}),
	)
	app.Run()
}
