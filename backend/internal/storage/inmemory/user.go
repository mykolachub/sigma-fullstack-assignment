package inmemory

import (
	"sigma-test/internal/entity"
)

// In case of gorm the struct will have db field
type UsersRepo struct {
	// db *gorm.DB
	Users []entity.User
}

func NewUsersRepo() *UsersRepo {
	// using simple slice of users as database
	users := []entity.User{
		{
			ID:       "2dXeeVItnr9Ci7trhCp10LFvtNJ",
			Email:    "user",
			Password: "$2a$10$9jrmat9.0TVc8nkgiSpUPu9ZHOrf7rHpq7csnkxUaESLFrzwQKtdO", // test123
			Role:     "user",
		},
		{
			ID:       "2dXecVszaUb8AUybj33pcb4lR60",
			Email:    "admin",
			Password: "$2a$10$9jrmat9.0TVc8nkgiSpUPu9ZHOrf7rHpq7csnkxUaESLFrzwQKtdO", // test123
			Role:     "admin",
		},
	}
	repo := UsersRepo{
		Users: users,
	}

	return &repo
}

func (r UsersRepo) GetUserIdxByEmail(email string) (int, bool) {
	for i, v := range r.Users {
		if v.Email == email {
			return i, true
		}
	}
	return 0, false
}

func (r UsersRepo) GetUserIdxById(id string) (int, bool) {
	for i, v := range r.Users {
		if v.ID == id {
			return i, true
		}
	}
	return 0, false
}

func (r *UsersRepo) AddUser(user entity.User) (entity.User, error) {
	r.Users = append(r.Users, user)
	return user, nil
}

func (r *UsersRepo) DeleteUser(idx int) error {
	r.Users = append(r.Users[:idx], r.Users[idx+1:]...)
	return nil
}

func (r *UsersRepo) GetUserByIdx(idx int) (entity.User, error) {
	return r.Users[idx], nil
}

func (r *UsersRepo) GetAllUsers() ([]entity.User, error) {
	return r.Users, nil
}

func (r *UsersRepo) UpdateUser(user entity.User, idx int) (entity.User, error) {
	r.Users[idx] = user
	return r.Users[idx], nil
}
