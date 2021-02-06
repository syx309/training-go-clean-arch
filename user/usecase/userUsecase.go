package usecase

import (
	"errors"
	"fmt"
	"gitlab.com/alfred_soegiarto/training-clean-arch/domain"
	"log"
	"strconv"
)

type userUsecase struct{
	userRepo domain.UserRepository
	itemRepo domain.ItemRepository
}

func NewUserUsecase(ur domain.UserRepository, ir domain.ItemRepository) domain.UserUsecase{
	return &userUsecase{
		userRepo: ur,
		itemRepo: ir,
	}
}

func (u *userUsecase) FetchAll() (res []domain.User, err error) {
	result, err := u.userRepo.FetchAll()
	if err != nil {
		log.Fatal("uc: error fetching all user data")
		return result, err
	}
	for _, user := range result{
		fmt.Println(user)
	}
	return result, nil
}

func (u *userUsecase) GetById(id int64) (user domain.User, err error) {
	result, err := u.userRepo.GetById(id)
	if err != nil {
		log.Fatal("uc: error fetching user data by id")
		return result, err
	}
	fmt.Println(result)
	return result, nil
}

func (u *userUsecase) GetUserItem(email string) (item []domain.Item, err error) {
	if email == ""{
		return []domain.Item{}, errors.New("invalid email")
	}
	curUser, err := u.userRepo.GetUserByEmail(email)
	id, err := strconv.ParseInt(curUser.Id,10,64)
	result, err := u.itemRepo.FetchAll(id)
	if err != nil {
		log.Fatal("error fetching all user's item data")
		return result, err
	}
	for _, item := range result{
		fmt.Println(item)
	}
	return result, nil
}

func (u *userUsecase) Update(id int64, user *domain.NewUser) error {
	if id < 0{
		return errors.New("Invalid ID")
	}
	curData, err := u.GetById(id)
	if err != nil{
		return err
	}
	if curData.Name == user.Name && curData.Email == user.Email && curData.Password == user.Password{
		return errors.New("No changes made")
	}
	return u.userRepo.Update(id,user)
}

func (u *userUsecase) Insert(user *domain.NewUser) error {
	if user.Name == "" || user.Email == "" || user.Password == ""{
		return errors.New("Credentials can't be null")
	}
	return u.userRepo.Insert(user)
}

func (u *userUsecase) Delete(id int64) error {
	if id < 0{
		return errors.New("id is invalid")
	}
	return u.userRepo.Delete(id)
}
