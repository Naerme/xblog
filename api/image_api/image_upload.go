package image_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/utils"
	"blogx_server/utils/file"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
)

func (i ImageApi) ImageUploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	//大小判断
	s := global.Conifg.Upload.Size
	if fileHeader.Size > s*1024*1024 {
		fmt.Println(global.Conifg.Upload.Size)
		res.FailWithMsg(fmt.Sprintf("文件大小大于%dMB", s), c)
		return
	}
	//后缀判断
	filename := fileHeader.Filename
	suffix, err := file.ImageSuffixJudge(filename)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	//文件hash
	file, err := fileHeader.Open()
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	byteData, _ := io.ReadAll(file)
	hash := utils.Md5(byteData)
	var model models.ImageModel
	err = global.DB.Take(&model, "hash = ?", hash).Error
	if err == nil {
		logrus.Infof("上传图片重复%s <==> %s  %s", filename, model.Filename, hash)
		res.Ok(model.WebPath(), "上传成功", c)
		return
	}
	//文件名称一样，文件不一样直接用hash做文件名

	filepath := fmt.Sprintf("uploads/%s/%s.%s", global.Conifg.Upload.UploadDir, hash, suffix)
	//入库
	model = models.ImageModel{
		Filename: filename,
		Path:     filepath,
		Size:     fileHeader.Size,
		Hash:     hash,
	}
	err = global.DB.Create(&model).Error
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	c.SaveUploadedFile(fileHeader, filepath)
	res.Ok(model.WebPath(), "图片上传成功", c)
}

//func imageSuffixJudge(filename string) (suffix string, err error) {
//	_list := strings.Split(filename, ".")
//	if len(_list) == 1 {
//		err = errors.New("错误的文件名")
//		return
//	}
//	suffix = _list[len(_list)-1]
//	if !utils.InList(suffix, global.Conifg.Upload.WhiteList) {
//		err = errors.New("文件非法")
//		return
//	}
//	return
//}
