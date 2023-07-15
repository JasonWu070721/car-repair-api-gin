package models

import "gorm.io/gorm"

type Customer struct {
    gorm.Model
    Name       string `json:"name"`
    LicensePlate       string `json:"license_plate"`
    CarColor       string `json:"car_color"`
    CarYear        string `json:"car_year"`
}

type UpdateCustomerRequestBody struct {
    Name       string `json:"name"`
    LicensePlate       string `json:"license_plate"`
    CarColor       string `json:"car_color"`
    CarYear        string `json:"car_year"`
}

type AddCustomerRequestBody struct {
    Name       string `json:"name"`
    LicensePlate       string `json:"license_plate"`
    CarColor       string `json:"car_color"`
    CarYear        string `json:"car_year"`
}