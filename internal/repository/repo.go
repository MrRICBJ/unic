package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strconv"
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
	q := `INSERT INTO users (name, password, role) VALUES ($1, $2, $3) RETURNING id`

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

func (r *Repo) GetUsers() ([]entity.User, error) {
	q := `Select * From users`

	users := make([]entity.User, 0)

	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.Role, pq.Array(&user.ResultTest), pq.Array(&user.Students))
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *Repo) GetResultTest(id int64) ([]string, error) {
	q := `Select result_test From users WHERE id = $1`

	res := make([]string, 0)

	rows, err := r.db.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tmp := ""
		err = rows.Scan(&tmp)
		if err != nil {
			return nil, err
		}
		res = append(res, tmp+"/13")
	}
	return res, nil
}

func (r *Repo) CreateResultTest(count int, id int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	q := `select result_test from users where id = $1`

	var resTest []string
	err = tx.QueryRow(q, id).Scan(&resTest)
	if err != nil {
		return err
	}

	resTest = append(resTest, strconv.Itoa(count)+"/13")
	q = `UPDATE users SET result_test = $1 WHERE id = $2`
	_, err = tx.Exec(q, resTest, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
