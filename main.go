package main

import (
	"context"

	"github.com/otaxhu/proyecto-go-hex-architecture/infrastructure/database"
	"github.com/otaxhu/proyecto-go-hex-architecture/settings"
	"github.com/pkg/errors"
)

func main() {
	ctx := context.Background()
	settingsDb := settings.NewDatabase()
	db := database.NewSqlConnection(ctx, settingsDb)
	if err := db.PingContext(ctx); err != nil {
		panic(errors.Wrap(err, "main.go"))
	}
}
