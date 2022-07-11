package service

import (
	"chilindo/src/user-service/dto"
	"chilindo/src/user-service/entity"
	"chilindo/src/user-service/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type IUserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.UserLoginDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserServiceDefault(userRepository repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}
func (u *UserService) VerifyCredential(email string, password string) interface{} {
	res := u.UserRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (u *UserService) CreateUser(user dto.UserLoginDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := u.UserRepository.InsertUser(userToCreate)
	return res
}

func (u *UserService) FindByEmail(email string) entity.User {
	return u.UserRepository.FindByEmail(email)
}

func (u *UserService) IsDuplicateEmail(email string) bool {
	res := u.UserRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
func (service *UserService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.UserRepository.UpdateUser(userToUpdate)
	return updatedUser
}
