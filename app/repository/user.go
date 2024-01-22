package repository

import (
	"gorm.io/gorm"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/stori_challenge/app/models"
	"github.com/FrancoLiberali/stori_challenge/app/repository/conditions"
)

type IUserRepository interface {
	GetByEmail(db *gorm.DB, email string) (*models.User, error)
	Save(db *gorm.DB, user *models.User) error
}

type UserRepository struct{}

func (repository UserRepository) GetByEmail(db *gorm.DB, email string) (*models.User, error) {
	return cql.Query[models.User](
		db,
		conditions.User.Email.Is().Eq(email),
	).FindOne()
}

func (repository UserRepository) Save(db *gorm.DB, user *models.User) error {
	return db.Save(user).Error
}
