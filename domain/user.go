package domain

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email 	 string `json:"email"`
	Password string `json:"password"`
}

type NewUser struct{
	Name     string `json:"name"`
	Email 	 string `json:"email"`
	Password string `json:"password"`
}

type UserUsecase interface {
	FetchAll() (users []User, err error)
	GetById(id int64) (user User, err error)
	GetUserItem(email string) (item []Item, err error)
	Update(id int64, user *NewUser) error
	Insert(user *NewUser) error
	Delete(id int64) error
}

type UserRepository interface {
	FetchAll() (res []User, err error)
	GetById(id int64) (user User, err error)
	GetUserByEmail(email string) (u User, err error)
	Update(id int64, user *NewUser) error
	Insert(user *NewUser) error
	Delete(id int64) error
}
