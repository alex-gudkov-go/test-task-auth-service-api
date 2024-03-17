package store

import (
	"context"
	"test-task-auth-service-api/models"
)

type Store interface {
	SaveRefreshToken(ctx context.Context, refreshToken *models.RefreshToken) error
}
