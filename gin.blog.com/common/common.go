package common

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"time"
)
const (
	// 可自定义盐值
	TokenSalt = "salt_stone"
)

type JsonData struct {
	Code	int	`json:"code"`
	Msg		string	`json:"msg"`
}

//登录注册面以及 token md5加密
func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}


//
//// 指定加密密钥
//var jwtSecret = []byte(setting.JwtSecret)
//
//type Claims struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//	jwt.StandardClaims
//}
//
////生成token
//func GenerateToken(username, password string) (string, error) {
//	nowTime := time.Now()
//	expireTime := nowTime.Add(3 * time.Hour)
//
//	claims := Claims{
//		username,
//		password,
//		jwt.StandardClaims {
//			ExpiresAt : expireTime.Unix(),
//			Issuer : "gin.blog.com",
//		},
//	}
//
//	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	token, err := tokenClaims.SignedString(jwtSecret)
//
//	return token, err
//}
//
////解析token
//func ParseToken(token string) (*Claims, error) {
//	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return jwtSecret, nil
//	})
//
//	if tokenClaims != nil {
//		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
//			return claims, nil
//		}
//	}
//
//	return nil, err
//}
//
//

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

var Secret =  "stone_"

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//生成token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin.blog.com",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

//解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
