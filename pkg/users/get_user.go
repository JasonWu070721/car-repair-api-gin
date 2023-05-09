package users

import (
	"fmt"
	"net/http"

	"car_repair_api_go/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetUser(ctx *gin.Context) {
    id := ctx.Param("id")
    fmt.Println("id", id)

    var user models.User

    if result := h.DB.First(&user, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &user)
}
