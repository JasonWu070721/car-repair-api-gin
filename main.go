package main

import (
	"car_repair_api_go/pkg/auth"
	"car_repair_api_go/pkg/cars"
	"car_repair_api_go/pkg/common/db"
	"car_repair_api_go/pkg/maintenances"
	"car_repair_api_go/pkg/users"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
    viper.SetConfigName("env")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("pkg/common/envs/")

    viper.ReadInConfig()
    port := viper.Get("PORT").(string)

    router := SetRouter()

    router.Run(port)
}

func SetRouter() *gin.Engine {

    dbUrl := viper.Get("DB_URL").(string)

    router := gin.Default()
    dbHandler := db.Init(dbUrl)

    users.RegisterRoutes(router, dbHandler)
    auth.RegisterRoutes(router, dbHandler)
    cars.RegisterRoutes(router, dbHandler)
    maintenances.RegisterRoutes(router, dbHandler)

    router.GET("/", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "dbUrl": dbUrl,
        })
    })

    return router
}
