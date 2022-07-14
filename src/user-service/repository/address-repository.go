package repository

import (
	"chilindo/src/user-service/entity"
	"gorm.io/gorm"
	"log"
)

type IAddressRepository interface {
	CreateAddress(address *entity.Address) (*entity.Address, error)
	UpdateAddress(address *entity.Address) (*entity.Address, error)
	DeleteAddress(id string) error
}

type AddressRepositoryDefault struct {
	db *gorm.DB
}

func NewAddressRepositoryDefault(db *gorm.DB) *AddressRepositoryDefault {
	return &AddressRepositoryDefault{db: db}
}

func (a AddressRepositoryDefault) CreateAddress(address *entity.Address) (*entity.Address, error) {
	if errCheckEmptyField := address.Validate(); errCheckEmptyField != nil {
		log.Println("CreateAddress: Error empty field in package repository", errCheckEmptyField)
		return nil, errCheckEmptyField
	}
	result := a.db.Create(&address)
	if result.Error != nil {
		log.Println("CreateAddress: Error Create in package repository", result)
		return nil, result.Error
	}
	return address, nil

}

func (a *AddressRepositoryDefault) UpdateAddress(address *entity.Address) (*entity.Address, error) {
	if errCheckEmptyField := address.Validate(); errCheckEmptyField != nil {
		log.Println("CreateAddress: Error empty field in package repository", errCheckEmptyField)
		return nil, errCheckEmptyField
	}

	var matchedAddress *entity.Address
	var count int64

	record := a.db.Where("user_id = ? AND id = ?", address.UserId, address.ID).Find(&matchedAddress).Count(&count)
	if record.Error != nil {
		log.Println("Error ne thang lol")
		return nil, record.Error
	}
	if count == 0 {
		log.Println("=0")
		return nil, nil
	}
	matchedAddress = address
	recordUpdate := a.db.Updates(&matchedAddress)
	if recordUpdate.Error != nil {
		log.Println("Error ne thang lol")
		return nil, recordUpdate.Error
	}
	return matchedAddress, nil
}

func (a AddressRepositoryDefault) DeleteAddress(id string) error {
	var deleteAddress *entity.Address
	resultFind := a.db.Where("id = ?", id).Find(&deleteAddress)
	if resultFind.Error != nil {
		log.Println("DeleteAddress: Error to find Address  in package repository", resultFind)
		return resultFind.Error
	}
	resultDelete := a.db.Delete(&deleteAddress)
	if resultDelete.Error != nil {
		log.Println("DeleteAddress: Error to Deleted Address  in package repository", resultDelete)
		return resultDelete.Error
	}
	return nil
}
