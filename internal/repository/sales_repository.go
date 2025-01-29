package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/adityapersonal6/pos/internal/models"
)

type SalesRepository struct {
	db *sql.DB
}

// NewSalesRepository creates a new instance of SalesRepository
func NewSalesRepository(db *sql.DB) *SalesRepository {
	return &SalesRepository{db}
}

// CreateSale creates a new sale
func (r *SalesRepository) CreateSale(sale *models.Sale) error {
	query := "INSERT INTO sales (item, quantity, price) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, sale.Item, sale.Quantity, sale.Price).Scan(&sale.ID)
	if err != nil {
		log.Printf("could not insert sale: %v", err)
		return err
	}
	return nil
}

// GetSalesByID retrieves a sale by its ID
func (r *SalesRepository) GetSaleByID(id int) (*models.Sale, error) {
	query := "SELECT id, item, quantity, price FROM sales WHERE id = $1"
	sale := &models.Sale{}
	err := r.db.QueryRow(query, id).Scan(&sale.ID, &sale.Item, &sale.Quantity, &sale.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("sale with ID %d not found", id)
		}
		log.Printf("Error retrieving sale: %v", err)
		return nil, err
	}
	return sale, nil
}

func (r *SalesRepository) GetAllSales() ([]*models.Sale, error) {
	query := "SELECT id, item, quantity, price FROM sales"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("could not retrieve sales: %v", err)
		return nil, err
	}
	defer rows.Close()
	var sales []*models.Sale
	for rows.Next() {
		sale := &models.Sale{}
		err := rows.Scan(&sale.ID, &sale.Item, &sale.Quantity, &sale.Price)
		if err != nil {
			log.Printf("could not scan sale: %v", err)
			return nil, err
		}
		sales = append(sales, sale)
	}
	if err := rows.Err(); err != nil {
		log.Printf("rows error: %v", err)
		return nil, err
	}
	return sales, nil
}

// UpdateSale updates a sale
func (r *SalesRepository) UpdateSale(sale *models.Sale) error {
	query := "UPDATE sales SET item = $1, quantity = $2, price = $3 WHERE id = $4"
	_, err := r.db.Exec(query, sale.Item, sale.Quantity, sale.Price, sale.ID)
	if err != nil {
		log.Printf("could not update sale: %v", err)
		return err
	}
	return nil
}

// DeleteSale deletes a sale
func (r *SalesRepository) DeleteSale(id int) error {
	query := "DELETE FROM sales WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("could not delete sale: %v", err)
		return err
	}
	return nil
}
