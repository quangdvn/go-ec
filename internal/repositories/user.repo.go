package repositories

import (
	"github.com/quangdvn/go-ec/global"
	"github.com/quangdvn/go-ec/internal/model"
)

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
	// SELECT * FROM user where email = '??' ORDER BY email
	row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	return row != NilNumber
}
