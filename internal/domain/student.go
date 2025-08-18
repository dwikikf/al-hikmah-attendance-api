package domain

import (
	"time"

	"gorm.io/gorm"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type StudentStatus string

const (
	StatusActive    StudentStatus = "active"
	StatusGraduated StudentStatus = "graduated"
	StatusInactive  StudentStatus = "inactive"
)

type Student struct {
	gorm.Model
	StudentNISN string        `gorm:"column:student_NISN;type:varchar(10);not null;unique"`
	FullName    string        `gorm:"column:full_name;type:varchar(100);not null"`
	Gender      Gender        `gorm:"column:gender;type:enum('male,'female');not null"`
	DateOfBirth *time.Time    `gorm:"column:date_of_birth"`
	Addrress    *string       `gorm:"column:address;type:text"`
	Status      StudentStatus `gorm:"column:status;type:enum('active','graduated','inactive')"`
	ClassID     uint          `gorm:"not null"` // Foreign Key Filed

	// Belongs To: Satu Siswa milik satu Kelas
	Class Class `gorm:"foreignKey:ClassID"`

	// Has Many: Satu Siswa memiliki banyak data Absensi
	Attendances []Attendance `gorm:"foreignKey:StudentID"`
}
