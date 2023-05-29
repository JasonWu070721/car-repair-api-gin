package models

import "gorm.io/gorm"

type Car struct {
    gorm.Model
    Year       string `json:"year"`
    Make       string `json:"make"`
    Type       string `json:"type"`
	BodyStyles       string `json:"body_styles"`
}