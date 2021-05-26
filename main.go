package main

import (
	"golang.org/x/sync/errgroup"
	"servermanager/conf"
	"servermanager/globel"
	"servermanager/router"
)

func main() {

	var (
		g errgroup.Group
	)

	//定时取Redis任务
	g.Go(func() error {
		return InitWeb()
	})

	//g.Go(func() error {
	//	return InitTail()
	//})

	if err := g.Wait(); err != nil {
		println(err)
	}



}

func InitWeb() (err error) {

	router.InitRouter()
	globel.R.Run(conf.Config.App.Server.Port)
	return

}

//func InitTail() (err error) {
//
//	st := service.ServiceTail{}
//	st.TailTest("")
//	return
//
//}
