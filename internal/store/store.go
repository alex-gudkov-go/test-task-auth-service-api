package store

import (
	"context"
	"test-task-auth-service-api/internal/models"
)

type Store interface {
	SaveRefreshToken(ctx context.Context, refreshToken *models.RefreshToken) error
	FindRefreshTokenById(ctx context.Context, id string) (*models.RefreshToken, error)
	DeleteRefreshTokenById(ctx context.Context, id string) error
}
