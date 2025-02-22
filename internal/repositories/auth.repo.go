package repositories

import (
	"fmt"
	"time"

	"github.com/quangdvn/go-ec/global"
)

type IAuthRepository interface {
	AddOTP(email string, otp int, expiration int64) error
}

type authRepository struct {
}

func NewAuthRepository() IAuthRepository {
	return &authRepository{}
}

// AddOTP implements IAuthRepository.
func (ar *authRepository) AddOTP(email string, otp int, expiration int64) error {
	key := fmt.Sprintf("usr:%s:otp", email) // :usr:email:otp
	return global.Cache.SetEx(ctx, key, otp, time.Duration(expiration)).Err()
}
