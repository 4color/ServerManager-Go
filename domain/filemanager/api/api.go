package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"servermanager/domain/filemanager/model"
	"servermanager/domain/filemanager/service"
	"servermanager/globel"
	"servermanager/utils"
	"servermanager/utils/utilsIo"
	"servermanager/utils/utilsShell"
	"strings"
)

const (
	fileName = "./static/java.json"
)

type FileMangerApi struct {
}

func (p *FileMangerApi) GetList(gc *gin.Context) {

	res := utils.NewResponseBodyModel()

	bfile, err := utilsIo.PathExists(fileName)
	if err != nil {
		res.Message = err.Error()
		return
	}

	if !bfile {

		fo, err := os.Create(fileName)
		defer fo.Close()

		if err != nil {
			res.Message = err.Error()
			return
		}
		fo.WriteString("[]")
		fo.Sync()
	}

	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	//str := string(b)

	list := []model.JavaEntity{}
	json.Unmarshal(b, &list)

	for i := 0; i < len(list); i++ {
		list[i].Status, _ = utilsShell.GetProcStatus(list[i].Path)
	}

	res.Data = list
	res.Message = "获取成功"
	res.Status = http.StatusOK

	gc.JSON(200, res)
}

func (p *FileMangerApi) Save(gc *gin.Context) {

	res := utils.NewResponseBodyModel()

	param := model.JavaEntity{}
	err := gc.ShouldBind(&param)

	if err != nil {
		res.Message = err.Error()
		return
	}
	if param.Path == "" || param.Name == "" {
		res.Message = "请填写必填项"
		return
	}
	edit := true
	if param.Id == "" {
		edit = false
		param.Id = utils.Guid()
	}

	param.Time = utils.Time{}.Now().String()

	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	list := []model.JavaEntity{}
	json.Unmarshal(b, &list)

	for i := 0; i < len(list); i++ {
		if list[i].Id == param.Id {

			list[i].Name = param.Name
			list[i].Path = param.Path
			list[i].Time = param.Time
			list[i].Vars = param.Vars
			break
		}
	}
	if edit == false {
		list = append(list, param)
	}

	//替换内容

	fo, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0666) // just pass the file name
	defer fo.Close()
	if err != nil {
		fmt.Print(err)
	}

	bstr, _ := json.Marshal(list)

	fo.WriteString(string((bstr)))

	res.Data = list
	res.Message = "更新成功"
	res.Status = http.StatusOK

	gc.JSON(200, res)
}

func (p *FileMangerApi) Delete(gc *gin.Context) {

	res := utils.NewResponseBodyModel()

	id := gc.Param("id")

	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	list := []model.JavaEntity{}
	json.Unmarshal(b, &list)

	for i := 0; i < len(list); i++ {
		if list[i].Id == id {
			list = append(list[:i], list[i+1:]...)
			break
		}
	}
	//替换内容

	fo, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0666) // just pass the file name
	defer fo.Close()
	if err != nil {
		fmt.Print(err)
	}

	bstr, _ := json.Marshal(list)

	fo.WriteString(string((bstr)))

	res.Data = list
	res.Message = "删除成功"
	res.Status = http.StatusOK

	gc.JSON(200, res)
}

func (p *FileMangerApi) Stop(gc *gin.Context) {

	res := utils.NewResponseBodyModel()

	id := gc.Param("id")

	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	list := []model.JavaEntity{}
	json.Unmarshal(b, &list)

	path := ""
	for i := 0; i < len(list); i++ {
		if list[i].Id == id {

			path = list[i].Path
			break
		}
	}

	if path != "" {
		_, port := utilsShell.GetProcStatus(path)
		if port != "" {
			result := utilsShell.StopProc(port)
			res.Message = "返回消息：" + result
			res.Status = http.StatusOK
		} else {
			res.Message = "程序未启动"
		}
	} else {
		res.Message = "程序不存在"
	}

	gc.JSON(200, res)
}

