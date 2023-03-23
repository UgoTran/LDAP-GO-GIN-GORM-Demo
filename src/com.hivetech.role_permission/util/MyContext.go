package util

import "github.com/gin-gonic/gin"

// Engine use context global
var Engine = gin.Default()

type LoginRequest struct {
	Username    string `json:"username,omitempty"`
	Password    string `json:"pw,omitempty"`
	Application string `json:"app,omitempty"`
}
