package handler

import (
	"Mongo/internal/database"
	"Mongo/internal/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

func GetDocuments(c *gin.Context) {
	cursor, err := database.Documents.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch documents"})
		return
	}

	var documents []model.Documents
	err = cursor.All(c, &documents)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch documents"})
		return
	}
	c.JSON(http.StatusOK, documents)
}

func GetDocumentById(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id provided"})
		return
	}
	result := database.Documents.FindOne(c, primitive.M{"_id": _id})
	document := model.Documents{}
	err = result.Decode(&document)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to find document"})
		return
	}
	c.JSON(http.StatusOK, document)
}
func GetDocumentByUserId(c *gin.Context) {
	userId := c.Param("userId")
	_id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id provided"})
		return
	}
	result, err := database.Documents.Find(c, bson.M{"userId": _id}, options.Find().SetProjection(bson.M{"_id": 1, "userId": 1}))
	var documents []model.Documents
	err = result.All(c, &documents)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to find document for this user"})
		return
	}
	c.JSON(http.StatusOK, documents)
}

func CreateDocument(c *gin.Context) {
	var body model.CreateDocumentsRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request body"})
		return
	}
	res, err := database.Documents.InsertOne(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "enable to create document"})
		return
	}

	document := model.Documents{
		ID:       res.InsertedID.(primitive.ObjectID),
		Date:     primitive.NewDateTimeFromTime(time.Now()),
		ClientId: body.ClientId,
		UserId:   body.UserId,
		Type:     body.Type,
	}
	c.JSON(http.StatusCreated, document)

}
func DeleteDocumentById(c *gin.Context) {
	// TODO : implement the methods
}
