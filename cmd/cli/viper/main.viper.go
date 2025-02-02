package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Dbname   string `mapstructure:"dbname"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config %w \n", err))
	}

	fmt.Println("Server port::", viper.GetInt("server.port"))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("Failed to unmarshal config %w \n", err))
	}

	fmt.Println("Server port::", config.Server.Port)

	for _, db := range config.Database {
		fmt.Printf("DB user:: %s, password:: %s, host:: %s \n", db.User, db.Password, db.Host)
	}
}
