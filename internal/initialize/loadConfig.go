package initialize

import (
	"fmt"

	"github.com/quangdvn/go-ec/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config %w", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("failed to unmarshal config %w", err))
	}

	// fmt.Println("Server port::", config.Server.Port)

	// for _, db := range config.Database {
	// 	fmt.Printf("DB user:: %s, password:: %s, host:: %s \n", db.User, db.Password, db.Host)
	// }
}
