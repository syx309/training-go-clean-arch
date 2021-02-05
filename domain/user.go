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
	FetchAll() (res []User, err error)
	GetById(id int64) (u User, err error)
	GetUserItem(email string) (i Item, err error)
	Update(id int64, u *NewUser) error
	Insert(u *NewUser) error
	Delete(id int64) error
}

type UserRepository interface {
	FetchAll() (res []User, err error)
	GetById(id int64) (u User, err error)
	Update(id int64, u *NewUser) error
	Insert(u *NewUser) error
	Delete(id int64) error
}
