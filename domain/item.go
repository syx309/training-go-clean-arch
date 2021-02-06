package domain

type Item struct {
	Id       	 string `json:"id"`
	User_id		 string `json:"user_id"`
	App_name     string `json:"app_name"`
	App_email 	 string `json:"app_email"`
	App_password string `json:"app_password"`
}

type ItemRepository interface {
	FetchAll(user_id int64) (res []Item, err error)
	GetByName(user_id int64, appName string) (item Item, err error)
}
