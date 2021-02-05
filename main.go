package main

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/lib/pq"
	_userRepo "gitlab.com/alfred_soegiarto/training-clean-arch/user/repository"
)

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

	dbConn, err := sql.Open("postgres", connString)
	if err != nil {
		fmt.Println("Failed to connect")
		panic(err)
	}

	userRepo := _userRepo.NewPostgreUserRepository(dbConn)
	fmt.Println(userRepo)
	fmt.Println("Successfully connected to Database!")
}
