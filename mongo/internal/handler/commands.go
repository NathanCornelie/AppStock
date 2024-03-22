package handler

import (
	"Mongo/internal/database"
	"Mongo/internal/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GetCommands(c *gin.Context) {
	cursor, err := database.Commands.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch commands"})
		return
	}
	var commands []model.Command
	err = cursor.All(c, &commands)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch commands"})
		return
	}
	c.JSON(http.StatusOK, commands)
}

func CreateCommand(c *gin.Context) {
	var body model.CreateCommandRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request body"})
		return
	}
	res, err := database.Commands.InsertOne(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create product "})
		return
	}

	command := model.Command{
		ID:         res.InsertedID.(primitive.ObjectID),
		DocumentId: body.DocumentId,
		ItemId:     body.ItemId,
		Quantity:   body.Quantity,
		Discount:   body.Discount,
	}
	c.JSON(http.StatusCreated, command)

}

func GetCommandById(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id provided"})
		return
	}
	result := database.Commands.FindOne(c, primitive.M{"_id": _id})
	command := model.Command{}
	err = result.Decode(&command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to find command"})
		return
	}
	c.JSON(http.StatusOK, command)
}

func DeleteCommandById(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id provided"})
		return
	}

	res, err := database.Commands.DeleteOne(c, bson.M{"_id": _id})
	if res.DeletedCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "command not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "command deleted"})
}
