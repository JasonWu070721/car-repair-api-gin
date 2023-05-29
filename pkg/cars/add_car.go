package cars

import (
	"net/http"

	"car_repair_api_go/pkg/common/models"

	"github.com/gin-gonic/gin"
)

type AddRequestBody struct {
    Year       string `json:"year"`
    Make       string `json:"make"`
    Type       string `json:"type"`
	BodyStyles       string `json:"body_styles"`
}

func (h handler) AddCar(ctx *gin.Context) {
    body := AddRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var car models.Car

    car.Year = body.Year
    car.Make = body.Make
    car.Type = body.Type
	car.BodyStyles = body.BodyStyles

    if result := h.DB.Create(&car); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusCreated, &car)
}
