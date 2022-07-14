package repository

import (
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
	"gorm.io/gorm"
)

type ProductOptionRepository interface {
	CreateOption(b *dto.CreateOptionDTO) (*entity.ProductOption, error)
	GetOptions(b *dto.ProductIdDTO) (*[]entity.ProductOption, error)
	GetOptionByID(b *dto.OptionByIdDTO) (*entity.ProductOption, error)
	DeleteOption(b *dto.OptionIdDTO) (*entity.ProductOption, error)
	UpdateOption(b *dto.UpdateOptionDTO) (*entity.ProductOption, error)
}
type productOptionRepository struct {
	connection *gorm.DB
}

func newProductOptionRepository(connection *gorm.DB) *productOptionRepository {
	return &productOptionRepository{connection: connection}
}
