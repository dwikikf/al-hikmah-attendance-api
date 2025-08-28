package usecase

import (
	"context"

	"github.com/dwikikf/al-hikmah-attendance-api/internal/domain"
	"github.com/dwikikf/al-hikmah-attendance-api/internal/handler/httpHandler/dto"
)

type AcademicYearRepository interface {
	Create(ctx context.Context, academic_year *domain.AcademicYear) error
	FindById(ctx context.Context, id uint) (*domain.AcademicYear, error)
}

type AcademicYearUsecase interface {
	Create(ctx context.Context, req dto.CreateAcademicYearRequest) (*domain.AcademicYear, error)
	FindById(ctx context.Context, id uint) (*domain.AcademicYear, error)
}
