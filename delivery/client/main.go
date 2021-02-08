package main

import (
	delivery "gitlab.com/alfred_soegiarto/training-clean-arch/delivery/gRPC"
	"gitlab.com/alfred_soegiarto/training-clean-arch/domain"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var mockUser = &domain.User{
	Id: "1",
	Name: "Alfred",
	Email: "test1@gmail.com",
	Password: "asdasd",
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	client := delivery.NewUserServiceClient(conn)

	testRequest := &delivery.FetchAllRequest{
	}
	response, err := client.FetchAll(context.Background(), testRequest)
	log.Println(response)

	emailRequest := &delivery.EmailRequest{
		Email: mockUser.Email,
	}
	response2, err := client.GetUserItem(context.Background(),emailRequest)
	log.Println(response2)

	userReq := &delivery.UserRequest{
		Id: 1,
	}
	response3, err := client.GetById(context.Background(),userReq)
	log.Println(response3)

	//updateReq := &delivery.UpdateRequest{
	//	UserId: 4,
	//	Name: "Bryan",
	//	Email: "emaildiganti@gmail.com",
	//	Password: mockUser.Password,
	//}
	//response4, err := client.Update(context.Background(),updateReq)
	//log.Println(response4)

	//insertReq := &delivery.InsertRequest{
	//	Name: "userBaru",
	//	Email: "emailBaru@gmail.com",
	//	Password: mockUser.Password,
	//}
	//response5, err := client.Insert(context.Background(),insertReq)
	//log.Println(response5)

	//deleteReq := &delivery.DeleteRequest{
	//	UserId: 6,
	//}
	//response6, err := client.Delete(context.Background(),deleteReq)
	//log.Println(response6)
}
