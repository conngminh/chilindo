package repository

import (
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
	"gorm.io/gorm"
	"log"
)

type ProductImageRepository interface {
	CreateImage(b *dto.CreateImageDTO) (*entity.ProductImages, error)
	GetImage(b *dto.ProductIdDTO) (*[]entity.ProductImages, error)
	ProductImageByID(b *dto.ProductDTO) (int64, error)
}

func (p productImageRepository) GetImage(b *dto.ProductIdDTO) (*[]entity.ProductImages, error) {
	//TODO implement me
	var images *[]entity.ProductImages
	var count int64
	record := p.connection.Where("product_id = ?", b.ProductId).Find(&images).Count(&count)
	if record.Error != nil {
		log.Println("GetOptions : Error to get all option", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		log.Println("GetOptions : Not found Options", count)
		return nil, nil
	}
	return images, nil
}

func (p productImageRepository) ProductImageByID(b *dto.ProductDTO) (int64, error) {
	var count int64
	record := p.connection.Model(&entity.Product{}).Where("id = ?", b.ProductId).Count(&count)
	if record.Error != nil {
		log.Println("CountProductById: Get product by ID", record.Error)
		return count, record.Error
	}
	return count, nil
}

func (p productImageRepository) CreateImage(b *dto.CreateImageDTO) (*entity.ProductImages, error) {
	//TODO implement me
	record := p.connection.Create(&b.Image)
	if record.Error != nil {
		log.Println("CreateOption: Error to create repository")
		return nil, record.Error
	}
	return b.Image, nil
}

type productImageRepository struct {
	connection *gorm.DB
}

func NewProductImageRepository(connection *gorm.DB) *productImageRepository {
	return &productImageRepository{connection: connection}
}
