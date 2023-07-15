package customers

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

    routes := router.Group("/customers")
    routes.POST("/", h.AddCustomer)
    routes.GET("/", h.GetCustomers)
    routes.GET("/:id", h.GetCustomer)
    routes.PUT("/:id", h.UpdateCustomer)
    routes.DELETE("/:id", h.DeleteCustomer)
}
