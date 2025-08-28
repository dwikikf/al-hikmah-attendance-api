package httphandler

import (
	"net/http"
	"strconv"

	"github.com/dwikikf/al-hikmah-attendance-api/internal/handler/httpHandler/dto"
	"github.com/dwikikf/al-hikmah-attendance-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type academicYearHandler struct {
	academicYearUsecase usecase.AcademicYearUsecase
}

func NewAcdemicYearHandler(uc usecase.AcademicYearUsecase) *academicYearHandler {
	return &academicYearHandler{
		academicYearUsecase: uc,
	}
}

func (h *academicYearHandler) Create(c *gin.Context) {
	var req dto.CreateAcademicYearRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	academic_year, err := h.academicYearUsecase.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create academic year"})
		return
	}

	response := dto.AcademicYearResponse{
		ID:        academic_year.ID,
		Year:      academic_year.Year,
		StartDate: academic_year.StartDate,
		EndDate:   academic_year.EndDate,
		Status:    string(academic_year.Status),
		CreatedAt: academic_year.CreatedAt,
		UpdatedAt: academic_year.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *academicYearHandler) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid academic year id"})
		return
	}

	academic_year, err := h.academicYearUsecase.FindById(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "academic year not found"})
		return
	}

	response := dto.AcademicYearResponse{
		ID:        academic_year.ID,
		Year:      academic_year.Year,
		StartDate: academic_year.StartDate,
		EndDate:   academic_year.EndDate,
		Status:    string(academic_year.Status),
		CreatedAt: academic_year.CreatedAt,
		UpdatedAt: academic_year.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}
