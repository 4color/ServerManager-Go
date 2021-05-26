package view

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//文件管理
func List(gc *gin.Context) {

	gc.HTML(http.StatusOK, "filemanager/index.html", gin.H{

	})
}
