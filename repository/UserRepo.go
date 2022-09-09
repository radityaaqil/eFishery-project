package repository

import (
	"efishery-project/entities"
	"errors"
	"log"

	"gorm.io/gorm"
)

type RepositoryUser interface {
	CreateUser(user entities.User) (entities.UserResponse, error)
	FindEmail(email string) (entities.User, error)
	GetUserById(id int) (entities.User, error)
}

type Repository_User struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *Repository_User {
	return &Repository_User{db}
}

func (r *Repository_User) CreateUser(user entities.User) (entities.UserResponse, error) {
	newUser := entities.User{user.Model, user.Username, user.Email, user.Password}

	result := r.db.Create(&newUser)

	err := result.Error

	userResponse := entities.UserResponse{newUser.ID, newUser.Username, newUser.Email}

	if err != nil {
		log.Panic("Error inserting new user")
		return userResponse, err
	}

	return userResponse, nil
}

func (r *Repository_User) FindEmail(email string) (entities.User, error) {
	var user entities.User

	result := r.db.Where("email = ?", email).First(&user)

	err := result.Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository_User) GetUserById(id int) (entities.User, error) {
	var currentUser entities.User
	result := r.db.Find(&currentUser, id)

	err := result.Error

	if result.RowsAffected == 0 {
		return currentUser, errors.New("No user found")
	}

	if err != nil {
		log.Panic("Error inserting new user")
		return currentUser, err
	}

	return currentUser, nil
}
