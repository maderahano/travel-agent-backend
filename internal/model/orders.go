package model

import "time"

type Order struct {
	ID                 int       `json:"id" form:"id"`
	ID_Product         int       `json:"id_product" form:"id_product"`
	ID_Travel_Agent    int       `json:"id_travel_agent" form:"id_travel_agent"`
	Order_Date         time.Time `json:"order_date" form:"order_date"`
	Order_Name         string    `json:"order_name" form:"order_name"`
	Order_Phone_Number string    `json:"order_phone_number" form:"order_phone_number"`
	Order_address      string    `json:"order_address" form:"order_address"`
	Created_At         time.Time `json:"created_at" form:"created_at"`
	Updated_At         time.Time `json:"updated_at" form:"updated_at"`
}
