package database

import (
	"berjcode/dependency/config"
	"berjcode/dependency/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDB(filename string) (*gorm.DB, error) {
	cfg, err := config.LoadConfiguration(filename)
	if err != nil {
		return nil, err
	}
	connectionString := cfg.Database.Username + ":" + cfg.Database.Password + "@tcp(" + cfg.Database.Host + ":" + cfg.Database.Port + ")/" + cfg.Database.Name + "?charset=utf8&parseTime=True&loc=Local"
	return gorm.Open("mysql", connectionString)
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}, &models.Plan{}, &models.PlanStatus{}, &models.Lesson{}).Error; err != nil {
		return err
	}
	return nil
}
