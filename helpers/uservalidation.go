package helpers

import (
	"berjcode/dependency/constant"
	"berjcode/dependency/models"
	"errors"
)

func ValidateLoginForm(loginModel models.UserLogin) error {
	if loginModel.Password == "" || loginModel.UsernameAndEmail == "" {
		return errors.New(constant.RequriedField)
	}

	if len(loginModel.Password) > 2 && len(loginModel.Password) < 20 {
		return errors.New(constant.MustBetweenTwoAndTwentyCharacters)
	}

	if len(loginModel.UsernameAndEmail) > 2 && len(loginModel.UsernameAndEmail) < 30 {
		return errors.New(constant.MustBeLargerTwoCharacters)
	}

	return nil
}
