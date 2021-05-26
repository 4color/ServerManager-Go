package service

import (
	"fmt"
	"github.com/hpcloud/tail"
	"servermanager/globel"
	"servermanager/utils"
	"servermanager/utils/utilsShell"
	"strings"
	"time"
)

type ServiceTail struct {
}

func (p *ServiceTail) TailLog(logpath string, done chan int) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("协程被强制关闭", err)
		}
	}()

	println("启动协程,读取：" + logpath)

	if strings.Contains(logpath, "`") {

		paths := strings.Split(logpath, "`")
		rq := paths[1]
		println(rq)
		rqresult := utilsShell.ExecCommand(rq)

		rqresult = strings.Replace(rqresult, "\n", "", -1)
		logpath = paths[0] + rqresult + paths[2]
	}

	logpath = strings.Trim(logpath, " ")
	t, _ := tail.TailFile(logpath, tail.Config{Follow: true})

	time.Sleep(time.Second * 2)

	for line := range t.Lines {

		select {
		case <-done:
			fmt.Println("退出协程...")
			done <- 1
			return
		default:
		}

		globel.Melody.Broadcast([]byte(line.Text))
	}
}

func (p *ServiceTail) TailTest(logpath string, done chan int) {

	println("启动协程")
	time.Sleep(time.Second * 1)
	for {

		select {
		case <-done:
			fmt.Println("退出协程...")
			done <- 1
			return
		default:
		}
		globel.Melody.Broadcast([]byte(logpath + ":" + utils.Timenow()))
		time.Sleep(time.Second * 1)
	}
}
