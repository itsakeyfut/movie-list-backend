package user

type UserRepository interface {
	Create(u *User) (*User, error)
	FindByEmail(email string) (*User, error)
}
