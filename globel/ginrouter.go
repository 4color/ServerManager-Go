package globel

import "github.com/gin-gonic/gin"

var R *gin.Engine

//全局参数
type GVarStruct struct {
	Version string //版本
}

var Gvar *GVarStruct

func init() {

	var1 := GVarStruct{}
	var1.Version = "V1.8 Build 20210414"

	Gvar = &var1
}
