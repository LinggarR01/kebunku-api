package dto

type Tanaman struct {
	Id          string  `json:"id" form:"id" `
	Name        string  `json:"name" form:"name" binding:"required"`
	Description *string `json:"description" form:"description" `
	Class       string  `json:"class" form:"class" binding:"required"`
	Order       string  `json:"order" form:"order" binding:"required"`
	Family      string  `json:"family" form:"family" binding:"required"`
	Image       string  `json:"image"`
}
