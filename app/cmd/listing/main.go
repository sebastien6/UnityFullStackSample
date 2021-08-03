package main

import (
	// +build tools
	// _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	// _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	// _ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	// _ "google.golang.org/protobuf/cmd/protoc-gen-go"
	// _ "github.com/envoyproxy/protoc-gen-validate"

	"context"
	"flag"
	"log"

	"github.com/sebastien6/UnityFullStackSample/app/pkg/listing"
	"github.com/sebastien6/UnityFullStackSample/app/pkg/logger"
)

var (
	endpoint              = flag.String("endpoint", "9090", "TCP port to bind with GRPC server default 5001")
	mongoConnectionString = flag.String("mongo-connection-string", "mongodb://localhost:27017", "mongodb connection string, default \"mongodb://localhost:27017\"")
	mongoDatabase         = flag.String("mongo-database", "demo", "Database on which MongoDB client will connect to, default \"DEMO\"")
	mongoCollection       = flag.String("mongo-collection", "games", "Database on which MongoDB client will connect to, default \"DEMO\"")
	logLevel              = flag.Int("log-level", 0, "Global log level Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5), default 0")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	// initialize logger
	if err := logger.Init(*logLevel, ""); err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	opts := listing.Options{
		Port: ":" + *endpoint,
		Mongo: listing.MongoOptions{
			URL:        *mongoConnectionString,
			Database:   *mongoDatabase,
			Collection: *mongoCollection,
		},
	}

	if err := listing.Run(ctx, opts); err != nil {
		logger.Log.Sugar().Fatal(err)
	}
}
