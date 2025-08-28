package app

import (
	httphandler "github.com/dwikikf/al-hikmah-attendance-api/internal/handler/httpHandler"
	repository_mysql "github.com/dwikikf/al-hikmah-attendance-api/internal/repository/mysql"
	"github.com/dwikikf/al-hikmah-attendance-api/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Dependency Injection
	// repo
	academicYearRepo := repository_mysql.NewAcademicYearRepository(db)
	// service
	academicYearUsecase := usecase.NewAcademicYearService(academicYearRepo)
	// handler
	academicYearhandler := httphandler.NewAcdemicYearHandler(academicYearUsecase)

	// route
	api := router.Group("/api/v1")
	{
		academicyears := api.Group("/academicyears")
		{
			academicyears.POST("/", academicYearhandler.Create)
			academicyears.GET("/:id", academicYearhandler.FindById)
		}
	}
	return router
}
