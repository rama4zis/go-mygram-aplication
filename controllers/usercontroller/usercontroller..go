package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rama4zis/go-mygram-aplication/helpers"
	"github.com/rama4zis/go-mygram-aplication/models"
)

func Index(c *gin.Context) {
	var users []models.User

	result := models.DB.Find(&users)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Failed to fetch users!",
		})
		return
	}

	c.JSON(200, gin.H{"data": users})
}

func Show(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	result := models.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Failed to fetch user!",
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

func Register(c *gin.Context) {
	db := models.DB
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create user!",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id":       User.Id,
		"username": User.Username,
	})
}

func Login(c *gin.Context) {
	db := models.DB
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	password := ""

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Email or password is wrong!",
			"message": "Failed to login!",
		})
	}

	comparePass := helpers.ComparePassword([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Email or password is wrong!",
			"message": "Failed to login!",
		})
		return
	}

	token := helpers.GenerateToken(User.Id, User.Username, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Update(c *gin.Context) {
	// get userId from authotization header
	userId := c.MustGet("userId").(float64)

	// get user from database
	var user models.User
	result := models.DB.First(&user, userId)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Failed to fetch user!",
		})
		return
	}

	// get user input
	contentType := helpers.GetContentType(c)
	if contentType == "application/json" {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	// update user
	result = models.DB.Save(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Failed to update user!",
		})
		return
	}

	c.JSON(200, gin.H{
		"id":         user.Id,
		"email":      user.Email,
		"username":   user.Username,
		"age":        user.Age,
		"updated_at": user.UpdatedAt,
	})
}
