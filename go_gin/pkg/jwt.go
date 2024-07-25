package pkg

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"go_gin/entity"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (

	// 过期时间
	TOKEN_EXPIRATION_TIME = 24 * time.Hour
)

type InplantJWT struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	// jwt的Map，保证一个用户一个jwt
	Tokens sync.Map
}

func (j *InplantJWT) Init() {

}

func (j *InplantJWT) Generate(info Claims) (jwtStr string, err error) {
	nowTime := time.Now().Local()
	// 创建 Claims
	claims := Claims{
		MFD:      nowTime.Unix(),
		EXP:      nowTime.Add(TOKEN_EXPIRATION_TIME).Unix(),
		UserId:   info.UserId,
		UserName: info.UserName,
		Role:     info.Role,
	}

	// 使用指定的签名方法创建token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	// 使用指定的密钥签名并获得完整的编码后的签名字符串， 存入map
	jwtStr, err = token.SignedString(j.PrivateKey)
	if err == nil {
		j.Tokens.Store(info.UserId, jwtStr)
	}
	return
}

func (j *InplantJWT) Parse(jwtStr string) (info Claims, err error) {

	tempFunc := func(*jwt.Token) (interface{}, error) {
		if j.PublicKey == nil {
			return nil, errors.New("PublicKey is nil")
		}
		return j.PublicKey, nil
	}

	token, err := jwt.ParseWithClaims(jwtStr, &Claims{}, tempFunc)
	if err != nil {
		entity.PrintSqlErr(fmt.Errorf("parseJWT failed , err: %s", err.Error()))
		return info, err
	}

	if token == nil || !token.Valid || token.Claims == nil {
		return info, errors.New("token is invalid")
	}

	// 类型断言
	tokenClaims, ok := token.Claims.(*Claims)
	if !ok {
		return info, errors.New("token is invalid")
	}

	info = *tokenClaims
	valueJwt, bSave := j.Tokens.Load(info.UserId)
	if !bSave {
		return info, errors.New("user has no jwt")
	}

	if valueJwt != jwtStr {
		return info, errors.New("user jwt be updated")
	}

	return info, nil
}

// 自定义声明
type Claims struct {
	MFD      int64  // 生产时间 manufacturing date
	EXP      int64  // 过期时间 expiring date
	UserId   int64  // 用户id， 唯一
	UserName string // 用户名
	Role     int64  // 权限
}

// 验证是否过期
func (c *Claims) Valid() error {
	if c.EXP > time.Now().Local().Unix() {
		return errors.New("token is overdue")
	}
	return nil
}
