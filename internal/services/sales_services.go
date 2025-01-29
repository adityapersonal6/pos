package services

import (
	"errors"

	"github.com/adityapersonal6/pos/internal/models"
	"github.com/adityapersonal6/pos/internal/repository"
)

type SalesService struct {
	repo repository.SalesRepository
}

// NewSalesService creates a new SalesService with the necessary dependencies
func NewSalesService(repo repository.SalesRepository) *SalesService {
	return &SalesService{repo: repo}
}

// CreateSale creates a new sale
func (s *SalesService) CreateSale(sale *models.Sale) error {
	if sale == nil {
		return errors.New("sale cannot be nil")
	}
	return s.repo.CreateSale(sale)
}

// GetSaleByID returns a sale by its ID
func (s *SalesService) GetSaleByID(id int) (*models.Sale, error) {
	return s.repo.GetSaleByID(id)
}

// GetAllSales returns all sales
func (s *SalesService) GetAllSales() ([]*models.Sale, error) {
	return s.repo.GetAllSales()
}

// UpdateSale updates a sale
func (s *SalesService) UpdateSale(sale *models.Sale) error {
	if sale == nil {
		return errors.New("sale cannot be nil")
	}
	return s.repo.UpdateSale(sale)
}

// DeleteSale deletes a sale
func (s *SalesService) DeleteSale(id int) error {
	return s.repo.DeleteSale(id)
}
