package entity

type User struct {
	Name string `json:"Name" valid:"required~Name is required"`
}
