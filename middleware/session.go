package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"servermanager/globel"
	"servermanager/utils"
	"time"
)

//保存管理员的会话
func SaveManagerSession(gc *gin.Context) {

	session := sessions.DefaultMany(gc, globel.SessionNameManager)
	session.Set(globel.SessionKeyManagerName, "App8303")
	session.Save()
}

//管理员的api认证
func AuthManagerSessionMiddleView() gin.HandlerFunc {
	return func(gc *gin.Context) {
		session := sessions.DefaultMany(gc, globel.SessionNameManager)
		sessionValue := session.Get(globel.SessionKeyManagerName)
		if sessionValue == nil {
			gc.Redirect(http.StatusFound, "/")
			gc.Abort()
			return
		}
		// 设置简单的变量
		gc.Set(globel.SessionKeyManagerName, sessionValue.(string))

		session.Set("time", utils.Time(time.Now()).String())
		session.Save()

		gc.Next()
		return
	}

	//return func(c *gin.Context) {
	//	c.Next()
	//	return
	//}
}

func AuthManagerSessionMiddleApi() gin.HandlerFunc {
	return func(gc *gin.Context) {
		session := sessions.DefaultMany(gc, globel.SessionNameManager)
		sessionValue := session.Get(globel.SessionKeyManagerName)
		if sessionValue == nil {
			gc.JSON(http.StatusUnauthorized, gin.H{
				"error": "会话已过期,请重新登陆",
			})
			return
		}
		// 设置简单的变量
		gc.Set(globel.SessionKeyManagerName, sessionValue.(string))

		session.Set("time", utils.Time(time.Now()).String())
		session.Save()

		gc.Next()
		return
	}

	//return func(c *gin.Context) {
	//	c.Next()
	//	return
	//}
}
