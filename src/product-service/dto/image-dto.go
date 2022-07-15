package dto

import "chilindo/src/product-service/entity"

type CreateImageDTO struct {
	Image *entity.ProductImages
}

func NewCreateImageDTO(image *entity.ProductImages) *CreateImageDTO {
	return &CreateImageDTO{Image: image}
}

type ImageDTO struct {
	ImageId string
}
