package repositories

// type UserRepo struct {}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetUserInfo() string {
// 	return "user"
// }

// !! WITH INTERFACE
type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
}

// NewUserRepository creates a new instance of IUserRepository
func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (ur *userRepository) GetUserByEmail(email string) bool {
	return true
}
