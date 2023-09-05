package dtos

type CreateItemInput struct {
	Name  string  `json:"name" binding:"required,alphaunicode,min=2"`
	Price float64 `json:"price" binding:"required,gte=1"`
}

type UpdateItemInput struct {
	Name  string  `json:"name" binding:"alphaunicode,min=2"`
	Price float64 `json:"price" binding:"gte=1"`
}
