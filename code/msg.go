package code

var MsgFlags = map[int]string{
	Success:                "ok",
	Error:                  "fail",
	InvalidCode:            "参数错误",
	ErrorExistUser:         "用户已经存在",
	ErrorFailEncryption:    "加密失败",
	ErrorExistUserNotFound: "用户不存在",
	ErrorNotCompare:        "密码错误",
	ErrorAuthToken:         "token验证失败",
	ErrorImageBig:          "图片上传过大",
	ErrorUploadFail:        "图片上传失败",
	ErrorSendEmail:         "邮件发送失败",
	ErrorProductImgUpload:  "商品图片上传失败",
	ErrorproductExist:      "商品已经存在",
	ErrorProductNotExist:   "商品不存在",
	ErrorBossNotExist:      "boss不存在",
	ErrorAddressNotFind:    "用户地址不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg //成功直接返回
	}
	return MsgFlags[Error] //有误返回错误
}
