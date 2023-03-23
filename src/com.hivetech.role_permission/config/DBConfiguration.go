package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func DBConnect() *gorm.DB {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil
	}
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.Get("DB_UN"),
		viper.Get("DB_PW"),
		viper.Get("DB_PORT"),
		viper.Get("DB_NAME"))
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})
	if error != nil {
		logrus.Fatal("Can't connect DB > STOP")
	}

	return db
}
