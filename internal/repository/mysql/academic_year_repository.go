package repository_mysql

import (
	"context"

	"github.com/dwikikf/al-hikmah-attendance-api/internal/domain"
	"github.com/dwikikf/al-hikmah-attendance-api/internal/usecase"
	"gorm.io/gorm"
)

type academicYearRepository struct {
	db *gorm.DB
}

func NewAcademicYearRepository(db *gorm.DB) usecase.AcademicYearRepository {
	return &academicYearRepository{
		db: db,
	}
}

func (r *academicYearRepository) Create(ctx context.Context, acacademicYear *domain.AcademicYear) error {
	return r.db.WithContext(ctx).Create(acacademicYear).Error
}

func (r *academicYearRepository) FindById(ctx context.Context, id uint) (*domain.AcademicYear, error) {
	var acacademic_year domain.AcademicYear
	err := r.db.WithContext(ctx).First(&acacademic_year, id).Error
	if err != nil {
		return nil, err
	}
	return &acacademic_year, nil

}
