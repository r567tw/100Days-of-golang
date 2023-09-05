package models

import (
	_ "gorm.io/gorm"
)

type Customer struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"username"`
	Bio		  string ` json:"bio"`
	Account   string `json:"account"`
}
