package repository

import (
	"chilindo/src/user-service/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	//CreateUser(user *entity.User) (*entity.User, error)
	VerifyCredential(email string, password string) interface{}
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userID string) entity.User
}

type UserRepositoryDefault struct {
	db *gorm.DB
}

func NewUserRepositoryDefault(db *gorm.DB) *UserRepositoryDefault {
	return &UserRepositoryDefault{db: db}
}

func (u *UserRepositoryDefault) InsertUser(user entity.User) entity.User {
	user.Password, _ = user.HashPassword(user.Password)
	u.db.Save(&user)
	return user
}

func (u *UserRepositoryDefault) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password, _ = user.HashPassword(user.Password)
	} else {
		var tempUser entity.User
		u.db.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	u.db.Save(&user)
	return user
}

func (u *UserRepositoryDefault) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return u.db.Where("email = ?", email).Take(&user)
}

func (u *UserRepositoryDefault) FindByEmail(email string) entity.User {
	var user entity.User
	u.db.Where("email = ?", email).Take(&user)
	return user
}

func (u *UserRepositoryDefault) ProfileUser(userID string) entity.User {
	var user entity.User
	u.db.Preload("Books").Preload("Books.User").Find(&user, userID)
	return user
}

func (u *UserRepositoryDefault) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := u.db.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}
