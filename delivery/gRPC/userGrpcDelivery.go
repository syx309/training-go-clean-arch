package delivery

import (
	"gitlab.com/alfred_soegiarto/training-clean-arch/domain"
	"golang.org/x/net/context"
	"strconv"
)

type Server struct {}
type userGRPCDelivery struct {
	userUsecase domain.UserUsecase
}

func NewUserGRPCDelivery(uc domain.UserUsecase) *userGRPCDelivery{
	return &userGRPCDelivery{
		userUsecase: uc,
	}
}

func (u *userGRPCDelivery) FetchAll(ctx context.Context, _ *FetchAllRequest) (*FetchAllReply, error) {
	res, err := u.userUsecase.FetchAll()
	var results []*User
	if err != nil{
		return &FetchAllReply{
			Users: results,
			Error: err.Error(),
		}, err
	}
	for _, item := range res{
		userId, _ := strconv.ParseInt(item.Id,10,64)
		user := User{
			UserId:   userId,
			Name:     item.Name,
			Email:    item.Email,
			Password: item.Password,
		}
		results = append(results, &user)
	}
	return &FetchAllReply{
		Users: results,
		Error: "No errors",
	}, err
}
