package domain

// User Entity struct
type User struct {
	ID       int
	Username string
	Password string
}

// UserRepository interface
type UserRepository interface {
	Create(user *User) error
	FindByUsername(username string) (*User, error)
}

// UserUsecase interface
type UserUsecase interface {
	Create(user *User) error
	FindByUsername(username string) (*User, error)
}
