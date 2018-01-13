package main

import (
	"fmt"
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	R "github.com/marlonpeterc/product-microservice/routes"
)

func main() {

	fmt.Print("Starting product-microservice...")

	app := iris.New()
	app.Use(recover.New())

	api := app.Party("/api")
	{
		api.Get("/ping", R.Ping)
		api.Get("/products", R.GetProducts)
		api.Get("/products/:id", R.GetProduct)
	}

	app.Run(iris.Addr(os.Getenv("PORT")))

}
