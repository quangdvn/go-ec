package services

import "github.com/quangdvn/go-ec/internal/repositories"

type UserService struct {
	UserRepo *repositories.UserRepo
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetUserService() string {
	return us.UserRepo.GetUserInfo()
}
