package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	_itemRepo "gitlab.com/alfred_soegiarto/training-clean-arch/item/repository"
	_userRepo "gitlab.com/alfred_soegiarto/training-clean-arch/user/repository"
	_userUsecase "gitlab.com/alfred_soegiarto/training-clean-arch/user/usecase"
	_userDelivery "gitlab.com/alfred_soegiarto/training-clean-arch/delivery/gRPC"
)

var DB *sql.DB

func main(){
	viper.SetConfigFile(`config.yml`)
	err := viper.ReadInConfig()

	//default if not found
	viper.SetDefault("database.dbname", "test_db")

	host     := viper.GetString("database.host")
	port     := viper.GetInt("database.port")
	user     := viper.GetString("database.user")
	password := viper.GetString("database.password")
	dbname   := viper.GetString("database.dbname")

	if err != nil {
		panic(err)
	}
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	InitDatabase(connString)
	defer CloseDatabase()

	userRepo := _userRepo.NewPostgreUserRepository(DB)
	itemRepo := _itemRepo.NewPostgreItemRepository(DB)
	uUsecase := _userUsecase.NewUserUsecase(userRepo, itemRepo)
	server := _userDelivery.NewUserGRPCDelivery(uUsecase)

	err = server.Serve()
	if err != nil{
		panic(err.Error())
	}
	fmt.Println("Application running")

}

func InitDatabase(connString string) {
	if DB == nil {
		DB = openDatabase(connString)
	}
}

func openDatabase(connString string) *sql.DB {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		fmt.Println("Failed to connect")
		panic(err)
	}
	fmt.Println("Successfully connected to Database!")
	return db
}

func CloseDatabase() {
	err := DB.Close()
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
