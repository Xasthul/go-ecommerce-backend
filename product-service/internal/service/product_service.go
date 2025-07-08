package service

import (
	"context"

	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository"
	db "github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository/db/gen"
)

type ProductService struct {
	productRepository  *repository.ProductRepository
	categoryRepository *repository.CategoryRepository
}

func NewProductService(
	productRepository *repository.ProductRepository,
	categoryRepository *repository.CategoryRepository,
) *ProductService {
	return &ProductService{
		productRepository:  productRepository,
		categoryRepository: categoryRepository,
	}
}

func (s *ProductService) GetProducts(ctx context.Context) ([]db.Product, error) {
	return s.productRepository.GetProducts(ctx)
}
