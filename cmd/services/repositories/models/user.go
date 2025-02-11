package models

type User struct {
	CPF      string `json:"cpf"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
