package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"username"`
	Bio		  string ` json:"bio"`
	CreatedAt int64      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt int64      `gorm:"autoUpdateTime" json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
