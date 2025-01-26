package repositories

type UserRepo struct {}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetUserInfo() string {
	return "user"
}