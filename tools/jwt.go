package tools

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

const (
	AccessTokenDuration  = 2 * time.Hour
	RefreshTokenDuration = 30 * 24 * time.Hour
	TokenIssuer          = "e-commerce-system-chen"
)

// GetToken 生成jwt
func GetToken(secretKey, username string, userId int64, key string) (atoken string, err error) {
	claims1 := jwt.MapClaims{
		"userId":   userId,
		"username": username,
		"key":      key,
		"iat":      time.Now().Unix(), // 设置令牌的签发时间
		"iss":      TokenIssuer,
		"exp":      time.Now().Add(RefreshTokenDuration).Unix(), // 设置过期时间为当前时间加上指定的秒数 2个小时失效
	}
	atoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims1).SignedString([]byte(secretKey))
	if err != nil {
		logx.Error("生成Atoken:", err)
	}
	return atoken, err
}

// VerityToken 验证token
func VerityToken(secretKey string, aToken string) (int64, string) {
	token, err := jwt.Parse(aToken, func(token *jwt.Token) (interface{}, error) {
		// 返回用于验证签名的密钥
		return []byte(secretKey), nil
	})
	fmt.Println(token)
	// 检查错误
	if err != nil {
		fmt.Println("Failed to parse JWT:", err)
		return 0, ""
	}
	// 检查 token 是否有效
	if token.Valid {
		// 从 token 的 Claims 中提取用户 ID 和姓名等信息
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			userID := claims["userId"].(float64)
			name := claims["username"].(string)
			return int64(userID), name
		} else {
			fmt.Println("Invalid token claims")
			return 0, ""
		}
	} else {
		fmt.Println("Invalid token")
		return 0, ""
	}
}

// GetSendToken 生成jwt
func GetSendToken(secretKey, username string, userId int64, email string, password string, OperationType string) (atoken string, err error) {
	claims1 := jwt.MapClaims{
		"userId":        userId,
		"username":      username,
		"email":         email,
		"password":      password,
		"OperationType": OperationType,
		"iat":           time.Now().Unix(), // 设置令牌的签发时间
		"iss":           TokenIssuer,
		"exp":           time.Now().Add(AccessTokenDuration).Unix(), // 设置过期时间为当前时间加上指定的秒数
	}
	atoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims1).SignedString([]byte(secretKey))
	if err != nil {
		logx.Error("生成Atoken:", err)
	}
	return atoken, err
}

// VeritySendToken 验证token
func VeritySendToken(secretKey string, aToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(aToken, func(token *jwt.Token) (interface{}, error) {
		// 返回用于验证签名的密钥
		return []byte(secretKey), nil
	})
	fmt.Println(token)
	// 检查错误
	if err != nil {
		fmt.Println("Failed to parse JWT:", err)
		return nil, err
	}
	// 检查 token 是否有效
	if token.Valid {
		// 从 token 的 Claims 中提取用户 ID 和姓名等信息
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			return claims, err
		} else {
			fmt.Println("Invalid token claims")
			return nil, err
		}
	} else {
		fmt.Println("Invalid token")
		return nil, err
	}
}
