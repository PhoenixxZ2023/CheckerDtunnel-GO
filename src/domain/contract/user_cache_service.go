package contract

import (
	"time"

	"github.com/PhoenixxZ2023/CheckerDtunnel-GO/src/domain/entity"
)

type UserCacheService interface {
	Set(value *entity.User, ttl time.Duration)
	Get(username string) (*entity.User, bool)
}
