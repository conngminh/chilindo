package repository

import (
	"chilindo/src/user-service/entity"
	"errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	//GetUserById(id int) (*entity.User, error)
	//CheckUsernameAndPassword(userName, passWord string) (string, error)
	//CreateAddress(userId uint)
}
type UserRepositoryDefault struct {
	db *gorm.DB
}

//func (u UserRepositoryDefault) CheckUsernameAndPassword(userName, passWord string) (string, error) {
//	var user *entity.User
//	errUsernameExist := u.db.Debug().Model(entity.User{}).Where("username = ? ", userName).Take(&user).Error
//
//	if errUsernameExist != nil {
//		errUsernameExist = errors.New("wrong username")
//		return "", errUsernameExist
//	} else {
//		isPasswordMatch := utils.VerifyPassword(user.Password, passWord)
//		if !isPasswordMatch {
//			errPassword := errors.New("wrong password")
//			return "", errPassword
//		}
//	}
//	return strconv.Itoa(user.Id), nil
//}

func (u UserRepositoryDefault) CreateUser(user *entity.User) (*entity.User, error) {
	var userMayDuplicated *entity.User
	u.db.Where("Username = ?", user.Username).Find(&userMayDuplicated)
	if userMayDuplicated.Username == user.Username {
		err := errors.New("username existed")
		return nil, err
	}
	//may the code to check password format be here

	result := u.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

//func (u UserRepositoryDefault) GetUserById(id int) (*entity.User, error) {
//	//TODO implement me
//	panic("implement me")
//}
//

func NewUserRepositoryDefault(db *gorm.DB) *UserRepositoryDefault {
	return &UserRepositoryDefault{db: db}
}
