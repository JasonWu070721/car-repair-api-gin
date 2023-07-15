package customers

import (
	"net/http"

	"car_repair_api_go/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteCustomer(ctx *gin.Context) {
    id := ctx.Param("id")

    var customer models.Customer

    if result := h.DB.First(&customer, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    h.DB.Delete(&customer)
    ctx.Status(http.StatusNoContent)
}
