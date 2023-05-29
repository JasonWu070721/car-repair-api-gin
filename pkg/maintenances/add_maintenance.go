package maintenances

import (
	"net/http"

	"car_repair_api_go/pkg/common/models"

	"github.com/gin-gonic/gin"
)

type AddRequestBody struct {
    UserID       string `json:"user_id"`
    CarID       string `json:"car_id"`
	AppointmentTime       string `json:"appointment_time"`
}

func (h handler) AddMaintenance(ctx *gin.Context) {
    body := AddRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var maintenance models.Maintenance

    maintenance.UserID = body.UserID
    maintenance.CarID = body.CarID
    maintenance.AppointmentTime = body.AppointmentTime

    if result := h.DB.Create(&maintenance); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusCreated, &maintenance)
}
