package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	server := gin.Default()
	server.Use(cors.Default())

	dbconnection, err := db.ConnectDB()
	if err != nil{
		panic(err)
	}
	//camda repository
	productRepository := repository.NewProductRepository(dbconnection)
	//camada usercase
	productUserCase := usecase.NewProductUseCase(productRepository)
	//camada de controllers
	productController := controller.NewProductController(productUserCase)


	//teste de ping
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	//GET
	server.GET("/products", productController.GetProducts)
	
	//POST 
	server.POST("/products", productController.CreateProduct)

	
	//delete 
	server.DELETE("/products/:id", productController.DeleteProduct)

	//update
	server.PUT("/products/:id", productController.PutProduct)


	server.Run(":8000")
}
