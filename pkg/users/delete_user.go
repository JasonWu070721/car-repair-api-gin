package users

import (
	"net/http"

	"car_repair_api_go/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteUser(ctx *gin.Context) {
    id := ctx.Param("id")

    var user models.User

    if result := h.DB.First(&user, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    h.DB.Delete(&user)

    ctx.Status(http.StatusNoContent)
}
