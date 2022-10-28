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

	// get user from user_id
	var user models.User
	db.Debug().Where("id = ?", Comment.UserId).First(&user)

	// get photo from photo_id
	var Photo = models.Photo{Id: Comment.PhotoId}
	db.Debug().First(&Photo)
	// c.JSON(http.StatusCreated, Photo)
	// return

	returnData := map[string]interface{}{
		"id":         Comment.Id,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    Comment.UserId,
		"updated_at": Comment.UpdatedAt,
		"created_at": Comment.CreatedAt,
		"user": struct {
			Id       uint   `json:"id"`
			Email    string `json:"email"`
			Username string `json:"username"`
		}{
			Id:       user.Id,
			Email:    user.Email,
			Username: user.Username,
		},
		"photo": struct {
			Id       uint   `json:"id"`
			Title    string `json:"title"`
			Caption  string `json:"caption"`
			PhotoUrl string `json:"photo_url"`
			UserId   uint   `json:"user_id"`
		}{
			Id:       Photo.Id,
			Title:    Photo.Title,
			Caption:  Photo.Caption,
			PhotoUrl: Photo.PhotoUrl,
			UserId:   Photo.UserId,
		},
	}

	c.JSON(http.StatusCreated, returnData)
}

func UpdateComment(c *gin.Context) {
	commentID := c.Param("commentId")
	db := models.DB

	Comment := models.Comment{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := db.Debug().Where("id = ?", commentID).First(&Comment).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to find comment!",
			"error":   err.Error(),
		})
		return
	}

	if Comment.UserId != userID {
		c.JSON(401, gin.H{
			"message": "Unauthorized!",
		})
		return
	}

	contentType := helpers.GetContentType(c)

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err = db.Debug().Save(&Comment).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to update comment!",
			"error":   err.Error(),
		})
		return
	}

	var dataPhoto models.Photo
	db.Debug().Where("id = ?", Comment.PhotoId).First(&dataPhoto)

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.Id,
		"title":      dataPhoto.Title,
		"caption":    dataPhoto.Caption,
		"photo_url":  dataPhoto.PhotoUrl,
		"user_id":    Comment.UserId,
		"updated_at": Comment.UpdatedAt,
	})
}

func DeleteComment(c *gin.Context) {
	commentID := c.Param("commentId")
	db := models.DB

	Comment := models.Comment{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := db.Debug().Where("id = ?", commentID).First(&Comment).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to find comment!",
			"error":   err.Error(),
		})
		return
	}

	if Comment.UserId != userID {
		c.JSON(401, gin.H{
			"message": "Unauthorized!",
		})
		return
	}

	err = db.Debug().Delete(&Comment).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to delete comment!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment deleted successfully!",
	})
}
