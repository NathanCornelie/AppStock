package internal

import (
	"Mongo/internal/handler"
	"github.com/gin-gonic/gin"
)

func AddProductRoutes(r *gin.Engine) {
	r.GET("/products", handler.GetProducts)
	r.POST("/products", handler.CreateProduct)
	r.GET("/products/:id", handler.GetProductById)
	r.PATCH("/products/:id/stock", handler.UpdateProductStockById)
	r.PATCH("/products/:id/price", handler.UpdateProductPriceById)
	r.DELETE("/products/:id", handler.DeleteProductById)
}

func AddClientsRoutes(r *gin.Engine) {
	r.GET("/clients", handler.GetClients)
	r.POST("/clients", handler.CreateClient)
	r.PATCH("/clients/:id", handler.EditClientInfos)
	r.DELETE("/clients/:id", handler.DeleteClientById)
}
