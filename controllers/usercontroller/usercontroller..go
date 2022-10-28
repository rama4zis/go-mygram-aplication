package usercontroller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
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

	// turn user age to int
	User.Age = uint(User.Age)

	// Create User
	User = models.User{Email: User.Email, Username: User.Username, Password: User.Password, Age: User.Age}
	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create user!",
			"error":   err,
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

	token := helpers.GenerateToken(User.Id, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Update(c *gin.Context) {
	db := models.DB

	userData := c.MustGet("userData").((jwt.MapClaims))
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	userId := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.Id = userId
	// User = models.User{Email: User.Email, Username: User.Username}
	err := db.Debug().Model(&User).Where("id = ?", userId).Updates(User).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to update user!",
			"error":   err,
		})
		return
	}

	userUpdate := models.User{}
	db.Debug().Where("id = ?", userId).Take(&userUpdate)

	c.JSON(200, gin.H{
		"id":         userUpdate.Id,
		"email":      userUpdate.Email,
		"username":   userUpdate.Username,
		"age":        userUpdate.Age,
		"updated_at": userUpdate.UpdatedAt,
	})
}

func Delete(c *gin.Context) {
	var user models.User
	userData := c.MustGet("userData").((jwt.MapClaims))
	userId := uint(userData["id"].(float64))

	result := models.DB.Delete(&user, userId)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Failed to delete user!",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": "User deleted!",
	})

}
