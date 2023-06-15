package util

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"
)


func RandtInt64(min int64, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	if i.Int64() < min {
		RandtInt64(min, max)
	}
	return i.Int64()
}



func MakeSign(path string, appSecret string, appkey string, senid string, nonce string, body string, timestamp string) string {
	pieces := strings.Split(path, "/")

	signStr := fmt.Sprintf( "a=%s&l=%s&p=%s&k=%s&i=%s&n=%s&t=%s&f=%s", pieces[3],pieces[2], pieces[1], appkey, senid, nonce, timestamp, body)
	key := []byte(appSecret)
	mac := hmac.New(sha1.New,key)
	mac.Write([]byte(signStr))
	base64Str := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return base64Str
}


