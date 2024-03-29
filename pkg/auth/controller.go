package auth

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

    routes := router.Group("/auth")
    routes.POST("/login", h.Login)
    routes.GET("/ping", h.GetPing)
}
