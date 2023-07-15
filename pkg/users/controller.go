package users

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

    routes := router.Group("/users")
    routes.POST("/", h.AddUser)
    routes.GET("/", h.GetUsers)
    routes.GET("/:id", h.GetUser)
    routes.PUT("/:id", h.UpdateUser)
    routes.DELETE("/:id", h.DeleteUser)

    routes.GET("/check_user", h.CheckUser)
}
