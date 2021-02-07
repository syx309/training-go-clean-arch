package delivery

//// Serve gRPC
//func main() {
//	lis, err := net.Listen("tcp", ":8080")
//	if err != nil {
//		fmt.Println("delivery: ")
//		log.Fatal(err.Error())
//		return err
//	}
//
//	s := Server{}
//	grpcServer := grpc.NewServer()
//	s.RegisterUserDeliveryServer(grpcServer, u)
//
//	if err := grpcServer.Serve(lis); err != nil {
//		fmt.Println("delivery: ")
//		log.Fatal(err.Error())
//		return err
//	}
//	return nil
//}
