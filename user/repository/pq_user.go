package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"gitlab.com/alfred_soegiarto/training-clean-arch/domain"
)

type postgreUserRepo struct {
	DB *sql.DB
}

// implements UserRepository
func NewPostgreUserRepository(db *sql.DB) domain.UserRepository{
	return &postgreUserRepo{
		DB: db,
	}
}

func (p *postgreUserRepo) FetchAll() (res []domain.User, err error) {
	var user domain.User
	rows, err := p.DB.Query("SELECT name, email, password FROM users")
	if err != nil {
		fmt.Println("Query error")
		panic(err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		err = rows.Scan(&user.Name, &user.Email, &user.Password)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return users, err
}

func (p *postgreUserRepo) GetById(id int64) (u domain.User, err error) {
	query := fmt.Sprintf(`SELECT id, name, email, password FROM users WHERE id = %s`, id)
	row := p.DB.QueryRow(query)

	var user domain.User
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			panic(err)
		} else {
			panic(err)
		}
	}

	return user, err
}

func (p *postgreUserRepo) GetUserByEmail(email string) (u domain.User, err error) {
	query := fmt.Sprintf(`SELECT id, name, email, password FROM users WHERE LOWER(email) = LOWER('%s')`, email)
	row := p.DB.QueryRow(query)
	var user domain.User
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			panic(err)
		} else {
			panic(err)
		}
	}
	return user, err
}

func (p *postgreUserRepo) Update(id int64, u *domain.NewUser) error {
	if  id< 0{
		return errors.New("Invalid ID")
	}
	query := fmt.Sprintf("UPDATE users SET name = \"%s\", email = \"%s\", password = \"%s\" WHERE id = %s", u.Name, u.Email, u.Password, id)
	_, err := p.DB.Exec(query)
	if err != nil {
		panic(err)
	}
	return err
}

func (p *postgreUserRepo) Insert(u *domain.NewUser) error {
	query := fmt.Sprintf("INSERT INTO users (name, email, password) VALUES (\"%s\", \"%s\", \"%s\")", u.Name, u.Email, u.Password)
	_, err := p.DB.Exec(query)
	if err != nil {
		panic(err)
	}
	return err
}

func (p *postgreUserRepo) Delete(id int64) error {
	query := fmt.Sprintf("DELETE FROM users WHERE $1")
	_, err := p.DB.Exec(query,1)
	if err != nil {
		panic(err)
	}
	return err
}

