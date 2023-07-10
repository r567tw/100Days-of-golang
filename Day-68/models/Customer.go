package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	ID        int64  `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Name      string `gorm:"type:varchar(20) NOT NULL;" json:"username,omitempty"`
}