func (p *FileMangerApi) Start(gc *gin.Context) {

	res := utils.NewResponseBodyModel()

	id := gc.Param("id")

	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	list := []model.JavaEntity{}
	json.Unmarshal(b, &list)

	mode := model.JavaEntity{}
	path := ""
	for i := 0; i < len(list); i++ {
		if list[i].Id == id {

			path = list[i].Path
			mode = list[i]
			break
		}
	}

	if path == "" {
		res.Message = "程序不存在"
		gc.JSON(200, res)
		return
	}

	if mode.Vars == "" {
		res.Message = "启动参数不能为空"
		gc.JSON(200, res)
		return
	}

	_, port := utilsShell.GetProcStatus(path)
	if port != "" {
		res.Message = "程序正在运行不能启动"
		gc.JSON(200, res)
		return
	}

	dirs := strings.LastIndex(path, "/")
	dir := path[0:dirs]

	oldpath, _ := os.Getwd()

	println("工作目录：" + oldpath)
	//更改工作目录
	os.Chdir(dir)

	newpath, _ := os.Getwd()

	println("新工作目录：" + newpath)

	if strings.Contains(mode.Vars, "nohup") {
		utilsShell.ExecCommandNotWaite(mode.Vars)
	} else {
		utilsShell.ExecCommandnohup(mode.Vars)
	}
	res.Message = "已调用命令"
	res.Status = 200

	//切换回工作目录
	os.Chdir(oldpath)

	gc.JSON(200, res)
}

func (p *FileMangerApi) ReadLog(gc *gin.Context) {

	res := utils.NewResponseBodyModel()

	id := gc.Param("id")

	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	list := []model.JavaEntity{}
	json.Unmarshal(b, &list)

	mode := model.JavaEntity{}
	path := ""
	for i := 0; i < len(list); i++ {
		if list[i].Id == id {

			path = list[i].Path
			mode = list[i]
			break
		}
	}

	if path == "" {
		res.Message = "程序不存在"
		gc.JSON(200, res)
		return
	}

	if mode.Vars == "" {
		res.Message = "启动参数不能为空"
		gc.JSON(200, res)
		return
	}

	dirs := strings.LastIndex(path, "/")
	if dirs == -1 {
		res.Message = "路径不正确"
		gc.JSON(200, res)
		return
	}
	dir := path[0:dirs]

	logIndex := strings.Split(mode.Vars, ">")
	if len(logIndex) < 2 {
		res.Message = "未指定日志存储位置"
		gc.JSON(200, res)
		return
	}
	log := strings.Replace(strings.Trim(logIndex[1], " "), "&", "", -1)

	logpath := dir + "/" + log

	st := service.ServiceTail{}
	//st.TailLog(logpath)
	//开启协助
	c := make(chan int)

	globel.TailFilesAdd(id, logpath, c)

	osname := runtime.GOOS

	if osname == globel.WINDOWS {
		go st.TailTest(id, c)
	}
	if osname == globel.LINUX {
		go st.TailLog(logpath, c)
	}

	res.Data = logpath
	res.Message = "读取日志" + logpath
	res.Status = 200
	gc.JSON(200, res)
	return
}

func (p *FileMangerApi) StopLog(gc *gin.Context) {

	res := utils.NewResponseBodyModel()

	id := gc.Param("id")

	ta := globel.TailFilesGet(id)

	osname := runtime.GOOS
	if osname == globel.LINUX {
		close(ta.Cid)
	}

	if osname == globel.WINDOWS {
		//停止协助的写法
		ta.Cid <- 1
		<-ta.Cid
	}

	globel.TailFilesDelet(id)

	res.Message = "已停止读取日志"
	res.Status = 200
	gc.JSON(200, res)
	return
}

var (
	uploadFileKey = "file"
)

//上传Jar包
func (p *FileMangerApi) UploadJar(gc *gin.Context) {

	res := utils.NewResponseBodyModel()

	header, err1 := gc.FormFile(uploadFileKey)
	if err1 != nil {
		//ignore
		res.Message = err1.Error()
		gc.JSON(http.StatusOK, res)
		gc.Abort()
		return
	}

	dir, err1 := filepath.Abs(filepath.Dir(os.Args[0]))
	if err1 != nil {
		//ignore
		res.Message = err1.Error()
		gc.JSON(http.StatusOK, res)
		gc.Abort()
		return
	}

	filename, err := utilsIo.Upload(gc, header, "jar")
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		gc.Abort()
		return
	}

	id := gc.Param("id")

	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	list := []model.JavaEntity{}
	json.Unmarshal(b, &list)

	path := ""
	for i := 0; i < len(list); i++ {
		if list[i].Id == id {
			path = list[i].Path
			break
		}
	}

	//覆盖文件
	resutl := utilsShell.CopyJar(dir+"/"+filename, path)

	res.Message = "执行结果:" + resutl
	res.Status = http.StatusOK
	gc.JSON(200, res)
}
