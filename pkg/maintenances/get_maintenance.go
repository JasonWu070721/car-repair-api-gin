package maintenances

import (
	"fmt"
	"net/http"

	"car_repair_api_go/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetMaintenance(ctx *gin.Context) {
    id := ctx.Param("id")
    fmt.Println("id", id)

    var car models.Car

    if result := h.DB.First(&car, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &car)
}
