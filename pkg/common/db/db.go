package db

import (
	"log"

	"car_repair_api_go/pkg/common/models"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"

    "context"
    "github.com/redis/go-redis/v9"
    "fmt"

)

func Init(url string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.User{})
    SetApiVersionCache()

    return db
}

var ctx = context.Background()

func SetApiVersionCache() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "",
        DB:       0,
    })

    err := rdb.Set(ctx, "api_verion", "1.0.0", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "api_verion").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("api_verion", val)
}