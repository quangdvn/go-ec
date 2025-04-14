package impls

import (
	"context"

	"github.com/quangdvn/go-ec/internal/database"
	"github.com/quangdvn/go-ec/internal/services"
)

type sUserLogin struct {
	r *database.Queries
}

func NewUserLoginImplement(r *database.Queries) services.IUserLogin {
	return &sUserLogin{
		r: r,
	}
}

func (s *sUserLogin) Login(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) Register(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) VerifyOtp(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) UpdatePassword(ctx context.Context) error {
	return nil
}
