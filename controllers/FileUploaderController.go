package controllers

import (
	"github.com/djamboe/mtools-file-uploader/interfaces"
	"github.com/djamboe/mtools-file-uploader/models"
)

type LoginController struct {
	interfaces.ILoginService
}

func (controller *LoginController) LoginProcess(username string, password string) (models.UserModel, error) {
	login, err := controller.DoLogin(username, password)
	if err != nil {
		panic(err)
	}
	return login, nil
}
