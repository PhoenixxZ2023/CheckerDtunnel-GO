package contract

import (
	"context"

	"github.com/PhoenixxZ2023/CheckerDtunnel-GO/src/domain/entity"
)

type UserDAO interface {
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
}
