package service

import "github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository"

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
