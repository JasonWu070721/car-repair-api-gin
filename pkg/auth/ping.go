package auth

import (
	"fmt"
	"net/http"

	"car_repair_api_go/pkg/utils/token"

	"github.com/gin-gonic/gin"
)

func (h handler) GetPing(ctx *gin.Context) {

    uid, err := token.ExtractTokenID(ctx)

    if (err != nil) {
        fmt.Println("err:", err)
        ctx.JSON(http.StatusUnauthorized, []string{`"error":"error"`})
    } else {
        fmt.Println("uid:", uid)
    }

    ctx.JSON(http.StatusOK,[]string{`{"error":"error"}`})
}
