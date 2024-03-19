package common

import "time"

type EntityBase struct {
	CreatedOn *time.Time `gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
	CreatedBy string     `gorm:"not null;type:nvarchar(60)"`
	UpdatedOn *time.Time `gorm:"type:datetime;default:NULL"`
	UpdatedBy string     `gorm:"type:varchar(60);default:NULL"`
	DeletedOn *time.Time `gorm:"type:datetime;default:NULL"`
	DeletedBy string     `gorm:"type:nvarchar(60);default:NULL"`
}
