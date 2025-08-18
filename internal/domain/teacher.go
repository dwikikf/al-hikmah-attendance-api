package domain

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	FullName    string  `gorm:"column:full_name;type:varchar(100); not null"`
	Email       *string `gorm:"column:email;type:varchar(100);unique"`
	PhoneNumber *string `gorm:"column:phone_number;type:varchar(20);unique"`

	// Has Many: Satu Guru bisa menjadi wali kelas di banyak kelas (misal di tahun ajaran berbeda)
	Classes []Class `gorm:"foreignKey:TeacherID"`
}
