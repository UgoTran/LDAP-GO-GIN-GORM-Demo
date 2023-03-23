package config

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"net/http"
	"pt_role_permission_go/src/com.hivetech.role_permission/storage/entity"
	"strconv"
	"strings"
	"time"
)

var privateThings = map[string]map[int64]string{
	"mike": {
		0: "MIKE: private string",
		1: "MIKE: secret thing",
		2: "MIKE: sneaky secret",
	},
	"rama": {
		0: "RAMA: private string",
		1: "RAMA: secret thing",
		2: "RAMA: sneaky secret",
	},
}

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

type LoginResponse struct {
	Token       string `json:"token"`
	Message     string `json:"message"`
	Application string `json:"application"`
	FullName    string `json:"fullName"`
}

var SecretKey = []byte("sceret")

func index(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "index"})
}

func PrivateApiTest(c *gin.Context) {
	uidStr := c.Param("uid")
	pidInt, _ := strconv.ParseInt(c.Param("pid"), 10, 64)

	secret, ok := privateThings[uidStr][pidInt]

	if ok {
		c.JSON(200, gin.H{"msg": secret})
		return
	}
	c.JSON(200, gin.H{"msg": "unknown pid"})
}

func extractBearerToken(header string) (string, error) {
	jwtToken := strings.Split(header, " ")
	if header == "" || len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return SecretKey, nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func JwtFilter(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad token",
		})
		return
	}

	claims, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}

	logrus.Info("token claims: ", claims)
	c.Set("claims", "chungtest")
	c.Next()
}

func BuildClaims(user entity.RpUser) jwt.MapClaims {
	return jwt.MapClaims{
		"user":         user.Username,
		"expAt":        time.Now().Add(10 * time.Minute),
		"applications": []string{"pt", "hiro"},
	}
}

//
//func PrivateACLCheck(c *gin.Context) {
//	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
//			Message: err.Error(),
//		})
//		return
//	}
//
//	token, err := parseToken(jwtToken)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
//			Message: "bad jwt token",
//		})
//		return
//	}
//
//	claims, OK := token.Claims.(jwt.MapClaims)
//	if !OK {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
//			Message: "unable to parse claims",
//		})
//		return
//	}
//
//	claimedUID, OK := claims["user"].(string)
//	if !OK {
//		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
//			Message: "no user property in claims",
//		})
//		return
//	}
//
//	uid := c.Param("uid")
//	if claimedUID != uid {
//		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
//			Message: "token uid does not match resource uid",
//		})
//		return
//	}
//
//	c.Next()
//}
