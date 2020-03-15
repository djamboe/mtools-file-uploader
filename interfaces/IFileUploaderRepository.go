package interfaces

import (
	"github.com/djamboe/mtools-file-uploader/models"
)

type ILoginRepository interface {
	GetUserByEmailAndPassword(username string, password string) (models.UserModel, error)
}
