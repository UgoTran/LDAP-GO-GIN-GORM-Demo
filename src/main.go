package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"pt_role_permission_go/src/com.hivetech.role_permission/config"
	"pt_role_permission_go/src/com.hivetech.role_permission/controller/regist"
	"pt_role_permission_go/src/com.hivetech.role_permission/util"
	"time"
)

var TokenRepo = make(map[string]time.Time)

func main() {
	regist.PrivateApiRegistration()
	regist.PublicApiRegistration()

	serverPort := fmt.Sprintf("%v", viper.Get("SERVER_PORT"))
	err := util.Engine.Run(serverPort)
	log.Printf("Gin server stared on port " + serverPort)
	if err != nil {
		log.Fatal("[Error] failed to start Gin server due to: " + err.Error())
		return
	}
}

func init() {
	/*		Config logger global	*/
	logFileName := "app_rp.log"
	logFile, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("ERROR -> Failed to create logfile: " + logFileName)
		fmt.Println(logFile.Name())
		panic(err)
	}
	//defer logFile.Close()

	logrus.SetOutput(io.MultiWriter(logFile, os.Stdout))
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	config.InitEnvConfig()
	config.DBConnect()
}
