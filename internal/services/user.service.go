package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/quangdvn/go-ec/internal/repositories"
	"github.com/quangdvn/go-ec/internal/utils"
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
	authRepo repositories.IAuthRepository
}

// NewUserService creates a new instance of IUserService
func NewUserService(
	userRepo repositories.IUserRepository,
	authRepo repositories.IAuthRepository,
) IUserService {
	return &userService{userRepo: userRepo, authRepo: authRepo}
}

func (us *userService) Register(email string, purpose string) int {
	fmt.Printf("Email:: %s\n", email)
	fmt.Printf("Purpose:: %s\n", purpose)
	// 0. hashEmail returns
	hashEmail := utils.GetHash(email)
	fmt.Printf("Hash nil:: %s\n", utils.GetHash(""))
	fmt.Printf("Hash email:: %s\n", hashEmail)

	// 1. Check exists in DB
	if us.userRepo.GetUserByEmail(email) {
		return responses.ErrCodeUserIsExist
	}

	// 2. Create new OTP
	otp := utils.GenerateSixRandomDigit()
	if purpose == "TEST_USER" {
		otp = 123456
	}
	fmt.Printf("OTP is:: %d\n", otp)

	// 3. Save OTP to Redis with expiration
	err := us.authRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return responses.ErrCodeInvalidOTP
	}

	// 4. Send email
	// err = utils.SendTextEmail([]string{email}, "quangdvn@gmail.com", strconv.Itoa(otp))
	err = utils.SendOTPTemplateEmail(
		[]string{email},
		"quangdvn@gmail.com",
		"otp-auth.html",
		map[string]interface{}{
			"otp": strconv.Itoa(otp),
		},
	)
	if err != nil {
		return responses.ErrCodeFailedEmail
	}
	// 5. Check OTP is exist

	// 6. Check spamming

	return responses.ErrCodeSuccess
}
