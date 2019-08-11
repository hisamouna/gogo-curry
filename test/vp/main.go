package main

import (
	"os"
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Data Data
}

type Data struct {
	Host string 
	Port string
	User string
	Password string
}

func readConfig() {
	var C Config

	viper.SetConfigName("config")

	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("config file read error %v\n",err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&C); err != nil {
		fmt.Printf("config file Unmarshal error %v\n",err)
		os.Exit(1)
	}

	fmt.Println(C)
	fmt.Println("Config: ",C.Data.Host)
}

func main() {
	readConfig()
}


