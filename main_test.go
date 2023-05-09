package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)


type ResponseUser struct {
	ID uint `json:"id"`
    UserName       string `json:"username"`
    Password       string `json:"password"`
    Email       string `json:"email"`
}

type ResponseUsers []struct {
	ResponseUser
}

func Test_setupRouter(t *testing.T) {

    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()

	router := SetRouter()

	id := test_AddUser(router, t)

	fmt.Printf("id: %d\n", id)
	
	test_GetUser(router, t, id)
	test_DeleteUser(router, t, id)
	test_GetUsers(router, t)
}

func test_GetUsers(router *gin.Engine, t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/", nil)
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println(w.Body.String())
	httpOut := w.Body.String()

	users := ResponseUsers{}

    err := json.Unmarshal([]byte(httpOut), &users)
    if err != nil {
		fmt.Println(err)
    }

	// fmt.Println(users)

	// assert.Contains(t, w.Body.String(), "Asia/Shanghai")
}

func test_AddUser(router *gin.Engine, t *testing.T) uint {

    // JSON body
	body := []byte(`{
		"username": "Jason Wu",
		"password": "MD5_password",
		"email": "jason_wu@aai.com.tw"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	fmt.Println(w.Body.String())
	
	assert.Equal(t, http.StatusCreated, w.Code)

	httpOut := w.Body.String()

	user := ResponseUser{}

    err := json.Unmarshal([]byte(httpOut), &user)
    if err != nil {
		fmt.Println(err)
    }

	fmt.Printf("user.ID: %d\n", user.ID)
	return user.ID
	// assert.Contains(t, w.Body.String(), "Asia/Shanghai")
}


func test_GetUser(router *gin.Engine, t *testing.T, id uint) {

	_id := strconv.FormatUint(uint64(id), 10)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/" + string(_id), nil)
	router.ServeHTTP(w, req)

	fmt.Println(w.Body.String())
	
	assert.Equal(t, http.StatusOK, w.Code)
	// assert.Contains(t, w.Body.String(), "Asia/Shanghai")
}

func test_DeleteUser(router *gin.Engine, t *testing.T, id uint) {

	_id := strconv.FormatUint(uint64(id), 10)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/" + string(_id), nil)
	router.ServeHTTP(w, req)

	fmt.Println(w.Body.String())
	
	assert.Equal(t, http.StatusNoContent, w.Code)
	// assert.Contains(t, w.Body.String(), "Asia/Shanghai")
}


