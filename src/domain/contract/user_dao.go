package contract

import (
	"context"

	"github.com/DTunnel0/CheckUser-Go/src/domain/entity"
)

type UserDAO interface {
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
}
