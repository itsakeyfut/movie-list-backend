package usecase

import (
	"backend/internal/domain/user"
)

type UserUsecase interface {
	Signup(u *user.User) (*user.User, error)
	Login(email, password string) (*user.User, error)
}

type userUsecase struct {
	userService user.UserService
}

func NewUserUsecase(us user.UserService) UserUsecase {
	return &userUsecase{userService: us}
}

func (uc *userUsecase) Signup(u *user.User) (*user.User, error) {
	return uc.userService.Signup(u)
}

func (uc *userUsecase) Login(email, password string) (*user.User, error) {
	return uc.userService.Login(email, password)
}
