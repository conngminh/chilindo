package repository

import (
	"chilindo/services/admin/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type AdminRepository interface {
	UpdateAdmin(admin entity.Administrator) entity.Administrator
}
type adminConnection struct {
	connection *gorm.DB
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
