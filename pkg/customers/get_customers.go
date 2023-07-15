package customers

import (
	"net/http"

	"car_repair_api_go/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetCustomers(ctx *gin.Context) {
    var customers []models.Customer

    if result := h.DB.Find(&customers); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &customers)
}
