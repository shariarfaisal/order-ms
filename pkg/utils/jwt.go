package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

func GenerateJWT(header string, payload map[string]interface{}, secret string, exp int64) (token string, err error) {
	h := hmac.New(sha256.New, []byte(secret))
	header64 := base64.StdEncoding.EncodeToString([]byte(header))
	payload["iat"] = time.Now().Unix()
	if exp > 0 {
		payload["exp"] = time.Now().Unix() + exp
	}
	payloadStr, err := json.Marshal(payload)

	if err != nil {
		err = errors.New("payload marshal error")
		return
	}

	payload64 := base64.StdEncoding.EncodeToString(payloadStr)
	message := header64 + "." + payload64
	fmt.Println("message", message)
	h.Write([]byte(message))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	token = message + "." + signature
	return
}

func ValidateJWT(token, secret string) (isValid bool, err error, decodedPayload map[string]interface{}) {
	errMsg := errors.New("Invalid token")

	tokenSplit := strings.Split(token, ".")
	if len(tokenSplit) != 3 {
		err = errMsg
		return
	}

	_, err = base64.StdEncoding.DecodeString(tokenSplit[0])
	if err != nil {
		err = errMsg
		return
	}

	payload, err := base64.StdEncoding.DecodeString(tokenSplit[1])
	if err != nil {
		err = errMsg
		return
	}

	message := tokenSplit[0] + "." + tokenSplit[1]
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	if signature != tokenSplit[2] {
		err = errMsg
		return
	}

	var payloadMap map[string]interface{}
	if err = json.Unmarshal(payload, &payloadMap); err != nil {
		err = errMsg
		return
	}

	exp, isExp := payloadMap["exp"]
	if isExp {
		if int64(exp.(float64)) < time.Now().Unix() {
			err = errors.New("Token expired")
			return
		}
	}

	isValid = true
	err = nil
	decodedPayload = payloadMap
	return
}
