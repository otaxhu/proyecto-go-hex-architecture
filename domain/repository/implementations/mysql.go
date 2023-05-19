package implementations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/otaxhu/proyecto-go-hex-architecture/domain/models"
	"github.com/otaxhu/proyecto-go-hex-architecture/domain/repository/ordering"
)

type mysqlProductsRepository struct {
	db *sql.DB
}

const (
	qryGetProducts = "SELECT products.id, products.name, products.description, products.stock, products.price, products.provider_id, providers.name AS provider_name, providers.region AS provider_region " +
		"FROM products JOIN providers ON products.provider_id = providers.id ORDER BY %s LIMIT ? OFFSET ?"
)

func NewMysqlProductsRepository(ctx context.Context, db *sql.DB) *mysqlProductsRepository {
	return &mysqlProductsRepository{db: db}
}

func (mysqlRepo mysqlProductsRepository) GetProducts(ctx context.Context, limit, offset uint64, orderBy ordering.OrderBy) ([]models.Product, error) {
	if limit == 0 {
		return nil, ErrLimitIsZero
	}
	tx, err := mysqlRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	var qryFmtd string
	switch orderBy {
	case ordering.OrderByIdAsc:
		qryFmtd = fmt.Sprintf(qryGetProducts, "id ASC")
	case ordering.OrderByIdDesc:
		qryFmtd = fmt.Sprintf(qryGetProducts, "id DESC")
	default:
		return nil, ErrOrderByIsNil
	}
	rows, err := tx.QueryContext(ctx, qryFmtd, limit, offset)
	if err != nil {
		return nil, err
	}
	products := []models.Product{}
	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Stock,
			&product.Price,
			&product.Provider.Id,
			&product.Provider.Name,
			&product.Provider.Region,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
