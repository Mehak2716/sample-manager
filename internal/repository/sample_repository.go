package repository

import (
	"fmt"

	"github.com/Mehak2716/sample-manager/internal/models"
	"gorm.io/gorm"
)

type SampleRepository struct {
	DB *gorm.DB
}

func (repo *SampleRepository) Save(sampleMapping *models.SampleMapping) (*models.SampleMapping, error) {
	res := repo.DB.Create(sampleMapping)
	fmt.Println(sampleMapping)
	if res.Error != nil {
		return nil, res.Error
	}
	return sampleMapping, nil
}

func (repo *SampleRepository) IsExists(customerSegment string, productID string) bool {
	var count int64
	repo.DB.Model(&models.SampleMapping{}).
		Where("sample_mappings.customer_segment = ? AND sample_mappings.product_id = ?", customerSegment, productID).
		Count(&count)

	return count > 0
}
