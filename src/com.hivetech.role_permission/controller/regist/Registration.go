package regist

import (
	"pt_role_permission_go/src/com.hivetech.role_permission/config"
	"pt_role_permission_go/src/com.hivetech.role_permission/controller"
	"pt_role_permission_go/src/com.hivetech.role_permission/util"
)

// GET, POST, PUT, DELETE Registration

func PublicApiRegistration() {
	util.Engine.POST("/login", controller.Login)
	util.Engine.Group("/api/v1/public").
		GET("/ping", controller.Ping).
		GET("/kanto", controller.Kanto).
		GET("/user", controller.CreateUser)
}
func PrivateApiRegistration() {
	privateRouter := util.Engine.Group("/api/v1")
	privateRouter.Use(config.JwtFilter).
		//privateRouter.Use(config.PrivateACLCheck).
		GET("/:uid/:pid", config.PrivateApiTest).
		GET("/p2", config.PrivateApiTest)
}
