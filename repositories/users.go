package repositories

import (
	"myproject/database"
	"myproject/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}



func NewUserRepository() *UserRepository {
	return &UserRepository{db: database.GetDb()}
}


func (u *UserRepository) GetAllUsers()([]models.User,error) {
	var users []models.User
	err := u.db.Find(&users).Error
	return users,err
}


func (u *UserRepository) CreateUser(user *models.User) error {
	err := u.db.Create(&user).Error
	return err
}


func (u *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, id).Error
	return &user, err
}

func (u *UserRepository) UpdateUser(user *models.User) error {
	err := u.db.Save(&user).Error
	return err
}


func (u *UserRepository) DeleteUserByID(user *models.User) error {
	err := u.db.Delete(&user).Error
	return err
}