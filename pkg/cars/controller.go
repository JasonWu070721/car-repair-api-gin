package cars

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
    DB *gorm.DB
}


func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {
    h := &handler{
        DB: db,
    }

    routes := router.Group("/cars")
    routes.POST("/", h.AddCar)
    routes.GET("/:id", h.GetCar)
}
