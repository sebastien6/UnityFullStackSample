package gateway

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// Endpoint describes a gRPC endpoint
type Endpoint struct {
	Network, Addr string
}

// Options is a set of options to be passed to Run
type Options struct {
	// Addr is the address to listen
	Addr string

	// GRPCServer defines an endpoint of a gRPC service
	GRPCServer Endpoint

	// OpenAPIDir is a path to a directory from which the server
	// serves OpenAPI specs.
	OpenAPIDir string

	// Mux is a list of options to be passed to the gRPC-Gateway multiplexer
	Mux []gwruntime.ServeMuxOption
}

// Run starts a HTTP server and blocks while running if successful.
// The server will be shutdown when "ctx" is canceled.
func Run(ctx context.Context, opts Options) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	conn, err := dial(ctx, opts.GRPCServer.Network, opts.GRPCServer.Addr)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		if err := conn.Close(); err != nil {
			log.Fatalf("Failed to close a client connection to the gRPC server: %v", err)
		}
	}()

	mux := http.NewServeMux()

	mux.Handle("/openapiv2/", openAPIServer(opts.OpenAPIDir))
	mux.HandleFunc("/healthz", healthzServer(conn))

	gw, err := newGateway(ctx, conn, opts.Mux)
	if err != nil {
		return err
	}
	mux.Handle("/", gw)

	s := &http.Server{
		Addr: opts.Addr,
		Handler: middleware.Recoverer(middleware.Logger(
			middleware.RealIP(
				middleware.RequestID(mux),
			))),
	}
	go func() {
		<-ctx.Done()
		log.Println("Shutting down the http server")
		if err := s.Shutdown(ctx); err != nil {
			log.Fatalf("Failed to shutdown http server: %v", err)
		}
	}()

	log.Printf("Starting listening at %s", opts.Addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("failed to listen and serve: %v", err)
	}
	return nil
}
