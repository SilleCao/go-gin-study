package entity

type Person struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Age       int8   `json:"age" binding:"gte=0,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}

type Video struct {
	Title       string `json:"title" binding:"min=2,max=10" validate:"is-cool"`
	Description string `json:"description" binding:"max=20"`
	URL         string `json:"url" binding:"required"`
	Author      Person `json:"author" binding:"required"`
}
