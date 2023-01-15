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

func GenerateToken(header string, payload map[string]string, secret string) (string, error) {
	// create a new hash of type sha256. We pass the secret key to it
	h := hmac.New(sha256.New, []byte(secret))
	header64 := base64.StdEncoding.EncodeToString([]byte(header))
	// We then Marshal the payload which is a map. This converts it to a string of JSON.
	payloadstr, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error generating Token")
		return string(payloadstr), err
	}
	payload64 := base64.StdEncoding.EncodeToString(payloadstr)

	// Now add the encoded string.
	message := header64 + "." + payload64

	// We have the unsigned message ready.
	unsignedStr := header + string(payloadstr)

	// We write this to the SHA256 to hash it.
	h.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	//Finally we have the token
	tokenStr := message + "." + signature
	return tokenStr, nil
}

func GenerateJWT(header string, payload map[string]interface{}, secret string, exp int64) (token string, err error) {
	if secret == "" {
		err = errors.New("secret is empty")
		return
	}
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

	h := hmac.New(sha256.New, []byte(secret))
	payload64 := base64.StdEncoding.EncodeToString(payloadStr)
	message := header64 + "." + payload64
	unsignedStr := header + string(payloadStr)

	h.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	token = message + "." + signature
	return
}

// This helps in validating the token
func ValidateJWT(token, secret string) (isValid bool, err error, decodedPayload map[string]interface{}) {
	errMsg := errors.New("Invalid token")

	// JWT has 3 parts separated by '.'
	splitToken := strings.Split(token, ".")
	if len(splitToken) != 3 {
		err = errMsg
		return
	}

	// decode the header and payload back to strings
	header, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		fmt.Println("decode header", err)
		err = errMsg
		return
	}

	payload, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		fmt.Println("decode payload", err)
		err = errMsg
		return
	}

	//again create the signature
	unsignedStr := string(header) + string(payload)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	if signature != splitToken[2] {
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
