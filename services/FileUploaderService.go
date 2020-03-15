package services

import (
	"github.com/djamboe/mtools-file-uploader/interfaces"
	"github.com/djamboe/mtools-file-uploader/models"
)

type LoginService struct {
	interfaces.ILoginRepository
}

func (service *LoginService) DoLogin(username string, password string) (models.UserModel, error) {
	var user models.UserModel
	user, err := service.GetUserByEmailAndPassword(username, password)
	if err != nil {
		panic(err)
	}
	return user, nil
}
