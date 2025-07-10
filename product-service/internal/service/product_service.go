package service

import (
	"context"

	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/rabbitmq"
	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository"
	db "github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository/db/gen"
	"github.com/google/uuid"
)

type ProductService struct {
	productRepository *repository.ProductRepository
}

func NewProductService(
	productRepository *repository.ProductRepository,
) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (s *ProductService) GetProducts(ctx context.Context) ([]db.Product, error) {
	return s.productRepository.GetProducts(ctx)
}

func (s *ProductService) GetProductById(ctx context.Context, productId uuid.UUID) (*db.Product, error) {
	return s.productRepository.GetProductById(ctx, productId)
}

func (s *ProductService) CreateProduct(
	ctx context.Context,
	categoryID int16,
	name string,
	description *string,
	priceCents int32,
	currency *string,
	stock *int32,
) error {
	return s.productRepository.CreateProduct(
		ctx,
		categoryID,
		name,
		description,
		priceCents,
		currency,
		stock,
	)
}

func (s *ProductService) UpdateProduct(
	ctx context.Context,
	productId uuid.UUID,
	categoryID *int16,
	name *string,
	description *string,
	priceCents *int32,
	currency *string,
	stock *int32,
) error {
	return s.productRepository.UpdateProduct(
		ctx,
		productId,
		categoryID,
		name,
		description,
		priceCents,
		currency,
		stock,
	)
}

func (s *ProductService) DeleteProduct(ctx context.Context, productId uuid.UUID) error {
	return s.productRepository.DeleteProduct(ctx, productId)
}

func (s *ProductService) DecreaseStock(ctx context.Context, payload *rabbitmq.OrderCreatedEvent) {
	rowsImpacted, err := s.productRepository.DecreaseStock(ctx, payload.ProductID, payload.Quantity)
	if err != nil || rowsImpacted <= 0 {
		// handle failed to decrease stock
	}
}
