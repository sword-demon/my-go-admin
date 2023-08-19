package define

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	// JwtKey 密钥
	JwtKey = "sys_admin_wx"
	// TokenExpire token 有效期
	TokenExpire = time.Now().Add(time.Second * 3600 * 24 * 7).Unix()
	// RefreshTokenExpire 刷新 token 有效期
	RefreshTokenExpire = time.Now().Add(time.Second * 3600 * 24 * 14).Unix()
	// DefaultSize 默认分页大小
	DefaultSize = 10
)

type UserClaim struct {
	Id      uint
	Name    string
	IsAdmin bool // 是否是超管
	jwt.StandardClaims
}
