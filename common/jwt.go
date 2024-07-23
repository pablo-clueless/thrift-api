package common

import (
	"crypto/subtle"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type ThriftClaims struct {
	Id string `json:"id" bson:"_id"`
	jwt.StandardClaims
}

var (
	JwtSecret       = "G23uA*!tF$%5s7v8Np@L2^&9hS6wRbQ!yC@3g2K*eA7zjP^kTdG#f1*H6w!Dm"
	JwtSigninMethod = jwt.SigningMethodHS256.Name
)

func GenerateToken(Id string) (string, error) {
	claims := ThriftClaims{
		Id: Id,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "thrift-api",
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func VerifyToken(str string) (ThriftClaims, error) {
	token, err := jwt.ParseWithClaims(str, &ThriftClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(JwtSecret), nil
	})
	if err != nil {
		return ThriftClaims{}, err
	}

	claims, ok := token.Claims.(*ThriftClaims)
	if !ok {
		return ThriftClaims{}, errors.New("invalid auth token")
	}

	if subtle.ConstantTimeCompare([]byte(claims.Id), []byte(claims.Id)) != 1 {
		return ThriftClaims{}, errors.New("invalid auth token")
	}

	return *claims, nil
}
