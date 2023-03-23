package controller

import "C"
import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"pt_role_permission_go/src/com.hivetech.role_permission/util"
)

var engine = util.Engine

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

// Kanto Gọi API từ nguồn khác, dùng con trỏ context *gin
//
// The * goes in front of a variable that holds a memory address and resolves it (it is therefore the counterpart to the & operator)
//
// The & Operator: goes in front of a variable when you want to get that variable's memory address.
// Deprecated: test unused func
func Kanto(c *gin.Context) {
	// dạng khai báo trực tiếp
	response, error := http.Get("https://reqres.in/api/users?page=1")

	if error != nil {
		fmt.Print(error.Error())
		return
	}
	// đọc gán vào byte[]
	//responseData, error := io.ReadAll(response.Body)
	//if error != nil {
	//	log.Fatal(error)
	//}
	// tạo 1 map string : interface any
	//var data = make(map[string]interface{})
	// parse byte[] thành map
	//error = json.Unmarshal(responseData, &data)
	//if error != nil {
	//	context.JSON(500, gin.H{"error": " call https://reqres.in/api/users?page=1 error"})
	//	return
	//}
	c.PureJSON(200, toMap(response))

}

func CreateUser(c *gin.Context) {
	var payloadMap = map[string]string{
		"name": "morpheus",
		"job":  "leader",
	}
	var payloadJson, error = json.Marshal(payloadMap)
	log.Println(error)

	response, errorHttp := http.Post(
		"https://reqres.in/api/users",
		"application/json",
		bytes.NewBuffer(payloadJson))
	if error != nil {
		log.Print(errorHttp.Error())
		return
	}

	c.PureJSON(200, toMap(response))
}

// Convert http.Response to map
func toMap(response *http.Response) map[string]interface{} {
	responseData, error := io.ReadAll(response.Body)
	if error != nil {
		log.Fatal(error)
		return map[string]interface{}{} // empty map
	}

	var data = make(map[string]interface{})
	json.Unmarshal(responseData, &data) // parse byte[] thành map
	return data
}
