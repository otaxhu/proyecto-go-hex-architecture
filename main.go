package main

import (
	"context"
	"fmt"

	"github.com/otaxhu/proyecto-go-hex-architecture/domain/repository"
	"github.com/otaxhu/proyecto-go-hex-architecture/domain/repository/ordering"
	"github.com/otaxhu/proyecto-go-hex-architecture/settings"
)

func main() {
	ctx := context.Background()
	settingsDatabase, err := settings.NewDatabase()
	if err != nil {
		panic(err)
	}
	repo, err := repository.NewProductsRepository(ctx, settingsDatabase)
	if err != nil {
		panic(err)
	}
	products, err := repo.GetProducts(ctx, 1, 1, ordering.OrderByIdAsc)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("%#v", p)
	}
}
