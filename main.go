package main

import (
	"fmt"
	cfg "github.com/marlonpeterc/product-microservice/config"
)

func main() {
	fmt.Print("Starting product-microservice...")

	db := cfg.DB()

	var (
		id               string
		productName      string
		productShortDesc string
		productDesc      string
	)
	products := []interface{}{}

	rows, err := db.Query("SELECT id, product_name, product_short_desc, product_desc FROM products ORDER BY product_name")
	cfg.Err(err)

	for rows.Next() {
		rows.Scan(&id, &productName, &productShortDesc, &productDesc)
		product := map[string]interface{}{
			"id":               id,
			"productName":      productName,
			"productShortDesc": productShortDesc,
			"productDesc":      productDesc,
		}
		fmt.Println(product)
		products = append(products, product)
	}
}
