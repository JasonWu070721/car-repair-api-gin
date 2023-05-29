package models

import "gorm.io/gorm"

type Maintenance struct {
    gorm.Model
    UserID       string `json:"user_id"`
    CarID       string `json:"car_id"`
	AppointmentTime       string `json:"appointment_time"`
}