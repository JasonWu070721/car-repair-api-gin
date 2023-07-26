package main

import (
	"car_repair_api_go/pkg/auth"
	"car_repair_api_go/pkg/cars"
	"car_repair_api_go/pkg/common/db"
	"car_repair_api_go/pkg/customers"
	"car_repair_api_go/pkg/maintenances"
	"car_repair_api_go/pkg/users"
	"time"

	"car_repair_api_go/docs"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
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
    dbHandler := db.Init(dbUrl)

    router := gin.Default()
    
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"*"},
        AllowHeaders:     []string{"*"},
        ExposeHeaders:    []string{"*"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
      }))

    apiGroup := router.Group("/api")
    v1Group := apiGroup.Group("/v1")

    users.RegisterRoutes(v1Group, dbHandler)
    auth.RegisterRoutes(v1Group, dbHandler)
    cars.RegisterRoutes(v1Group, dbHandler)
    maintenances.RegisterRoutes(v1Group, dbHandler)
    customers.RegisterRoutes(v1Group, dbHandler)

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
