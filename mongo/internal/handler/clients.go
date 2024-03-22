package handler

import (
	"Mongo/internal/database"
	"Mongo/internal/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GetClients(c *gin.Context) {
	cursor, err := database.Clients.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch clients"})
		return
	}
	var clients []model.Client
	err = cursor.All(c, &clients)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch clients"})
		return
	}
	c.JSON(http.StatusOK, clients)
}

func CreateClient(c *gin.Context) {
	var body model.CreateClientRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request body"})
		return
	}
	res, err := database.Clients.InsertOne(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create client "})
		return
	}
	client := model.Client{
		ID:        res.InsertedID.(primitive.ObjectID),
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Email:     body.Email,
	}
	c.JSON(http.StatusCreated, client)

}
func GetClientById(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id provided"})
		return
	}
	result := database.Clients.FindOne(c, primitive.M{"_id": _id})
	client := model.Client{}
	err = result.Decode(&client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to find client"})
		return
	}
	c.JSON(http.StatusOK, client)
}
func EditClientInfos(c *gin.Context) {

}

// TODO : edit client info function

func DeleteClientById(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id provided"})
		return
	}

	res, err := database.Clients.DeleteOne(c, bson.M{"_id": _id})
	if res.DeletedCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "client deleted"})
}
