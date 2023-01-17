package utils

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cristalhq/jwt"
)

const (
	SECRETKEY = "secret-key"
)

func JWTSign(uid string) (string, error) {
	signer, err := jwt.NewHS256([]byte(SECRETKEY))
	if err != nil {
		return "", err
	}
	tokenBuilder := jwt.NewTokenBuilder(signer)
	token, err := tokenBuilder.Build(jwt.StandardClaims{
		Audience:  jwt.Audience{"todo-service"},
		ExpiresAt: jwt.Timestamp(time.Now().Add(time.Minute * 30).UnixMilli()),
		ID:        uid,
		IssuedAt:  jwt.Timestamp(time.Now().UnixMilli()),
		Issuer:    "todo-service",
		Subject:   "todo-service",
	})
	if err != nil {
		return "", err
	}
	return string(token.Raw()), nil
}

func JWTVerify(stringToken string) (string, error) {
	signer, err := jwt.NewHS256([]byte(SECRETKEY))
	if err != nil {
		return "", err
	}
	token, err := jwt.ParseAndVerify([]byte(stringToken), signer)
	if err != nil {
		return "", err
	}

	payloadRaw := token.RawClaims()

	data := make(map[string]interface{})
	err = json.Unmarshal(payloadRaw, &data)
	if err != nil {
		return "", err
	}

	uid, ok := data["jti"].(string)
	if !ok {
		return "", fmt.Errorf("invalid token")
	}

	return uid, nil
}
