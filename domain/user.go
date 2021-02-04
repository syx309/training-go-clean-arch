package domain

import "context"

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
	GetById(ctx context.Context, id int64) (User, err error)
	GetUserItem(ctx context.Context, email string) (Item, err error)
	Update(ctx context.Context, u *User) error
	Insert(ctx context.Context, u *User) error
	Delete(ctx context.Context, u *User) error
}

type UserRepository interface {
	FetchAll() (res []UserOutput, err error)
	GetById(ctx context.Context, id int64) (User, err error)
	GetUserItem(ctx context.Context, email string) (Item, err error)
	Update(ctx context.Context, u *User) error
	Insert(ctx context.Context, u *User) error
	Delete(ctx context.Context, u *User) error
}
