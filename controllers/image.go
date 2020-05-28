/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-16 17:28:34
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-17 15:39:43
 */

package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/gin-gonic/gin"
)

var imageBasePath = "/webroot/oss/image"
var imageRE       = regexp.MustCompile(`^image/`)
var pathFormatRE  = regexp.MustCompile(`^(.{2})(.{2})`)

func init() {
	if runtime.GOOS == "windows" {
		imageBasePath = "C:" + imageBasePath
	}
}

// ImageUploadHandle 上传图片
func ImageUploadHandle(c *gin.Context) {
	userID, errStr := getUserID(c)
	if userID == 0 {
		cJSONUnauthorized(c, errStr)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	if file.Size > 1 << 20 {
		cJSONBadRequest(c, "不能大于1M")
		return
	}

	imageType := file.Header["Content-Type"][0]

	if imageRE.Match([]byte(imageType)) != true {
		cJSONBadRequest(c, "请上传图片文件")
		return
	}

	f, _ := file.Open()
	defer f.Close()

	MD5 := md5.New()
	io.Copy(MD5, f)
	MD5Str := hex.EncodeToString(MD5.Sum(nil))

	fileName := pathFormatRE.ReplaceAllString(MD5Str, "$1/$2/")
	fileName = imageRE.ReplaceAllString(imageType, fileName + ".$1")

	path := filepath.Join(imageBasePath, fileName[:6])
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	dst := filepath.Join(imageBasePath, fileName)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	cJSONOk(c, fileName)
}

// GetImageHandle 获取图片
func GetImageHandle(c *gin.Context) {
	c.File(filepath.Join(imageBasePath, c.Param("name")))
}
