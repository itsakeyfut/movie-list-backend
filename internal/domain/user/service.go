package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Signup(u *User) (*User, error)
	Login(email, password string) (*User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Signup(u *User) (*User, error) {
	existing, _ := s.repo.FindByEmail(u.Email)
	if existing != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashedPassword)

	return s.repo.Create(u)
}

func (s *userService) Login(email, password string) (*User, error) {
	u, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return u, nil
}
