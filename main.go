package main

import (
	"car_repair_api_go/pkg/auth"
	"car_repair_api_go/pkg/cars"
	"car_repair_api_go/pkg/common/db"
	"car_repair_api_go/pkg/customers"
	"car_repair_api_go/pkg/maintenances"
	"car_repair_api_go/pkg/users"

	"car_repair_api_go/docs"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
    customers.RegisterRoutes(router, dbHandler)

    router.GET("/", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "dbUrl": dbUrl,
        })
    })

	docs.SwaggerInfo.Title = "Car Repair API"
	docs.SwaggerInfo.Description = "This is a car repair management API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}


    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return router
}
