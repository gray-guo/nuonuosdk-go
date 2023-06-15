package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	URL "net/url"
	"nuonuosdk/util"
	"strconv"
	"time"
)

const VERSION = "2.0"
const AUTH_URL = "https://open.nuonuo.com/accessToken"

func SendPostSyncRequest(url string, senid string, appKey string, appSecret string, token string, taxnum string, method string, content string) string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := strconv.FormatInt(util.RandtInt64(10000, 1000000000), 10)

	finalUrl := fmt.Sprintf("%s?senid=%s&nonce=%s&timestamp=%s&appkey=%s", url, senid, nonce, timestamp, appKey)

	urlInfo, err := URL.Parse(finalUrl)
	if err != nil {
		fmt.Println(err.Error())
		panic("URL解析错误")
	}

	sign := util.MakeSign(urlInfo.Path, appSecret, appKey, senid, nonce, content, timestamp)

	client := http.DefaultClient

	req , _ := http.NewRequest("POST", finalUrl, bytes.NewBufferString(content))

	req.Header["Content-Type"] = []string{"application/json"}
	req.Header["X-Nuonuo-Sign"] = []string{sign}
	req.Header["accessToken"] = []string{token}
	req.Header["userTax"] = []string{taxnum}
	req.Header["method"] = []string{method}
	req.Header["skdVer"] = []string{VERSION}

	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func RefreshISVToken(refreshToken string, userId string, appSecret string) string {
	var uri URL.URL
	params := uri.Query()
	params.Set("client_id", userId)
	params.Set("client_secret", appSecret)
	params.Set("refresh_token", refreshToken)
	params.Set("grant_type", "refresh_token")

	query := params.Encode()
	req , _ := http.NewRequest("POST", AUTH_URL, bytes.NewBufferString(query))

	req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}

	client := http.DefaultClient
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func getISVToken(appKey string, appSecret string, code string, taxnum string, redirectUri string) string {

	var uri URL.URL
	params := uri.Query()
	params.Set("client_id", appKey)
	params.Set("client_secret", appSecret)
	params.Set("code", code)
	params.Set("taxNum", taxnum)
	params.Set("redirect_uri", redirectUri)
	params.Set("grant_type", "authorization_code")

	query := params.Encode()
	req , _ := http.NewRequest("POST", AUTH_URL, bytes.NewBufferString(query))

	req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}

	client := http.DefaultClient
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}



