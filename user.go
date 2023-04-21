package go_project

type User struct {
	Id int `json:"-" db:"id"`
	First_name string `json:"first_name" binding:"required"`
	Last_name string 	`json:"last_name" binding:"required"`
	Password string `json:"password" binding:"required"` 
	Email string `json:"email" binding:"required"` 
	Confirmpassword string `json:"confirmPassword" binding:required"`
}