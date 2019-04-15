package helper

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type EasyToken struct {
	Uid      int64
	Username string
	Password string
	Role     string
}

var (
	verifyKey  string
	ErrAbsent  = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" // 令牌无效
	ErrExpired = "token expired" // 令牌过期
	ErrOther   = "other error"   // 其他错误
)

func init() {
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	verifyKey = appConf.String("jwt::token")
}

func (e EasyToken) GetToken() (string, error){
	//自定义claim
	claim := jwt.MapClaims{
		"exp": 		time.Now().Add(time.Hour * time.Duration(25)).Unix(),
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"info":		e,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	return token.SignedString([]byte(verifyKey))
}

func (e EasyToken) ValidateToken(tokenString string) (bool, error) {
	if tokenString == "" {
		return false, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(verifyKey), nil
	})
	if token == nil {
		return false, errors.New(ErrInvalid)
	}
	if token.Valid {

		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New(ErrInvalid)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New(ErrExpired)
		} else {
			return false, errors.New(ErrOther)
		}
	} else {
		return false, errors.New(ErrOther)
	}
}

func ParseToken(tokenString string, key string) (interface{}, bool,string){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["info"], true,""
	} else {
		return "", false,err.Error()
	}
}