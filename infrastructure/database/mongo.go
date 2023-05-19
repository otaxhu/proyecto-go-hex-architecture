package database

import (
	"context"
	"fmt"

	"github.com/otaxhu/proyecto-go-hex-architecture/settings"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection(ctx context.Context, settings *settings.Database) (*mongo.Database, error) {
	uri := fmt.Sprintf("%s://%s:%d", settings.Driver, settings.Host, settings.Port)
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.NewMongoConnection()")
	}
	return client.Database(settings.Name), nil
}
