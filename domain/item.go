package domain

import "context"

type Item struct {
	Id       	 string `json:"id"`
	User_id		 string `json:"user_id"`
	App_name     string `json:"app_name"`
	App_email 	 string `json:"app_email"`
	App_password string `json:"app_password"`
}

type ItemOutput struct {
	App_name     string `json:"app_name"`
	App_email 	 string `json:"app_email"`
	App_password string `json:"app_password"`
}

type ItemRepository interface {
	FetchAll() (res []ItemOutput, err error)
	GetById(ctx context.Context) (Item, error)
}
