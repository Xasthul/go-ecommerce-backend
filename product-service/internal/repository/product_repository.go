package repository

import (
	"context"

	gen "github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository/db/gen"
	"github.com/google/uuid"
)

type ProductRepository struct {
	q *gen.Queries
}

func NewProductRepository(q *gen.Queries) *ProductRepository {
	return &ProductRepository{q: q}
}

func (r *ProductRepository) GetProducts(ctx context.Context) ([]gen.Product, error) {
	products, err := r.q.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetProductById(ctx context.Context, productId uuid.UUID) (*gen.Product, error) {
	product, err := r.q.GetProductByID(ctx, productId)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) CreateProduct(
	ctx context.Context,
	categoryID int16,
	name string,
	description *string,
	priceCents int32,
	currency *string,
	stock *int32,
) error {
	params := gen.CreateProductParams{
		CategoryID: categoryID,
		Name:       name,
		PriceCents: priceCents,
	}
	if description != nil {
		params.Description = *description
	}
	if currency != nil {
		params.Currency = *currency
	}
	if stock != nil {
		params.Stock = *stock
	}

	return r.q.CreateProduct(ctx, params)
}
