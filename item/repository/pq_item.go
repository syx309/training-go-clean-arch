package repository

import (
	"database/sql"
	"fmt"
	"gitlab.com/alfred_soegiarto/training-clean-arch/domain"
)

type postgreItemRepo struct {
	DB *sql.DB
}

// implements ItemRepository
func NewPostgreItemRepository(db *sql.DB) domain.ItemRepository{
	return &postgreItemRepo{db}
}

func (p *postgreItemRepo) FetchAll(user_id int64) (res []domain.Item, err error) {
	var item domain.Item
	rows, err := p.DB.Query("SELECT id, user_id, app_name, app_email, app_password FROM items WHERE $1", user_id)
	if err != nil {
		fmt.Println("Query error")
		panic(err)
	}
	defer rows.Close()

	var items []domain.Item
	for rows.Next() {
		err = rows.Scan(&item.Id, &item.User_id, &item.App_name, &item.App_email, &item.App_password)
		if err != nil {
			panic(err)
		}
		items = append(items, item)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return items, err
}

func (p *postgreItemRepo) GetByName(user_id int64, appName string) (i domain.Item, err error) {

	query := `SELECT items.id, user_id, app_name, app_email, app_password 
								FROM items JOIN users 
								ON items.user_id = users.id
								WHERE users.id = $1 
								AND LOWER(items.app_name) = LOWER($2)`
	row := p.DB.QueryRow(query, user_id, appName)

	var item domain.Item
	err = row.Scan(&item.Id, &item.User_id, &item.App_name, &item.App_email, &item.App_password)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			panic(err)
		} else {
			panic(err)
		}
	}
	return item, err
}

