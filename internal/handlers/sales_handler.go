package handlers

import (
	"net/http"

	"github.com/adityapersonal6/pos/internal/models"
	"github.com/adityapersonal6/pos/internal/services"
	"github.com/gin-gonic/gin"
)

type SalesHandler struct {
	SalesService services.SalesService
}

func NewSalesHandler(salesService services.SalesService) *SalesHandler {
	return &SalesHandler{SalesService: salesService}
}

func (sh *SalesHandler) CreateSale(c *gin.Context) {
	var sale models.Sale
	if err := c.ShouldBindJSON(&sale); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := sh.SalesService.CreateSale(&sale)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &sale)
}

func (sh *SalesHandler) GetAllSales(c *gin.Context) {
	sales, err := sh.SalesService.GetAllSales()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sales)
}
