package repository

import (
	"context"

	gen "github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository/db/gen"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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
		params.Description = pgtype.Text{String: *description, Valid: true}
	} else {
		params.Description = pgtype.Text{Valid: false}
	}

	if currency != nil {
		params.Currency = pgtype.Text{String: *currency, Valid: true}
	} else {
		params.Currency = pgtype.Text{Valid: false}
	}

	if stock != nil {
		params.Stock = pgtype.Int4{Int32: *stock, Valid: true}
	} else {
		params.Stock = pgtype.Int4{Valid: false}
	}

	return r.q.CreateProduct(ctx, params)
}

func (r *ProductRepository) UpdateProduct(
	ctx context.Context,
	productId uuid.UUID,
	categoryID *int16,
	name *string,
	description *string,
	priceCents *int32,
	currency *string,
	stock *int32,
) error {
	params := gen.UpdateProductParams{ID: productId}

	if categoryID != nil {
		params.CategoryID = pgtype.Int2{Int16: *categoryID, Valid: true}
	} else {
		params.CategoryID = pgtype.Int2{Valid: false}
	}

	if name != nil {
		params.Name = pgtype.Text{String: *name, Valid: true}
	} else {
		params.Name = pgtype.Text{Valid: false}
	}

	if description != nil {
		params.Description = pgtype.Text{String: *description, Valid: true}
	} else {
		params.Description = pgtype.Text{Valid: false}
	}

	if priceCents != nil {
		params.PriceCents = pgtype.Int4{Int32: *priceCents, Valid: true}
	} else {
		params.PriceCents = pgtype.Int4{Valid: false}
	}

	if currency != nil {
		params.Currency = pgtype.Text{String: *currency, Valid: true}
	} else {
		params.Currency = pgtype.Text{Valid: false}
	}

	if stock != nil {
		params.Stock = pgtype.Int4{Int32: *stock, Valid: true}
	} else {
		params.Stock = pgtype.Int4{Valid: false}
	}

	return r.q.UpdateProduct(ctx, params)
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, productId uuid.UUID) error {
	return r.q.DeleteProduct(ctx, productId)
}

func (r *ProductRepository) DecreaseStock(
	ctx context.Context,
	productId uuid.UUID,
	quantity int,
) (int64, error) {
	return r.q.DecreaseStock(ctx, gen.DecreaseStockParams{
		ID:    productId,
		Stock: int32(quantity),
	})
}
