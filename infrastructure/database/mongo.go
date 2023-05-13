package database

import (
	"context"
	"fmt"

	"github.com/otaxhu/proyecto-go-hex-architecture/settings"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection(ctx context.Context, settings *settings.Database) *mongo.Database {
	uri := fmt.Sprintf("")
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(errors.Wrap(err, "infrastructure.database.NewMongoConnection()"))
	}
	return client.Database(settings.Name)
}
