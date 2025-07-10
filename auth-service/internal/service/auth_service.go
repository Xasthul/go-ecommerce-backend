package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/Xasthul/go-ecommerce-backend/auth-service/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository  *repository.UserRepository
	tokenRepository *repository.TokenRepository
	jwtSecret       []byte
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAuthService(
	userRepository *repository.UserRepository,
	tokenRepository *repository.TokenRepository,
	jwtSecret string,
	accessTokenTTL, refreshTokenTTL int,
) *AuthService {
	return &AuthService{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
		jwtSecret:       []byte(jwtSecret),
		accessTokenTTL:  time.Duration(accessTokenTTL) * time.Minute,
		refreshTokenTTL: time.Duration(refreshTokenTTL) * time.Hour,
	}
}

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

type tokenClaims struct {
	Role int16 `json:"role"`
	jwt.RegisteredClaims
}

func (s *AuthService) RegisterUser(ctx context.Context, email, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.userRepository.CreateUser(ctx, email, string(hashed))
}

type issuedTokensDTO struct {
	AccessToken  string
	RefreshToken string
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*issuedTokensDTO, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, &AppError{Code: 401, Message: "invalid credentials"}
	}
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		return nil, &AppError{Code: 401, Message: "invalid credentials"}
	}
	return s.issueTokens(ctx, user.ID, user.Role)
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*issuedTokensDTO, error) {
	claims, err := s.parseRefresh(refreshToken)
	if err != nil {
		return nil, &AppError{Code: 401, Message: "invalid refresh token"}
	}

	hash := sha256.Sum256([]byte(refreshToken))

	savedRefreshToken, err := s.tokenRepository.GetRefreshToken(ctx, fmt.Sprintf("%x", hash[:]))
	if err != nil {
		return nil, &AppError{Code: 401, Message: "invalid refresh token"}
	}

	if s.tokenRepository.DeleteRefreshToken(ctx, savedRefreshToken.TokenHash) != nil {
		return nil, &AppError{Code: 500, Message: "failed to remove refresh token"}
	}

	return s.issueTokens(ctx, savedRefreshToken.UserID, claims.Role)
}

func (s *AuthService) issueTokens(ctx context.Context, userId uuid.UUID, role int16) (*issuedTokensDTO, error) {
	now := time.Now()
	accessTokenClaims := tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId.String(),
			ID:        uuid.NewString(),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Role: role,
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString(s.jwtSecret)
	if err != nil {
		return nil, &AppError{Code: 500, Message: "failed to create access token"}
	}

	refreshTokenExpiresAt := now.Add(s.refreshTokenTTL)
	refreshTokenClaims := tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId.String(),
			ID:        uuid.NewString(),
			ExpiresAt: jwt.NewNumericDate(refreshTokenExpiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Role: role,
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString(s.jwtSecret)
	if err != nil {
		return nil, &AppError{Code: 500, Message: "failed to create refresh token"}
	}

	refreshTokenHash := sha256.Sum256([]byte(refreshToken))
	if s.tokenRepository.CreateRefreshToken(ctx, userId, fmt.Sprintf("%x", refreshTokenHash[:]), refreshTokenExpiresAt) != nil {
		return nil, &AppError{Code: 500, Message: "failed to store refresh token"}
	}

	return &issuedTokensDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *AuthService) parseRefresh(token string) (*tokenClaims, error) {
	claims := &tokenClaims{}
	parsed, err := jwt.ParseWithClaims(
		token,
		claims,
		func(t *jwt.Token) (any, error) { return a.jwtSecret, nil },
	)
	if err != nil || !parsed.Valid {
		return nil, err
	}
	return claims, nil
}
