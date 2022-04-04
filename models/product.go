package models

import (
	"time"
)

type (
	Product struct {
		ID                uint            `json:"id" gorm:"primary_key"`
		Title             string          `json:"title"`
		Year              int             `json:"year"`
		ProductCategoryID uint            `json:"productCategoryID"`
		CreatedAt         time.Time       `json:"created_at"`
		UpdatedAt         time.Time       `json:"updated_at"`
		Product           ProductCategory `json:"-"`
	}
)
