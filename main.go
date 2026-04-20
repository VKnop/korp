package main

import (
	"korp/controller"
	"korp/db"
	"korp/repository"
	"korp/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	service := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada de repository
	productRepository := repository.NewProductRepository(dbConnection)
	//Camada usecase
	productUseCase := usecase.NewProductUseCase(productRepository)
	//Camada de controllers
	productController := controller.NewProductController(productUseCase)

	service.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	service.GET("/products", productController.GetProducts)
	service.GET("/product/:id", productController.GetProductById)
	service.POST("/product", productController.CreateProduct)
	service.PUT("/product/:id", productController.EditProduct)

	service.Run(":8000")
}
