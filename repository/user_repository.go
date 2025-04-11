package repository

import "password-manager/models"

type UserRepository interface {
	Create(user models.User) error
	UpdateByUUID(uuid string, dataUser *models.User) (*models.User, error)
	DeleteByUUID(uuid string) error
	ListarUsers() ([]models.User, error)
}
