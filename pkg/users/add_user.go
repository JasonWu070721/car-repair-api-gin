package users

import (
	"fmt"
	"net/http"

	"car_repair_api_go/pkg/common/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AddUserRequestBody struct {
    UserName       string `json:"username"`
    Password       string `json:"password"`
    Email       string `json:"email"`
}

func (h handler) AddUser(ctx *gin.Context) {
    body := AddUserRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var user models.User

	pass := []byte(body.Password)

    // Hashing the password
    hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(hash))

    user.UserName = body.UserName
    user.Password = string(hash)
    user.Email = body.Email

    if result := h.DB.Create(&user); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusCreated, &user)
}
