package users

import (
	"car_repair_api_go/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheckUserInput struct {
	UserName string `json:"username" binding:"required"`
}

func (h handler) CheckUser(ctx *gin.Context) {
    // body := AddUserRequestBody{}

    checkUserInput := CheckUserInput{}

	if err := ctx.ShouldBindJSON(&checkUserInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    username := checkUserInput.UserName
    users := []models.User{}

	err := h.DB.Where("user_name = ?", username).Find(&users)

	if err != nil {
		return
	}
	
    // err := h.DB.Model(user).Where("user_name = ?", username).Take(&user).Error

	// if err != nil {
	// 	return
	// }


    ctx.JSON(http.StatusCreated, &users)
}
