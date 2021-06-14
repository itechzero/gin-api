package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func SendAlarm(message string) []byte {
	url := sign()
	content, data := make(map[string]string), make(map[string]interface{})
	content["content"] = message
	data["msgtype"] = "text"
	data["text"] = content
	//data["at"] = ""
	dataJson, _ := json.Marshal(data)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(dataJson))
	if err != nil {

		return nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func sign() string {
	webhook := "https://oapi.dingtalk.com/robot/send?access_token="
	secret := ""
	milliTimeStamp := time.Now().UnixNano() / 1e6
	stringToSign := fmt.Sprintf("%d\n%s", milliTimeStamp, secret)
	sign := hmacSha256(stringToSign, secret)
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", webhook, milliTimeStamp, sign)
	return url
}

func hmacSha256(stringToSign string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
