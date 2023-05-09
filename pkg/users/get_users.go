package users

import (
	"net/http"

	"car_repair_api_go/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetUsers(ctx *gin.Context) {
    var users []models.User

    if result := h.DB.Find(&users); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &users)
}
