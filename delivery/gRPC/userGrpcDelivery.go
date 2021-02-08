package delivery

import (
	"errors"
	"fmt"
	"gitlab.com/alfred_soegiarto/training-clean-arch/domain"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type userGRPCDelivery struct {
	userUsecase domain.UserUsecase
	UnimplementedUserServiceServer
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
		Error: "Fetch User Successful",
	}, err
}

func (u *userGRPCDelivery) GetById(ctx context.Context, request *UserRequest) (*User, error) {
	response, err := u.userUsecase.GetById(request.GetId())
	if err!=nil{
		return &User{}, errors.New("Failed to get user by id")
	}
	id, err := strconv.ParseInt(response.Id, 10, 64)
	return &User{
		UserId:   id,
		Name:     response.Name,
		Email:    response.Email,
		Password: response.Password,
	}, err
}

func (u *userGRPCDelivery) GetUserItem(ctx context.Context, request *EmailRequest) (*FetchAllItemReply, error) {
	response, err := u.userUsecase.GetUserItem(request.GetEmail())
	var items []*Item
	if err != nil{
		return &FetchAllItemReply{
			Items: items,
			Error: err.Error(),
		}, err
	}
	for _, index := range response{
		itemId, _ := strconv.ParseInt(index.Id,10,64)
		item := &Item{
			ItemId: itemId,
			AppName: index.App_name,
			AppEmail: index.App_email,
			AppPassword: index.App_password,
		}
		items = append(items, item)
	}
	return &FetchAllItemReply{
		Items: items,
		Error: "Fetch Item Successful",
	},err
}

func (u *userGRPCDelivery) Update(ctx context.Context, request *UpdateRequest) (*GeneralReply, error) {
	err := u.userUsecase.Update(request.GetUserId(),
		&domain.NewUser{
		Name: request.GetName(),
		Email: request.GetEmail(),
		Password: request.GetPassword(),
	})
	if err != nil{
		return &GeneralReply{
			Error: "Failed to update user",
		},err
	}
	return &GeneralReply{
		Error: "Update User Successful",
	},err
}

func (u *userGRPCDelivery) Insert(ctx context.Context, request *InsertRequest) (*GeneralReply, error) {
	err := u.userUsecase.Insert(&domain.NewUser{
		Name: request.GetName(),
		Email: request.GetEmail(),
		Password: request.GetPassword(),
	})
	if err != nil{
		return &GeneralReply{
			Error: "Failed to insert user",
		},err
	}
	return &GeneralReply{
		Error: "Insert User Successful",
	},err
}

func (u *userGRPCDelivery) Delete(ctx context.Context, request *DeleteRequest) (*GeneralReply, error) {
	err := u.userUsecase.Delete(request.GetUserId())
	if err!=nil{
		return &GeneralReply{
			Error: "Failed to get user by id",
		}, err
	}
	return &GeneralReply{
		Error: "Successfully deleted the user by id",
	}, err
}

// Serve the gRPC service
func (u *userGRPCDelivery) Serve() error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("delivery: ")
		log.Fatal(err.Error())
		return err
	}

	grpcServer := grpc.NewServer()
	RegisterUserServiceServer(grpcServer, u)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("delivery: ")
		log.Fatal(err.Error())
		return err
	}
	return nil
}
