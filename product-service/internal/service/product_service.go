package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/rabbitmq"
	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository"
	db "github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository/db/gen"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type ProductService struct {
	productRepository *repository.ProductRepository
	redisClient       *redis.Client
}

func NewProductService(
	productRepository *repository.ProductRepository,
	redisClient *redis.Client,
) *ProductService {
	return &ProductService{
		productRepository: productRepository,
		redisClient:       redisClient,
	}
}

func (s *ProductService) GetProducts(ctx context.Context) ([]db.Product, error) {
	return s.productRepository.GetProducts(ctx)
}

func (s *ProductService) GetProductById(ctx context.Context, productId uuid.UUID) (*db.Product, error) {
	cachedProduct, err := s.getCachedProduct(ctx, productId)
	if err == nil {
		return cachedProduct, nil
	}

	product, err := s.productRepository.GetProductById(ctx, productId)
	if err != nil {
		return nil, err
	}

	s.cacheProduct(ctx, product)

	return product, nil
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
	product, err := s.productRepository.UpdateProduct(
		ctx,
		productId,
		categoryID,
		name,
		description,
		priceCents,
		currency,
		stock,
	)
	if err != nil {
		return err
	}

	s.cacheProduct(ctx, product)

	return nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, productId uuid.UUID) error {
	err := s.productRepository.DeleteProduct(ctx, productId)
	if err != nil {
		return err
	}

	s.clearCacheForProduct(ctx, productId)

	return nil
}

func (s *ProductService) DecreaseStock(ctx context.Context, payload *rabbitmq.OrderCreatedEvent) {
	product, err := s.productRepository.DecreaseStock(ctx, payload.ProductID, payload.Quantity)
	if err != nil || product == nil {
		// handle failed to decrease stock
		return
	}

	s.cacheProduct(ctx, product)
}

func (s *ProductService) cacheProduct(ctx context.Context, product *db.Product) {
	cacheKey := s.getProductCacheKey(product.ID)

	bytes, _ := json.Marshal(product)
	s.redisClient.Set(ctx, cacheKey, bytes, 10*time.Minute)
}

func (s *ProductService) getCachedProduct(ctx context.Context, productId uuid.UUID) (*db.Product, error) {
	cacheKey := s.getProductCacheKey(productId)

	cached, err := s.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var product db.Product
		_ = json.Unmarshal([]byte(cached), &product)
		return &product, nil
	}
	return nil, err
}

func (s *ProductService) clearCacheForProduct(ctx context.Context, productId uuid.UUID) {
	cacheKey := s.getProductCacheKey(productId)

	s.redisClient.Del(ctx, cacheKey)
}

func (s *ProductService) getProductCacheKey(productId uuid.UUID) string {
	return fmt.Sprintf("product:%s", productId)
}
