package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	apiv1 "github.com/sebastien6/UnityFullStackSample/app/pkg/api/v1"
	"github.com/sebastien6/UnityFullStackSample/app/pkg/listing/service"
	"github.com/sebastien6/UnityFullStackSample/app/pkg/logger"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mongoRepo struct {
	client     *mongo.Client
	database   string
	collection *mongo.Collection
	timeout    time.Duration
}

type GameItem struct {
	Id                  primitive.ObjectID `bson:"_id,omitempty"`
	Category            string             `bson:"category,omitempty"`
	Title               string             `bson:"title,omitempty"`
	Subtitle            string             `bson:"subtitle,omitempty"`
	Description         string             `bson:"description,omitempty"`
	Images              []*apiv1.Images    `bson:"images,omitempty"`
	Type                int32              `bson:"type,omitempty"`
	Tags                []string           `bson:"tags,omitempty"`
	Author              string             `bson:"author,omitempty"`
	ReplayBundleUrlJson string             `bson:"replayBundleUrlJson,omitempty"`
	Duration            float64            `bson:"duration,omitempty"`
	IsDownloadable      bool               `bson:"isDownloadable,omitempty"`
	IsStreamable        bool               `bson:"isStreamable,omitempty"`
	Version             string             `bson:"version,omitempty"`
}

func newMongoClient(mongoURL string, mongoTimeout time.Duration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	} else {
		log.Println("Successfully connected to mongodb")
	}

	return client, nil
}

func NewMongoRepo(mongoURL, mongoDB, collection string) (service.ServiceResository, error) {
	repo := &mongoRepo{
		timeout:  time.Duration(10) * time.Second,
		database: mongoDB,
	}
	client, err := newMongoClient(mongoURL, repo.timeout)
	if err != nil {
		return nil, errors.Wrap(err, "database.newMongoRepo")
	}
	repo.client = client
	repo.collection = repo.client.Database(repo.database).Collection(collection)

	return repo, nil
}

func (r *mongoRepo) FindAll(ctx context.Context) (*apiv1.Games, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("empty cursor creation failure: %v", err),
		)
	}

	defer cursor.Close(ctx)
	response := apiv1.Games{}
	counter := 0

	for cursor.Next(ctx) {
		data := &GameItem{}
		err := cursor.Decode(data)
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		logger.Log.Sugar().Infof("fetching data id: %s", data.Id.Hex())

		response.Listings = append(response.Listings, &apiv1.Game{
			Id:                  data.Id.Hex(),
			Category:            data.Category,
			Title:               data.Title,
			Subtitle:            data.Subtitle,
			Description:         data.Description,
			Images:              data.Images,
			Type:                data.Type,
			Tags:                data.Tags,
			Author:              data.Author,
			ReplayBundleUrlJson: data.ReplayBundleUrlJson,
			Duration:            data.Duration,
			IsDownloadable:      data.IsDownloadable,
			IsStreamable:        data.IsStreamable,
			Version:             data.Version,
		})

		counter += 1
	}
	logger.Log.Sugar().Infof("%d items fetched from database", counter)
	return &response, nil

}

func (r *mongoRepo) Store(ctx context.Context, game *apiv1.Game) (*apiv1.Game, error) {
	result, err := r.collection.InsertOne(
		ctx,
		bson.M{
			"Category":            game.GetCategory(),
			"Title":               game.GetTitle(),
			"Subtitle":            game.GetSubtitle(),
			"Description":         game.GetDescription(),
			"Images":              game.GetImages(),
			"Type":                game.GetType(),
			"Tags":                game.GetTags(),
			"Author":              game.GetAuthor(),
			"ReplayBundleUrlJson": game.GetReplayBundleUrlJson(),
			"Duration":            game.GetDuration(),
			"IsDownloadable":      game.GetIsDownloadable(),
			"IsStreamable":        game.GetIsStreamable(),
			"Version":             game.GetVersion(),
		},
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	game.Id = result.InsertedID.(primitive.ObjectID).Hex()
	logger.Log.Sugar().Infof("new record successsfully added to database with id %s", game.Id)
	return game, nil
}
