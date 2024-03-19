package models

import "gorm.io/gorm"

type SampleMapping struct {
	gorm.Model
	CustomerSegment string `gorm:"not null"`
	ProductID       string `gorm:"not null"`
	SampleProductID string `gorm:"not null"`
}
