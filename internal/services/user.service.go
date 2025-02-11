package services

import (
	"github.com/quangdvn/go-ec/internal/repositories"
	"github.com/quangdvn/go-ec/pkg/responses"
)

// type UserService struct {
// 	UserRepo *repositories.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{}
// }

// func (us *UserService) GetUserService() string {
// 	return us.UserRepo.GetUserInfo()
// }

// !! WITH INTERFACE
type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repositories.IUserRepository
}

// NewUserService creates a new instance of IUserService
func NewUserService(userRepo repositories.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

func (us *userService) Register(email string, purpose string) int {
	// return us.userRepo.Register(email, purpose)
	// 1. Check email exists
	if us.userRepo.GetUserByEmail(email) {
		return responses.ErrCodeUserIsExist

	}
	return responses.ErrCodeSuccess
}
