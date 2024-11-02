package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/mutezebra/forum/config"
	"github.com/pkg/errors"
)

var (
	jwtSecret []byte
)

// CheckToken 接收一个token,返回uid,valid,err
func CheckToken(token string) (string, bool, error) {
	if len(jwtSecret) == 0 {
		jwtSecret = []byte(config.Secret.JwtSecret)
	}

	if token[0] == 'B' {
		bearer := token[0:6]
		if bearer == "Bearer" {
			token = token[7:]
		}
	}

	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return "", false, errors.Errorf("failed parse token,error: %v", err)
	}

	uid, ok := parseToken.Claims.(jwt.MapClaims)["sub"].(string)
	if !ok {
		return "", false, errors.Errorf("failed get uid from token")
	}

	return uid, parseToken.Valid, nil
}
