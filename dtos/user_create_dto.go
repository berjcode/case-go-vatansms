package dtos

type UserCreateDto struct {
	UserName     string
	NameSurname  string
	Email        string
	PasswordHash string
	CreatedBy    string
}
