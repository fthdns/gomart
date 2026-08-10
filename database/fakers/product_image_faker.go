package fakers

import (
	"log"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/fthdns/gomart/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ProductImageFaker(db *gorm.DB) *models.ProductImage {
	product := ProductFaker(db)
	err := db.Create(&product).Error
	if err != nil {
		log.Fatal(err)
	}

	return &models.ProductImage{
		ID:        uuid.New().String(),
		Product:   models.Product{},
		ProductID: product.ID,
		Path:      faker.URL(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}
