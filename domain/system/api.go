package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"servermanager/conf"
	"servermanager/middleware"
	"servermanager/utils"
)

type tmp struct {
	Password string
	Xmguid   string
}

func Login(gc *gin.Context) {

	res := utils.NewResponseBodyModel()
	var tt tmp
	err := gc.BindJSON(&tt)
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		gc.Abort()
		return
	}
	if tt.Password != conf.Config.App.Pwd {
		res.Message = "密码不正确"
		gc.JSON(http.StatusOK, res)
		gc.Abort()
		return
	}

	//写入Session
	middleware.SaveManagerSession(gc)

	res.Message = "登陆成功"
	res.Status = http.StatusOK

	gc.JSON(200, res)
}
