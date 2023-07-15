package customers

import (
	"fmt"
	"net/http"

	"car_repair_api_go/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetCustomer(ctx *gin.Context) {
    id := ctx.Param("id")
    fmt.Println("id", id)

    var customer models.Customer

    if result := h.DB.First(&customer, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &customer)
}
