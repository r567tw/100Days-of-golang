package models

type Customer struct {
	ID        int64  `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Name      string `gorm:"type:varchar(20) NOT NULL;" json:"username,omitempty"`
	CreatedAt string `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt string `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}
