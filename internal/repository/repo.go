package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strconv"
	"strings"
	"university/internal/entity"
)

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetTestAnswers() ([]string, error) {
	q := `Select test_a From data`

	testA := make([]string, 0)
	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tmp := ""
		err := rows.Scan(&tmp)
		if err != nil {
			return nil, err
		}

		testA = append(testA, tmp)
	}

	return testA, nil
}

func (r *Repo) GetTestQuestions() ([]string, error) {
	q := `Select test_q From data`

	testQ := make([]string, 0)
	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tmp := ""
		err := rows.Scan(&tmp)
		if err != nil {
			return nil, err
		}

		testQ = append(testQ, tmp)
	}

	return testQ, nil
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
	var est string
	var tmp *[]string
	var t bool
	err = tx.QueryRow(q, id).Scan(&tmp)
	if err != nil {
		_ = tx.QueryRow(q, id).Scan(&est)
		est = strings.ReplaceAll(est, "{", "")
		est = strings.ReplaceAll(est, "}", "")
		t = true
	}
	if t {
		resTest = append(resTest, est, strconv.Itoa(count)+"/13")
	} else {
		resTest = append(resTest, strconv.Itoa(count)+"/13")
	}
	q = `UPDATE users SET result_test = $1 WHERE id = $2`
	_, err = tx.Exec(q, pq.Array(resTest), id)
	//_, err = tx.Exec(q, fmt.Sprintf("{%s}", strings.Join(resTest, ",")), id)
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

func (r *Repo) AddStudent(idStudent int, id int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	q := `select students from users where id = $1`

	var student []string
	var tmp *[]string
	var est string
	var t bool

	err = tx.QueryRow(q, id).Scan(&tmp)
	if err != nil {
		_ = tx.QueryRow(q, id).Scan(&est)
		est = strings.ReplaceAll(est, "{", "")
		est = strings.ReplaceAll(est, "}", "")
		t = true
	}
	if t {
		student = append(student, est, strconv.Itoa(idStudent))
	} else {
		student = append(student, strconv.Itoa(idStudent))
	}

	q = `UPDATE users SET students = $1 WHERE id = $2`
	_, err = tx.Exec(q, pq.Array(student), id)
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

func (r *Repo) GetMyStudent(id int64) ([]entity.User, error) {
	q := `Select students From users WHERE id = $1`

	res := make([]entity.User, 0)
	str := ""

	err := r.db.QueryRow(q, id).Scan(&str)
	if err != nil {
		return nil, err
	}

	listStud := conv(str)

	for _, v := range listStud {
		q = `Select id, name, role, result_test From users WHERE id = $1`

		var user entity.User

		if err := r.db.QueryRow(q, v).
			Scan(&user.Id, &user.Name, &user.Role, pq.Array(&user.ResultTest)); err != nil {
			return nil, err
		}

		res = append(res, user)
	}
	return res, nil
}

func conv(str string) []int {
	str = strings.Trim(str, "{}")
	strValues := strings.Split(str, ",") // Разделяем элементы по запятой

	var intSlice []int
	for _, val := range strValues {
		num, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			fmt.Printf("Ошибка преобразования числа: %v\n", err)
			continue
		}
		intSlice = append(intSlice, num)
	}
	return intSlice
}
