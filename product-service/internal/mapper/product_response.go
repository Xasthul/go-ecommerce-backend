package mapper

import (
	"time"

	db "github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository/db/gen"
	"github.com/google/uuid"
)

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	CategoryID  int16     `json:"category_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PriceCents  int32     `json:"price_cents"`
	Currency    string    `json:"currency"`
	Stock       int32     `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewProductResponse(p db.Product) *ProductResponse {
	description := ""
	if p.Description.Valid {
		description = p.Description.String
	}

	return &ProductResponse{
		ID:          p.ID,
		CategoryID:  p.CategoryID.Int16,
		Name:        p.Name,
		Description: description,
		PriceCents:  p.PriceCents,
		Currency:    p.Currency,
		Stock:       p.Stock,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func NewProductListResponse(products []db.Product) []ProductResponse {
	response := make([]ProductResponse, 0, len(products))
	for _, p := range products {
		response = append(response, *NewProductResponse(p))
	}
	return response
}
