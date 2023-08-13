package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
    ID           uint `gorm:"primarykey" json:"id"`
    CreatedAt time.Time         `json:"created_at"`
    UpdatedAt time.Time         `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Customer struct {
	BaseModel
	Name         string         `json:"name"`
	LicensePlate string         `json:"license_plate"`
	CarColor     string         `json:"car_color"`
	CarYear      string         `json:"car_year"`
}

type UpdateCustomerRequestBody struct {
	Name         string `json:"name"`
	LicensePlate string `json:"license_plate"`
	CarColor     string `json:"car_color"`
	CarYear      string `json:"car_year"`
}

type AddCustomerRequestBody struct {
	Name         string `json:"name"`
	LicensePlate string `json:"license_plate"`
	CarColor     string `json:"car_color"`
	CarYear      string `json:"car_year"`
}