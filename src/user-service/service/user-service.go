package service

import (
	"chilindo/src/user-service/dto"
	"chilindo/src/user-service/entity"
	"chilindo/src/user-service/repository"
	"github.com/mashingan/smapping"
	"log"
)

type IUserService interface {
	Update(user *dto.UserUpdateDTO) *entity.User
	VerifyCredential(loginDTO *dto.UserLoginDTO) (*entity.User, error)
	CreateUser(user *entity.User) *entity.User
	FindByEmail(email string) *entity.User
	IsDuplicateEmail(email string) bool
}

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserServiceDefault(userRepository repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}
func (u *UserService) VerifyCredential(loginDTO *dto.UserLoginDTO) (*entity.User, error) {
	user, err := u.UserRepository.VerifyCredential(loginDTO)

	if err != nil {
		log.Println("SignIn: Error VerifyCredential in package service")
		return nil, err
	}
	return user, nil
}

func (u *UserService) CreateUser(user *entity.User) *entity.User {
	res, _ := u.UserRepository.InsertUser(user)
	return res
}

func (u *UserService) FindByEmail(email string) *entity.User {
	return u.UserRepository.FindByEmail(email)
}

func (u *UserService) IsDuplicateEmail(email string) bool {
	res := u.UserRepository.IsDuplicateEmail(email)
	return res
}

//func ComparePassword(hashedPwd string, plainPassword []byte) bool {
//	byteHash := []byte(hashedPwd)
//	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
//	if err != nil {
//		log.Println(err)
//		return false
//	}
//	return true
//}

func (service *UserService) Update(user *dto.UserUpdateDTO) *entity.User {
	var userToUpdate *entity.User
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.UserRepository.UpdateUser(userToUpdate)
	return updatedUser
}
