package service

import (
	"chilindo/src/user-service/entity"
	"chilindo/src/user-service/repository"
	"log"
)

type IUserService interface {
	SignUp(user *entity.User) (*entity.User, error)
	//SignIn(dto *dto.LoginDTO) (*entity.User, error)
}

type UserService struct {
	UserRepository repository.UserRepository
}

func (u UserService) SignUp(user *entity.User) (*entity.User, error) {
	newUser, err := u.UserRepository.CreateUser(user)
	if err != nil {
		log.Println("SignUp: Error CreateUser in package service")
		return nil, err
	}
	return newUser, nil
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}
