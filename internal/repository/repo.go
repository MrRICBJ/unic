package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"university/internal/entity"
)

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetTheory() (string, error) {
	q := `Select theory From data`

	theory := ""
	err := r.db.QueryRow(q).Scan(&theory)
	if err != nil {
		return "", err
	}
	return theory, nil
}

func (r *Repo) CreateUser(user *entity.User) (*entity.User, error) {
	q := `INSERT INTO users (name, password, role) VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.db.QueryRow(q, user.Name, user.Password, user.Role).Scan(&user.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repo) GetUser(name, password string) (*entity.User, error) {
	q := `Select * From users WHERE name = $1 AND password = $2`

	var user entity.User

	if err := r.db.QueryRow(q, name, password).
		Scan(&user.Id, &user.Name, &user.Password, &user.Role, pq.Array(&user.ResultTest), pq.Array(&user.Students)); err != nil {
		return nil, err
	}

	return &user, nil
}
