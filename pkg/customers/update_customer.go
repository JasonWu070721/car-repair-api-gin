package customers

import (
	"net/http"

	"car_repair_api_go/pkg/common/models"

	"github.com/gin-gonic/gin"
)


func (h handler) UpdateCustomer(ctx *gin.Context) {
    id := ctx.Param("id")
    body := models.UpdateCustomerRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var customer models.Customer

    if result := h.DB.First(&customer, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    customer.Name = body.Name
    customer.LicensePlate = body.LicensePlate
    customer.CarColor = body.CarColor
    customer.CarYear = body.CarYear

    h.DB.Save(&customer)

    ctx.JSON(http.StatusOK, &customer)
}
