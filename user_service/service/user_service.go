package service

import "shared/repository"

type UserService interface {
	Find() (any, error)
	FindById() (any, error)
	Create() (int64, error)
	Update() error
	Delete() error
}

type userService struct {
	repository.Repository
}

type User struct {
	Login    string
	Password string
	Id       int64
}

func (u *userService) Create(user User) (int64, error) {
	var id int64

	err := u.Get(&id, "INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id",
		user.Login,
		user.Password,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userService) FindById(id int64) (*User, error) {
	var user User

	err := u.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userService) Update(id int64, login, password string) error {
	_, err := u.Exec("UPDATE users SET login = $1, password = $2 WHERE id = $3",
		login,
		password,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) Find() (any, error) {
	panic("unimplemented")
}

// Delete implements UserService.
func (u *userService) Delete() error {
	panic("unimplemented")
}

func NewUserService(repo repository.Repository) UserService {
	return &userService{repo}
}
