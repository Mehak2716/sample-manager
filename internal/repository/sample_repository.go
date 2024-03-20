package repository

import (
	"github.com/Mehak2716/sample-manager/internal/models"
	"gorm.io/gorm"
)

type SampleRepository struct {
	DB *gorm.DB
}

func (repo *SampleRepository) Save(sampleMapping *models.SampleMapping) (*models.SampleMapping, error) {
	res := repo.DB.Create(sampleMapping)
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

func (repo *SampleRepository) FetchSampleIDs(customerSegments []string, productIDs []string) ([]models.SampleMapping, error) {

	query := `SELECT product_id,sample_product_id,customer_segment,id
	          FROM (SELECT  product_id, sample_product_id,customer_segment,
			  ROW_NUMBER() OVER (PARTITION BY product_id ORDER BY id) AS rn,id
	          FROM sample_mappings WHERE customer_segment IN (?) AND product_id IN (?))
	          AS subquery WHERE rn = 1`

	var sampleMappings []models.SampleMapping
	if err := repo.DB.Raw(query, customerSegments, productIDs).Scan(&sampleMappings).Error; err != nil {
		return nil, err
	}

	return sampleMappings, nil

}
