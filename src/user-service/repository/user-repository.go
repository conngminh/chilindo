package repository

import (
	"chilindo/src/user-service/dto"
	"chilindo/src/user-service/entity"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	VerifyCredential(loginDTO *dto.UserLoginDTO) (*entity.User, error)
	InsertUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) *entity.User
	IsDuplicateEmail(email string) bool
	FindByEmail(email string) *entity.User
	ProfileUser(userID string) *entity.User
}

type UserRepositoryDefault struct {
	db *gorm.DB
}

func NewUserRepositoryDefault(db *gorm.DB) *UserRepositoryDefault {
	return &UserRepositoryDefault{db: db}
}

func (u UserRepositoryDefault) InsertUser(user *entity.User) (*entity.User, error) {
	if errHashPassword := user.HashPassword(user.Password); errHashPassword != nil {
		log.Println("CreateUser: Error in package repository", errHashPassword)
		return nil, errHashPassword
	}
	result := u.db.Create(&user)
	if result.Error != nil {
		log.Println("CreateUser: Error in package repository", result.Error)
		return nil, result.Error
	}
	return user, nil
}

func (u UserRepositoryDefault) UpdateUser(user *entity.User) *entity.User {
	//if user.Password != "" {
	//	user.Password, _ = user.HashPassword(user.Password)
	//} else {
	//	var tempUser entity.User
	//	u.db.Find(&tempUser, user.ID)
	//	user.Password = tempUser.Password
	//}
	//
	//u.db.Save(&user)
	return user
}

func (u UserRepositoryDefault) IsDuplicateEmail(email string) bool {
	var user *entity.User
	u.db.Where("email = ?", email).Find(&user)
	if user.Email == email {
		return true
	}
	return false
}

func (u UserRepositoryDefault) FindByEmail(email string) *entity.User {
	var user *entity.User
	u.db.Where("email = ?", email).Find(&user)

	return user
}

func (u UserRepositoryDefault) ProfileUser(userID string) *entity.User {
	var user *entity.User
	u.db.Preload("Books").Preload("Books.User").Find(&user, userID)
	return user
}

func (u UserRepositoryDefault) VerifyCredential(loginDTO *dto.UserLoginDTO) (*entity.User, error) {
	var user *entity.User

	//errCheckEmptyField := user.Validate("login")
	//
	//if errCheckEmptyField != nil {
	//	log.Println("VerifyCredential: Error empty field in package repository", errCheckEmptyField)
	//	return nil, errCheckEmptyField
	//}
	res := u.db.Where("email = ?", loginDTO.Email).Find(&user)
	fmt.Println(user)
	if res.Error != nil {
		log.Println("VerifyCredential: Error find username in package repository", res.Error)
		return nil, res.Error
	}
	if err := user.CheckPassword(loginDTO.Password); err != nil {
		log.Println("VerifyCredential: Error in check password package repository")
		return nil, err
	}
	return user, nil
}
