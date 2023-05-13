package database

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/otaxhu/proyecto-go-hex-architecture/settings"
	"github.com/pkg/errors"
)

func NewSqlConnection(ctx context.Context, settings *settings.Database) *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		settings.User,
		settings.Password,
		settings.Host,
		settings.Port,
		settings.Name,
	)
	db, err := sqlx.ConnectContext(ctx, settings.Driver, dsn)
	if err != nil {
		panic(errors.Wrap(err, "infrastructure.database.NewConnection()"))
	}
	return db
}
