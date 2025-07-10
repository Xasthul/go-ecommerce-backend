package service

import "github.com/Xasthul/go-ecommerce-backend/order-service/internal/repository"

type OrderService struct {
	r *repository.OrderRepository
}

func NewOrderService(r *repository.OrderRepository) *OrderService {
	return &OrderService{r: r}
}
