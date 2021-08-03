package main

import (
	"context"
	"flag"
	"log"

	"github.com/sebastien6/UnityFullStackSample/app/pkg/gateway"
)

var (
	endpoint   = flag.String("endpoint", "localhost:9090", "endpoint of the gRPC service, default: \"localhost:9090\"")
	network    = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
	httpPort   = flag.String("http-port", "8080", "TCP port HTTP server will listen on, default: \"8080\"")
	openAPIDir = flag.String("openapi-dir", "/app/openapiv2", "path to the directory which contains OpenAPI definitions, default: \"/app/openapiv2\"")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	opts := gateway.Options{
		Addr: ":" + *httpPort,
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *endpoint,
		},
		OpenAPIDir: *openAPIDir,
	}

	if err := gateway.Run(ctx, opts); err != nil {
		log.Fatal(err)
	}
}
