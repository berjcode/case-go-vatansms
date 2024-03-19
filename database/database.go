package database

import (
	"berjcode/dependency/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDB() (*gorm.DB, error) {
	return gorm.Open("mysql", "root:123456@tcp(localhost:3306)/newdb?charset=utf8&parseTime=True&loc=Local")
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}, &models.Plan{}, &models.PlanStatus{}).Error; err != nil {
		return err
	}
	return nil
}
