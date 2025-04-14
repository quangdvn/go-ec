package services

import "context"

type (
	//.. Define all the interfaces in one place by team leader
	IUserLogin interface {
		Login(ctx context.Context) error
		Register(ctx context.Context) error
		VerifyOtp(ctx context.Context) error
		UpdatePassword(ctx context.Context) error
	}
	IUserInfo interface {
		GetInfoByUserId(ctx context.Context) error
		GetAllUsers(ctx context.Context) error
	}
	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
		FindOneUser(ctx context.Context) error
	}
)

var (
	localUserLogin IUserLogin
	localUserAdmin IUserAdmin
	localUserInfo  IUserInfo
)

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("implement not found for IUserAdmin, forgot register?")
	}
	return localUserAdmin
}

func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("implement not found for IUserInfo, forgot register?")
	}
	return localUserInfo
}

func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}

func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("implement not found for IUserLogin, forgot register?")
	}
	return localUserLogin
}

func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}
