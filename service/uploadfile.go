package service

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadAvatarToLocalStatic(file multipart.File, userid int64, name string) (filepath string, err error) {
	bId := strconv.Itoa(int(userid))
	basePath := "./Image/user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + name + ".jpg"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		logx.Error("读取文件失败！", err)
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		logx.Error("写文件错误err:", err)
		return "", err
	}
	return fmt.Sprintf("user%s/%s.jpg", bId, name), err
}

func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func CreateDir(dirname string) bool {
	if err := os.MkdirAll(dirname, 755); err != nil {
		return false
	}
	return true
}

func UploadProductToLocalStatic(file multipart.File, userid int64, name string) (filepath string, err error) {
	bId := strconv.Itoa(int(userid))
	basePath := "./Image/boss" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + name + ".jpg"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		logx.Error("读取文件失败！", err)
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		logx.Error("写文件错误err:", err)
		return "", err
	}
	return fmt.Sprintf("boss%s/%s.jpg", bId, name), err
}
