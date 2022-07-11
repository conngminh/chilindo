package repository

import (
	"chilindo/src/admin-service/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type AdminRepository interface {
	UpdateAdmin(admin entity.Administrator) entity.Administrator
	VerifyCredential(email string, password string) interface{}
	InsertAdmin(admin entity.Administrator) entity.Administrator
	IsDuplicateEmail(email string) (tx *gorm.DB)
}
type adminConnection struct {
	connection *gorm.DB
}

func (db adminConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	//TODO implement me
	var admin entity.Administrator
	return db.connection.Where("email = ?", email).Take(&admin)
}

func (db adminConnection) InsertAdmin(admin entity.Administrator) entity.Administrator {
	//TODO implement me
	admin.Password = hashAndSalt([]byte(admin.Password))
	db.connection.Save(&admin)
	return admin
}

func (db adminConnection) VerifyCredential(email string, password string) interface{} {
	//TODO implement me
	var user entity.Administrator
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminConnection{
		connection: db,
	}
}
func (db adminConnection) UpdateAdmin(admin entity.Administrator) entity.Administrator {
	//TODO implement me
	if admin.Password != "" {
		admin.Password = hashAndSalt([]byte(admin.Password))
	} else {
		var tempUser entity.Administrator
		db.connection.Find(&tempUser, admin.Id)
		admin.Password = tempUser.Password
	}

	db.connection.Save(&admin)
	return admin
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
