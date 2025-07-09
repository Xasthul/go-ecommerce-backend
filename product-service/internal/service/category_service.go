package service

import (
	"context"

	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository"
	db "github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository/db/gen"
)

type CategoryService struct {
	categoryRepository *repository.CategoryRepository
}

func NewCategoryService(
	categoryRepository *repository.CategoryRepository,
) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, name string) (*db.Category, error) {
	return s.categoryRepository.CreateCategory(ctx, name)
}
