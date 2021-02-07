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
}
