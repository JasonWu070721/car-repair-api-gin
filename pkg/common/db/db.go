package db

import (
	"car_repair_api_go/pkg/common/models"
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.User{})
    db.AutoMigrate(&models.Customer{})
    SetApiVersionCache()

    return db
}

var ctx = context.Background()

func SetApiVersionCache() {

    redis_addr := viper.Get("REDIS_ADDR").(string)
    rdb := redis.NewClient(&redis.Options{
        Addr:     redis_addr,
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