package model

import "time"

type TravelAgent struct {
	ID           int       `json:"id" form:"id"`
	Name         string    `json:"name" form:"name"`
	Address      string    `json:"address" form:"address"`
	Phone_Number string    `json:"phone_number" form:"phone_number"`
	Created_At   time.Time `json:"created_at" form:"created_at"`
	Updated_At   time.Time `json:"updated_at" form:"updated_at"`
}
