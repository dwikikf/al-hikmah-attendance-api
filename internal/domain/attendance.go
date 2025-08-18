package domain

import (
	"time"

	"gorm.io/gorm"
)

type AttendanceStatus string

const (
	Present AttendanceStatus = "present"
	Sick    AttendanceStatus = "sick"
	Excused AttendanceStatus = "excused"
	Absent  AttendanceStatus = "absent"
)

type Attendance struct {
	gorm.Model
	AttendanceDate time.Time        `gorm:"column:attendance_date;type:date;not null"`
	Status         AttendanceStatus `gorm:"column:status;type:enum('present','sick','excused','absent');not null"`
	CheckInTime    *time.Time       `gorm:"column:check_in_time;type:time;"`
	Notes          *string          `gorm:"column:notes;type:text;"`
	RecordedBy     *uint            `gorm:"column:recorded_by"`
	StudentID      uint             `gorm:"not null"` //Foreign Key Field

	// Belongs To: Satu data Absensi milik satu Siswa
	Student Student `gorm:"foreignKey:StudentID"`
}
