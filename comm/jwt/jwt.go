/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 02:14:16
 * @LastEditTime : 2020-07-14 02:24:01
 * @Description  : jwt
 */ 
 
package jwt

 import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("secret")

type Claims struct {
	UserId uint64 `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken 生成token
// userId 用户id
// expireDate 过期时间
func GenerateToken(userId uint64, expireDate int) (string, error) {
	nowTime := time.Now()
	// 过期时间默认3
	if expireDate == 0 {
		expireDate = 3
	}
	expireTime := nowTime.Add(time.Duration(expireDate) * time.Hour)

	claims := Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  nowTime.Unix(), //发行时间
			Issuer:    "micro",        //发行人
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 解析token
// ParseToken parsing token
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