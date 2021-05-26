package router

import (
	fileapi "servermanager/domain/filemanager/api"
	"servermanager/domain/filemanager/view"
	"servermanager/domain/system"
	"servermanager/globel"
	"servermanager/middleware"
)

func RouterManager() {

	//首页
	globel.R.GET("/", system.Index)
	globel.R.POST("/api/login", system.Login)

	ws := system.ApiWs{}
	globel.R.GET("/socket.io", ws.WsConnect)
	globel.R.GET("/wstest", ws.WebSokectTest)

	author := globel.R.Group("/", middleware.AuthManagerSessionMiddleView())
	{
		author.GET("/app", system.App)

		java := author.Group("/java")
		{
			java.GET("/list", view.List)
		}
	}

	api := globel.R.Group("/api", middleware.AuthManagerSessionMiddleApi())
	{

		fapi := fileapi.FileMangerApi{}
		api.POST("/java/list", fapi.GetList)
		api.POST("/java/save", fapi.Save)
		api.POST("/java/delete/:id", fapi.Delete)
		api.POST("/java/stop/:id", fapi.Stop)
		api.POST("/java/upload/:id", fapi.UploadJar)
		api.POST("/java/start/:id", fapi.Start)
		api.POST("/java/log/:id", fapi.ReadLog)
		api.POST("/java/stoplog/:id", fapi.StopLog)

	}
}
