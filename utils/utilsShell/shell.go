package utilsShell

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"servermanager/utils"
	"strings"
)

func ExecCommand(strCommand string) string {
	cmd := exec.Command("/bin/bash", "-c", strCommand)

	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return ""
	}

	out_bytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return ""
	}
	return string(out_bytes)
}

func ExecCommandNotWaite(strCommand string) {
	cmd := exec.Command("/bin/bash", "-c", strCommand)
	cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return
	}
}

func ExecCommandnohup(strCommand string) {
	cmd := exec.Command("java", strCommand)
	cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return
	}
}

func ExecCommandStone(strCommand string) (result string, err error) {
	cmd := exec.Command(strCommand)

	stdout, _ := cmd.StdoutPipe()
	if err1 := cmd.Start(); err1 != nil {
		fmt.Println("Execute failed when Start:" + err1.Error())
		err = utils.NewErrorDefault(err1.Error())
		return
	}

	out_bytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err2 := cmd.Wait(); err2 != nil {
		fmt.Println("Execute failed when Wait:" + err2.Error())
		err = utils.NewErrorDefault(err2.Error())
		return
	}

	result = string(out_bytes)
	return
}

func GetProcStatus(path string) (status string, port string) {

	status = "无效"
	osname := runtime.GOOS
	//取path的最后一级
	paths := strings.Split(path, "/")
	name := paths[len(paths)-1]

	cmdResult := ""
	if osname == "linux" {
		cmdResult = ExecCommand("ps -aux | grep " + name)
		print(cmdResult)
	}
	if osname == "windows" {
		//cmdResult = "root     31118  0.0  0.0 113180  1220 pts/1    S+   17:39   0:00 /bin/bash -c ps -aux | grep reportservice.1.0.0-RELEASE.jar\nroot     31120  0.0  0.0 112728   956 pts/1    S+   17:39   0:00 grep reportservice.1.0.0-RELEASE.jar\nroot     48214  0.7  8.0 4067900 304524 ?      Sl   1月05 1477:43 java -jar -Xms50m -Xmx400m reportservice.1.0.0-RELEASE.jar --spring.profiles.active=dev\n"
	}
	lines := strings.Split(cmdResult, "\n")

	iIndex := -1
	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "java") && strings.Contains(lines[i], "-jar") {
			ports := strings.Split(lines[i], " ")
			for ii := 0; ii < len(ports); ii++ {
				if ports[ii] != "" {
					iIndex++
				}
				if iIndex == 1 {
					port = ports[ii]
					println(port)
					break
				}
			}
		}
	}

	if port != "" {

		status = "运行"

		//cmdResult = ""
		//if osname == "linux" {
		//	cmdResult = ExecCommand("ll /proc/" + port + "/cwd")
		//	println(cmdResult)
		//}
		//
		//if osname == "windows" {
		//	cmdResult = "lrwxrwxrwx. 1 root root 0 5月  20 17:53 /proc/48214/cwd -> /home/maque/run/reportToExcel"
		//}
		//
		//if cmdResult != "" {
		//	lines = strings.Split(cmdResult, "->")
		//	if len(lines) > 1 {
		//		path1 := lines[len(lines)-1]
		//
		//		if path1+"/"+name == path {
		//			status = "运行"
		//		}
		//	}
		//}
	}

	return
}

func StopProc(port string) (result string) {

	osname := runtime.GOOS

	if osname == "linux" {
		result = ExecCommand("kill -9 " + port)
		print(result)
	}

	return
}

func CopyJar(srcFileName string, dstFileName string) (result string) {

	result = "复制完成"

	println("source:" + srcFileName)
	println("target:" + dstFileName)

	if srcFileName == dstFileName {
		result = "源文件和目的文件名字不能相同"
		return
	}

	//只读方式打开源文件
	sF, err1 := os.Open(srcFileName)
	if err1 != nil {
		result = "err1 : " + err1.Error()
		return
	}

	//新建目的文件
	dF, err2 := os.Create(dstFileName)
	if err2 != nil {
		result = "er2 = " + err2.Error()
		return
	}

	//操作完毕，需要关闭文件
	defer sF.Close()
	defer dF.Close()

	//核心处理，从源文件读取内容，往目的文件写，读多少写多少
	buf := make([]byte, 4*1024) //4k大小临时缓冲区
	for {
		n, err := sF.Read(buf) //从源文件读取内容
		if err != nil {
			if err == io.EOF { //文件读取完毕
				break
			} else {
				result = "err = " + err.Error()
			}
		}
		//往目的文件写，读多少写多少
		dF.Write(buf[:n])
	}

	return
}
