package implementations

import (
	"context"

	"github.com/otaxhu/proyecto-go-hex-architecture/domain/models"
	"github.com/otaxhu/proyecto-go-hex-architecture/domain/repository/ordering"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoProductsRepository struct {
	db *mongo.Database
}

func NewMongoProductsRepository(db *mongo.Database) *mongoProductsRepository {
	return &mongoProductsRepository{db: db}
}

func (mongoRepo *mongoProductsRepository) GetProducts(ctx context.Context, limit, offset uint64, orderBy ordering.OrderBy) ([]models.Product, error) {
	if limit == 0 {
		return nil, ErrLimitIsZero
	}
	var sort bson.D
	switch orderBy {
	case ordering.OrderByIdAsc:
		sort = bson.D{{"_id", 1}}
	case ordering.OrderByIdDesc:
		sort = bson.D{{"_id", -1}}
	default:
		return nil, ErrOrderByIsNil
	}
	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset)).SetSort(sort)
	filter := bson.D{}
	cursor, err := mongoRepo.db.Collection("products").Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	products := []models.Product{}
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}
