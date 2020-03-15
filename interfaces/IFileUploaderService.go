package interfaces

import "github.com/djamboe/mtools-file-uploader/models"

type ILoginService interface {
	DoLogin(username string, password string) (models.UserModel, error)
}
