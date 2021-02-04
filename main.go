package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main(){
	viper.SetConfigFile(`config.yml`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	//default if not found
	viper.SetDefault("database.dbname", "test_db")

	fmt.Println("Database is\t", viper.GetString("database.dbname"))
	fmt.Println("Port is\t\t", viper.GetInt("server.port"))


}
