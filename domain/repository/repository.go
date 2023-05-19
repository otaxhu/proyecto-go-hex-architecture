package repository

import (
	"context"

	"github.com/otaxhu/proyecto-go-hex-architecture/domain/models"
	"github.com/otaxhu/proyecto-go-hex-architecture/domain/repository/implementations"
	"github.com/otaxhu/proyecto-go-hex-architecture/domain/repository/ordering"
	"github.com/otaxhu/proyecto-go-hex-architecture/infrastructure/database"
	"github.com/otaxhu/proyecto-go-hex-architecture/settings"
	"github.com/pkg/errors"
)

type ProductsRepository interface {
	// If the limit param is 0 or the orderBy param is nil then it will return an error
	GetProducts(ctx context.Context, limit, offset uint64, orderBy ordering.OrderBy) ([]models.Product, error)
	InsertProducts(products ...models.Product) error
}

var (
	ErrRepositoryImplementationUndefined error = errors.New("repository implementation undefined")
)

// Factory of repositories, depending on the driver settings, it will take the corresponding case
func NewProductsRepository(ctx context.Context, settings *settings.Database) (ProductsRepository, error) {
	switch settings.Driver {
	case "mysql":
		db, err := database.NewSqlConnection(ctx, settings)
		if err != nil {
			return nil, errors.Wrap(err, "repository.NewProductsRepository() case \"mysql\"")
		}
		return implementations.NewMysqlProductsRepository(ctx, db), nil

	case "mongodb":
		db, err := database.NewMongoConnection(ctx, settings)
		if err != nil {
			return nil, errors.Wrap(err, "repository.NewProductsRepository() case \"mongo\"")
		}
		return implementations.NewMongoProductsRepository(db), nil

	default:
		return nil, ErrRepositoryImplementationUndefined
	}
}
