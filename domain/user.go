package domain

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email 	 string `json:"email"`
	Password string `json:"password"`
}

type UserOutput struct{
	Name     string `json:"name"`
	Email 	 string `json:"email"`
	Password string `json:"password"`
}

type UserUsecase interface {
	FetchAll() (res []UserOutput, err error)
	GetById(id int64) (User, err error)
	GetUserItem(email string) (Item, err error)
	Update(u *User) error
	Insert(u *User) error
	Delete(u *User) error
}

type UserRepository interface {
	FetchAll() (res []UserOutput, err error)
	GetById(id int64) (User, err error)
	GetUserItem(email string) (Item, err error)
	Update(u *User) error
	Insert(u *User) error
	Delete(u *User) error
}
