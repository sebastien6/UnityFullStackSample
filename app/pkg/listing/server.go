package listing

import (
	"context"
	"time"

	"github.com/sebastien6/UnityFullStackSample/app/pkg/listing/protocol"
	"github.com/sebastien6/UnityFullStackSample/app/pkg/listing/repo/mongo"
	"github.com/sebastien6/UnityFullStackSample/app/pkg/listing/service"
	"github.com/sebastien6/UnityFullStackSample/app/pkg/logger"
)

type config struct {
	grpcPort              string
	mongoConnectionString string
	mongoDatabase         string
	mongoCollection       string
	LogLevel              int
}

type MongoOptions struct {
	URL        string
	Database   string
	Collection string
}

// Options is a set of options to be passed to Run
type Options struct {
	// Port is the address port to listen
	Port string

	// Mongo defines MongoDB server connection parameters
	Mongo MongoOptions
}

const (
	timeout int = 10
)

func Run(ctx context.Context, opts Options) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()

	// Set a new MongoDB repository
	repo, err := mongo.NewMongoRepo(opts.Mongo.URL, opts.Mongo.Database, opts.Mongo.Collection)
	if err != nil {
		logger.Log.Sugar().Fatalf("Failed to serve: %s", err)
	}

	svc := service.NewListingService(repo)
	return protocol.RunGrpcServer(ctx, svc, opts.Port)
	// return protocol.RunHttpServer(ctx, cfg.httpPort, cfg.swaggeruiDir, cfg.grpcServerEndpointListing)
}
