package service

import "github.com/Xasthul/go-ecommerce-backend/auth-service/internal/repository"

type AuthService struct {
	r *repository.UserRepository
}

func NewAuthService(r *repository.UserRepository) *AuthService {
	return &AuthService{r: r}
}
