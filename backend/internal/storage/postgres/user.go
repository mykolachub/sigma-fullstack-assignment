package postgres

import (
	"database/sql"
	"fmt"
	"sigma-test/internal/entity"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (s *UsersRepo) CreateUser(data entity.User) (entity.User, error) {
	user := entity.User{}
	now := time.Now().UTC()

	query := "INSERT INTO Users(email, password, role, created_at) VALUES($1, $2, $3, $4) RETURNING *"
	rows := s.db.QueryRow(query, data.Email, data.Password, data.Role, now)
	err := rows.Scan(&user.ID, &user.Email, &user.Role, &user.Password, &user.CreatedAt)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s *UsersRepo) GetUsers() ([]entity.User, error) {
	query := "SELECT * FROM users"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	users := []entity.User{}
	for rows.Next() {
		user := entity.User{}
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *UsersRepo) GetUser(id string) (entity.User, error) {
	user := entity.User{}

	query := "SELECT * FROM users WHERE id = $1"
	err := s.db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.CreatedAt)

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s *UsersRepo) GetUserByEmail(email string) (entity.User, error) {
	user := entity.User{}

	query := "SELECT * FROM users WHERE email = $1"
	err := s.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.CreatedAt)

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s *UsersRepo) UpdateUser(id string, data entity.User) (entity.User, error) {
	user := entity.User{}
	updates := []string{}

	if data.Email != "" {
		updates = append(updates, fmt.Sprintf("email = '%s'", data.Email))
	}
	if data.Password != "" {
		updates = append(updates, fmt.Sprintf("password = '%s'", data.Password))
	}
	if data.Role != "" {
		updates = append(updates, fmt.Sprintf("role = '%s'", data.Role))
	}

	fmt.Println(data, updates)

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $1 RETURNING *", strings.Join(updates, ", "))
	fmt.Println(id)
	err := s.db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.CreatedAt)

	if err != nil {
		return entity.User{}, nil
	}

	return user, nil
}

func (s *UsersRepo) DeleteUser(id string) (entity.User, error) {
	user := entity.User{}

	query := "DELETE FROM users WHERE id = $1 RETURNING *"
	err := s.db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.CreatedAt)

	if err != nil {
		return entity.User{}, nil
	}

	return user, nil
}
