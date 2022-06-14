package model

import "time"

type Location struct {
	ID         int       `json:"id" form:"id"`
	ID_Product int       `json:"id_product" form:"id_product"`
	Name       string    `json:"name" form:"name"`
	Created_At time.Time `json:"created_at" form:"created_at"`
	Updated_At time.Time `json:"updated_at" form:"updated_at"`
}
