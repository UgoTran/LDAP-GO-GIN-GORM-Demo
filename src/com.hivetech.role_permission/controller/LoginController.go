package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "pt_role_permission_go/src/com.hivetech.role_permission/config"
	"pt_role_permission_go/src/com.hivetech.role_permission/service"
	"pt_role_permission_go/src/com.hivetech.role_permission/util"
)

func Login(context *gin.Context) {
	credentials := util.LoginRequest{}
	err := context.ShouldBindJSON(&credentials)
	if err != nil || util.AnyEmpty(credentials.Password, credentials.Username, credentials.Application) {
		context.JSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad credentials: cannot parse body json or missed credential info: pw, username, app",
		})
		return
	}
	service.LoginChecker(context, credentials)
}
