package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"servermanager/conf"
	"servermanager/globel"
	"servermanager/middleware"
)

func InitRouter() {
	globel.R = gin.Default()
	//设置模式
	gin.SetMode(conf.Config.App.Mode)

	// 允许使用跨域请求  全局中间件
	globel.R.Use(middleware.Cors())

	//自定义分割符
	globel.R.Delims("{[{", "}]}")

	globel.R.Static("/static", "static")
	globel.R.Static("/upload", "upload")
	globel.R.LoadHTMLGlob("templates/**/*")

	store := cookie.NewStore([]byte("secret"))

	//设置会话时间
	store.Options(sessions.Options{
		MaxAge: int(20 * 60), //30min
		Path:   "/",
	})
	sessionNames := []string{globel.SessionNameManager}
	globel.R.Use(sessions.SessionsMany(sessionNames, store))

	RouterManager()

}
