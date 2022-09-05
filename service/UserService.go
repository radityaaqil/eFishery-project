package service

import (
	"efishery-project/entities"
	helper "efishery-project/helpers"
	"efishery-project/repository"
	"fmt"
	"log"
)

type ServiceUser interface {
	CreateUserService(user entities.User) (entities.User, error)
	LoginService(user entities.User) (entities.User, error)
	GetUserByIdService(id int) (entities.User, error)
}

type Service_User struct {
	repositoryuser repository.RepositoryUser
}

func NewServiceUser(repositoryuser repository.RepositoryUser) *Service_User {
	return &Service_User{repositoryuser}
}

func (s *Service_User) CreateUserService(user entities.User) (entities.User, error) {
	var createUser entities.User

	matchingEmail, errMatchingEmail := s.repositoryuser.FindEmail(user.Email)

	if matchingEmail.Email == user.Email {
		return user, fmt.Errorf("Please insert different email")
	}

	if errMatchingEmail != nil && matchingEmail.ID == 0 {
		var errCreateUser error
		var err error

		user.Password, err = helper.GeneratePassword(user.Password)

		if err != nil {
			log.Panic("Failed to hash password")
			return user, err
		}

		createUser, errCreateUser := s.repositoryuser.CreateUser(user)

		if errCreateUser != nil {
			log.Panic("Failed to create user")
			return createUser, errCreateUser
		}
		return createUser, nil
	}
	return createUser, nil
}

func (s *Service_User) LoginService(user entities.User) (entities.User, error) {

	loginUser, errLoginUser := s.repositoryuser.FindEmail(user.Email)

	if errLoginUser != nil {
		return user, errLoginUser
	}

	comparePassword := helper.ComparePassword([]byte(loginUser.Password), []byte(user.Password))

	if comparePassword != nil {
		return user, fmt.Errorf("Credential mismatch!")
	}

	return loginUser, nil
}

func (s *Service_User) GetUserByIdService(id int) (entities.User, error) {
	currentUser, errCurrentUser := s.repositoryuser.GetUserById(id)

	if errCurrentUser != nil {
		return currentUser, errCurrentUser
	}

	return currentUser, nil
}
