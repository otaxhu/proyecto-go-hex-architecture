package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/otaxhu/proyecto-go-hex-architecture/settings"
	"github.com/pkg/errors"
)

func NewSqlConnection(ctx context.Context, settings *settings.Database) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		settings.User,
		settings.Password,
		settings.Host,
		settings.Port,
		settings.Name,
	)
	db, err := sql.Open(settings.Driver, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "infrastructure.database.NewSqlConnection()")
	}
	return db, nil
}
