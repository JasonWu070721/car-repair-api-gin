package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    UserName       string `json:"username"`
    Password       string `json:"password"`
    Email       string `json:"email"`
}

type AddUserRequestBody struct {
    UserName       string `json:"username"`
    Password       string `json:"password"`
    Email       string `json:"email"`
}

type UpdateUserRequestBody struct {
    UserName       string `json:"username"`
    Password       string `json:"password"`
    Email       string `json:"email"`
}