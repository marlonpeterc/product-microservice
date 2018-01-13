package routes

import (
	"fmt"
	"log"
	"database/sql"

	"github.com/kataras/iris"
	p "github.com/marlonpeterc/product-microservice/product"
	cfg "github.com/marlonpeterc/product-microservice/config"
)

func Ping(ctx iris.Context) {
	json(ctx, "Pong")
}

func GetProducts(ctx iris.Context) {

	db := cfg.DB()

	var (
		id               string
		productName      string
		productShortDesc string
		productDesc      string
	)
	products := []p.Product{}

	rows, err := db.Query("SELECT id, product_name, product_short_desc, product_desc FROM products ORDER BY product_name")
	cfg.Err(err)

	for rows.Next() {
		rows.Scan(&id, &productName, &productShortDesc, &productDesc)
		product := p.Product{
			Id:        id,
			Name:      productName,
			ShortDesc: productShortDesc,
			Desc:      productDesc,
		}
		fmt.Println(product)
		products = append(products, product)
	}

	json(ctx, products)
}

func GetProduct(ctx iris.Context) {

	productId := ctx.Params().Get("id")
	query := "SELECT id, product_name, product_short_desc, product_desc FROM products WHERE id=$1"

	db := cfg.DB()

	var (
		id               string
		productName      string
		productShortDesc string
		productDesc      string
	)

	err := db.QueryRow(query, productId).Scan(&id, &productName, &productShortDesc, &productDesc)

	switch {
	case err == sql.ErrNoRows:
		json(ctx, "No Product with ID "+productId)
	case err != nil:
		log.Fatal(err)
	default:
		product := p.Product{
			Id:        id,
			Name:      productName,
			ShortDesc: productShortDesc,
			Desc:      productDesc,
		}
		json(ctx, product)
	}

}
