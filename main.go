package main

import (
	"cicd-demo/cmd"
	"cicd-demo/handlers"
	"cicd-demo/repositories"
	"cicd-demo/services"
	"context"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/fx"
)

func NewHTTPServer(
	lc fx.Lifecycle,
	productHandler *handlers.ProductHandler,
) *http.Server {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: cmd.Router(productHandler),
	}
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				ln, err := net.Listen("tcp", srv.Addr)
				if err != nil {
					return err
				}
				fmt.Println("Starting HTTP server at ", srv.Addr)
				go srv.Serve(ln)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return srv.Shutdown(ctx)
			},
		},
	)
	return srv
}

func main() {
	fx.New(
		fx.Provide(
			NewHTTPServer,
			services.ProvideProductService,
			repositories.ProvideProductRepository,
			handlers.ProvideProductHandler,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
