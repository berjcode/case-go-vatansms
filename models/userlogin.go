package models

type UserLogin struct {
	UsernameAndEmail string `json:"userNameAndEmail"`
	Password         string `json:"password"`
}
