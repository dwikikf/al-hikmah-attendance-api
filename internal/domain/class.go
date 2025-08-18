package domain

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name           string `gorm:"column:name;type:varchar(100);not null"`
	GradeLevel     int    `gorm:"column:grade_level;type:tinyint;not null"`
	AcademicYearID uint   `gorm:"not null"` // Foreign Key Fileds
	TeacherID      *uint  // Foreign Key Fileds

	// Belongs To: Satu Kelas milik satu Tahun Ajaran
	AcademicYear AcademicYear `gorm:"foreignKey:AcademicYearID"`

	//Belongs To: Satu Kelas memiliki satu Wali Kelas (Teacher)
	Teacher *Teacher `gorm:"foreignKey:TeacherID"`

	// Has Many: Satu Kelas memiliki banyak Siswa
	Students []Student `gorm:"foreignKey:ClassID"`
}
