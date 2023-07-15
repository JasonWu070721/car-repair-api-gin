package users

import (
	"net/http"

	"car_repair_api_go/pkg/common/models"

	"github.com/gin-gonic/gin"
)



func (h handler) UpdateUser(ctx *gin.Context) {
    id := ctx.Param("id")
    body := models.UpdateUserRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var user models.User

    if result := h.DB.First(&user, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    user.UserName = body.UserName
    user.Password = body.Password
    user.Email = body.Email

    h.DB.Save(&user)

    ctx.JSON(http.StatusOK, &user)
}
