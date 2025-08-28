package dto

import "time"

type CreateAcademicYearRequest struct {
	Year      string    `json:"year" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
	Status    string    `json:"status" binding:"required,oneof=active archived"`
}

type AcademicYearResponse struct {
	ID        uint      `json:"id"`
	Year      string    `json:"year"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
