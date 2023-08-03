package customers

import (
	"net/http"

	"car_repair_api_go/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) AddCustomer(ctx *gin.Context) {
    body := models.AddCustomerRequestBody{}
    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var customer models.Customer

    customer.Name = body.Name
    customer.LicensePlate = body.LicensePlate
    customer.CarColor = body.CarColor
    customer.CarYear = body.CarYear

    if result := h.DB.Create(&customer); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusCreated, &customer)
}
