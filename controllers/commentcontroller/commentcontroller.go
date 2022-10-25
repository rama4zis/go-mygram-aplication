package commentcontroller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rama4zis/go-mygram-aplication/helpers"
	"github.com/rama4zis/go-mygram-aplication/models"
)

func CreateComment(c *gin.Context) {
	db := models.DB
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserId = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create comment!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.Id,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    Comment.UserId,
		"created_at": Comment.CreatedAt,
	})
}
