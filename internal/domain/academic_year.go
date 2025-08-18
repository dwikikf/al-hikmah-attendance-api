package domain

import (
	"time"

	"gorm.io/gorm"
)

type AcademicYearStatus string

const (
	AcademicYearStatusActive   AcademicYearStatus = "active"
	AcademicYearStatusArchived AcademicYearStatus = "arcihived"
)

type AcademicYear struct {
	gorm.Model
	Year      string             `gorm:"column:year;type:varchar(10);not null;unique"`
	StartDate time.Time          `gorm:"column:start_date;not null"`
	EndDate   time.Time          `gorm:"column:end_date;not null"`
	Status    AcademicYearStatus `gorm:"column:status;type:enum('active','archived');not null"`

	// Has Many: Satu Tahun Ajaran memiliki banyak Kelas
	Classes []Class `gorm:"foreignKey:AcademicYearID"`
}
