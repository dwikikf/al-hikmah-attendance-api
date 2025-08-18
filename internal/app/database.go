package app

import (
	"fmt"
	"log"

	"github.com/dwikikf/al-hikmah-attendance-api/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Melakukan migrasi database...")
	if err := db.AutoMigrate(
		&domain.AcademicYear{},
		&domain.Teacher{},
		&domain.Class{},
		&domain.Student{},
		&domain.Attendance{},
	); err != nil {
		return nil, err
	}

	log.Println("Migrasi database berhasil.")

	return db, nil
}
