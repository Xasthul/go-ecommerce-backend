package service

import (
	"context"

	"github.com/Xasthul/go-ecommerce-backend/auth-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	r *repository.UserRepository
}

func NewAuthService(r *repository.UserRepository) *AuthService {
	return &AuthService{r: r}
}

func (s *AuthService) RegisterUser(ctx context.Context, email, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.r.CreateUser(ctx, email, string(hashed))
}
