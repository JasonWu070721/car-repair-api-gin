package auth

import (
	"fmt"
	"net/http"

	"car_repair_api_go/pkg/common/models"
	"car_repair_api_go/pkg/utils/token"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

type LoginInput struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func (h handler)Login(c *gin.Context) {
	
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.UserName = input.UserName
	u.Password = input.Password

	token, err := h.LoginCheck(u.UserName, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token":token})

}

func (h handler) LoginCheck(username string, password string) (string, error) {

	var err error

	user := models.User{}

	err = h.DB.Model(user).Where("user_name = ?", username).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	fmt.Println("user.ID: ", user.ID)

	token, err := token.GenerateToken(user.ID)


	if err != nil {
		return "", err
	}

	return token, nil

}
