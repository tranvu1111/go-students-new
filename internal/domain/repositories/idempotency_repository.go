package repositories

import (
	"context"

	"github.com/tranvu1111/go-students-new/internal/domain/entities"
)

type IdempotencyRepository interface {
	FindByKey(ctx context.Context, key string) (*entities.IdempotencyRecord, error)
	Create(ctx context.Context, record *entities.IdempotencyRecord) (*entities.IdempotencyRecord, error)
	Update(ctx context.Context, record *entities.IdempotencyRecord) (*entities.IdempotencyRecord, error)
}