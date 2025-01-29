package routes

import (
	"database/sql"

	"github.com/adityapersonal6/pos/internal/handlers"
	"github.com/adityapersonal6/pos/internal/repository"
	"github.com/adityapersonal6/pos/internal/services"
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the routes for the application
func SetupRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()

	repository := repository.NewSalesRepository(db)
	salesService := services.NewSalesService(*repository)
	salesHandler := handlers.NewSalesHandler(*salesService)

	router.POST("/api/v1/sales", salesHandler.CreateSale)
	router.GET("/api/v1/sales", salesHandler.GetAllSales)

	return router
}
