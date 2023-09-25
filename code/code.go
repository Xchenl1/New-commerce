package code

const (
	Success     = 200
	Error       = 500
	InvalidCode = 400

	//user模块
	ErrorExistUser         = 30001 //用户已存在
	ErrorFailEncryption    = 30002 //加密失败
	ErrorExistUserNotFound = 30003 //找不到用户
	ErrorNotCompare        = 30004 //密码错误
	ErrorAuthToken         = 30005 //token验证失败
	ErrorImageBig          = 30006 //图片上传过大
	ErrorUploadFail        = 30007 //图片上传失败
	ErrorSendEmail         = 30008

	//商品模块
	ErrorProductImgUpload = 40001 //商品图片加载失败
	ErrorproductExist     = 40002 //商品已经存在
	ErrorProductNotExist  = 40003 //商品不存在
	ErrorBossNotExist     = 40004 //boss不存在

	//地址模块
	ErrorAddressNotFind = 50001
)
