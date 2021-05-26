package utilsIo

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"path"
	"servermanager/utils"
	"strconv"
	"strings"
	"time"
)

func Upload(gc *gin.Context, header *multipart.FileHeader, NeedExt string) (uploadfile string, err error) {

	//创建文件夹
	year := time.Now().Year()
	month := int(time.Now().Month()) //time.Now().Month().String()
	day := time.Now().Day()

	var dir = "upload/" + strconv.Itoa(year) + "/" + strconv.Itoa(month) + "/" + strconv.Itoa(day)
	existPath, _ := PathExists(dir)
	if existPath == false {
		os.MkdirAll(dir, os.ModePerm)
	}

	//随机文件名
	var ext = strings.ToLower(path.Ext(header.Filename))

	if ext != ".jar" && ext != ".gif" && ext != ".xls" && ext != ".doc" && ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".pdf" && ext != ".docx" && ext != ".xlsx" && ext != ".csv" && ext != ".svg" {
		err = utils.NewError(500, "上传文件类型不允许")
		return
	}

	if NeedExt != "" && ext != "."+NeedExt {
		err = utils.NewError(500, "上传文件类型不允许,只能是"+NeedExt)
		return
	}

	uploadfile = dir + "/upload" + utils.Guid() + ext

	if err1 := gc.SaveUploadedFile(header, uploadfile); err1 != nil {
		err = utils.NewError(500, err1.Error())
		return
	}

	return

}
