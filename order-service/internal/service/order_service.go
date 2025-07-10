package service

import (
	"github.com/Xasthul/go-ecommerce-backend/order-service/internal/client"
	"github.com/Xasthul/go-ecommerce-backend/order-service/internal/repository"
)

type OrderService struct {
	orderRepository *repository.OrderRepository
	productClient   *client.ProductClient
}

func NewOrderService(
	orderRepository *repository.OrderRepository,
	productClient *client.ProductClient,
) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		productClient:   productClient,
	}
}
