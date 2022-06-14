package model

import "time"

type Product struct {
	ID              int       `json:"id" form:"id"`
	ID_Travel_Agent int       `json:"id_travel_agent" form:"id_travel_agent"`
	Name            string    `json:"name" form:"name"`
	Price           int       `json:"price" form:"price"`
	Description     string    `json:"description" form:"description"`
	Tax             int       `json:"tax" form:"tax"`
	Lodging         int       `json:"lodging" form:"lodging"`
	Airport_Pickup  int       `json:"airport_pickup" form:"airport_pickup"`
	Refund          int       `json:"refund" form:"refund"`
	Created_At      time.Time `json:"created_at" form:"created_at"`
	Updated_At      time.Time `json:"updated_at" form:"updated_at"`
}
