package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//首页
func Index(gc *gin.Context) {

	gc.HTML(http.StatusOK, "index/index.html", gin.H{

	})
}

//App
func App(gc *gin.Context) {

	gc.HTML(http.StatusOK, "index/app.html", gin.H{

	})
}
