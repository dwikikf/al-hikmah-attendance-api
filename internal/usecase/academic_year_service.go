package usecase

import (
	"context"

	"github.com/dwikikf/al-hikmah-attendance-api/internal/domain"
	"github.com/dwikikf/al-hikmah-attendance-api/internal/handler/httpHandler/dto"
)

type acacademicYearService struct {
	acacademicYearRepo AcademicYearRepository
}

func NewAcademicYearService(repo AcademicYearRepository) AcademicYearUsecase {
	return &acacademicYearService{
		acacademicYearRepo: repo,
	}
}

func (s *acacademicYearService) Create(ctx context.Context, req dto.CreateAcademicYearRequest) (*domain.AcademicYear, error) {
	acacademic_year := &domain.AcademicYear{
		Year:      req.Year,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Status:    domain.AcademicYearStatusActive,
	}

	err := s.acacademicYearRepo.Create(ctx, acacademic_year)
	if err != nil {
		return nil, err
	}
	return acacademic_year, nil
}

func (s *acacademicYearService) FindById(ctx context.Context, id uint) (*domain.AcademicYear, error) {
	return s.acacademicYearRepo.FindById(ctx, id)
}
