package encrypt

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
)

var Encrypt *Encryption

type Encryption struct {
	Key string
}

func init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *Encryption {
	return &Encryption{}
}

// PadPwd 填充密码长度
func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

// AesEncoding 加密  采用的是16字节  128位 电码本模式 ECB（Electronic Codebook Book）
func (k *Encryption) AesEncoding(src string) string {
	srcByte := []byte(src)
	block, err := aes.NewCipher([]byte(k.Key)) //添加密钥  块大小已经被固定大小位为16字节 1字节是8为位
	if err != nil {
		return src
	}
	NewSrcByte := PadPwd(srcByte, block.BlockSize()) //填充密码长度  明文长度必须是16字节的整数倍 才可以方便被分割 分别进行块加密
	dst := make([]byte, len(NewSrcByte))
	block.Encrypt(dst, NewSrcByte)                //默认加密算法  ecb加密算法  最后进行拼接
	pwd := base64.StdEncoding.EncodeToString(dst) //放在参数中会导致错误
	return pwd
}

// UnPadPwd 去掉填充部分
func UnPadPwd(dst []byte) ([]byte, error) {
	if len(dst) <= 0 {
		return dst, errors.New("长度有误！")
	}
	unpadNum := int(dst[len(dst)-1])
	strErr := "error"
	op := []byte(strErr)
	if len(dst) < unpadNum {
		return op, nil
	}
	str := dst[:(len(dst) - unpadNum)]
	return str, nil
}

// AesDecoding 解密
func (k *Encryption) AesDecoding(pwd string) string {
	pwdByte := []byte(pwd)
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return pwd
	}
	block, errblock := aes.NewCipher([]byte(k.Key))
	if errblock != nil {
		return pwd
	}
	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte)
	dst, err = UnPadPwd(dst) //去掉填充部分
	if err != nil {
		return "0"
	}
	return string(dst)
}

func (k *Encryption) SetKey(key string) {
	k.Key = key
}

func PadKey(key []byte) []byte {
	keySize := 16 // 128 比特密钥长度为 16 字节
	if len(key) >= keySize {
		return key[:keySize] // 如果密钥长度已经足够，直接返回前 16 字节
	}
	paddedKey := make([]byte, keySize)
	copy(paddedKey, key) // 将原始密钥复制到新的填充密钥中
	return paddedKey
}
