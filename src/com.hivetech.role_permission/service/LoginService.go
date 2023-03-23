package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	. "pt_role_permission_go/src/com.hivetech.role_permission/config"
	"pt_role_permission_go/src/com.hivetech.role_permission/storage/entity"
	"pt_role_permission_go/src/com.hivetech.role_permission/util"
)

func LoginChecker(context *gin.Context, credentials util.LoginRequest) {
	var rpUser entity.RpUser
	// 1 check user on DB
	DBConnect().First(&rpUser, "username=?", credentials.Username)
	errPwHashed := bcrypt.CompareHashAndPassword([]byte(rpUser.Password), []byte(credentials.Password))
	if rpUser.Username != credentials.Username || errPwHashed != nil {
		context.JSON(http.StatusNotFound, UnsignedResponse{
			Message: "bad credentials",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, BuildClaims(rpUser))
	tokenStr, err := token.SignedString(SecretKey)
	if err != nil {
		context.JSON(http.StatusInternalServerError, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, LoginResponse{
		Token:       tokenStr,
		Message:     "logged in",
		Application: credentials.Application,
		FullName:    rpUser.FullName,
	})
}
