package internal

import (
	"Mongo/internal/handler"
	"Mongo/internal/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"os"
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
	r.GET("/clients/:id", handler.GetClientById)
	r.POST("/clients", handler.CreateClient)
	r.PATCH("/clients/:id", handler.EditClientInfos)
	r.DELETE("/clients/:id", handler.DeleteClientById)
}

func AddCommandsRoutes(r *gin.Engine) {
	r.GET("/commands", handler.GetCommands)
	r.GET("/commands/:id", handler.GetCommandById)
	r.POST("/commands", handler.CreateCommand)
	r.DELETE("/commands/:id", handler.DeleteCommandById)
}

func AddDocumentsRoutes(r *gin.Engine) {
	r.GET("/documents", handler.GetDocuments)
	r.GET("/documents/:id", handler.GetDocumentById)
	r.GET("/documents/users/:userId", handler.GetDocumentByUserId)
	r.GET("/documents/client/:clientId", handler.GetDocumentByClientId)
	r.POST("/documents", handler.CreateDocument)
	r.DELETE("/documents/:id", handler.DeleteDocumentById)
}

func DefineRouter() *gin.Engine {
	r := gin.New()
	// logs
	f, _ := os.Create("logs.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r.Use(gin.LoggerWithFormatter(logger.FormatLogs))
	//cors policies
	r.Use(cors.Default())

	AddProductRoutes(r)
	AddClientsRoutes(r)
	AddCommandsRoutes(r)
	AddDocumentsRoutes(r)
	return r
}
