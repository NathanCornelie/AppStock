package handler

import (
	"Mongo/internal/database"
	"Mongo/internal/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GetProducts(c *gin.Context) {
	cursor, err := database.Products.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch products"})
		return
	}

	var products []model.Product
	err = cursor.All(c, &products)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var body model.CreateProductRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request body"})
		return
	}
	res, err := database.Products.InsertOne(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create product "})
		return
	}
	product := model.Product{
		ID:       res.InsertedID.(primitive.ObjectID),
		Name:     body.Name,
		Category: body.Category,
		Price:    body.Price,
		Stock:    body.Stock,
	}
	c.JSON(http.StatusCreated, product)

}

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id provided"})
		return
	}
	result := database.Products.FindOne(c, primitive.M{"_id": _id})
	product := model.Product{}
	err = result.Decode(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to find product"})
		return
	}
	c.JSON(http.StatusOK, product)
}
func UpdateProductStockById(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id provided"})
		return
	}
	var body struct {
		Stock int `json:"stock" binding:"required"`
	}
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request body"})
		return
	}
	_, err = database.Products.UpdateOne(c, bson.M{"_id": _id}, bson.M{"$set": bson.M{"stock": body.Stock}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update product stock"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"succes": "product stock updated"})
}
func UpdateProductPriceById(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id provided"})
		return
	}
	var body struct {
		Price float32 `json:"price" binding:"required"`
	}
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request body"})
		return
	}
	_, err = database.Products.UpdateOne(c, bson.M{"_id": _id}, bson.M{"$set": bson.M{"price": body.Price}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update product price"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "product price updated"})
}
func DeleteProductById(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id provided"})
		return
	}

	res, err := database.Products.DeleteOne(c, bson.M{"_id": _id})
	if res.DeletedCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "product deleted"})
}
