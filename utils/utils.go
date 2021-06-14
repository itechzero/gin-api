package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"sort"
	"time"
)

func Md5V1(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Md5V2(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func SignCheck(data map[string]string) string {
	var sliceKey []string
	//把key放入切片
	for k, _ := range data {
		sliceKey = append(sliceKey, k)
	}
	//对key进行升序排序
	sort.Strings(sliceKey)
	var str string
	for _, v := range sliceKey {
		str += fmt.Sprintf("%v=>%v", v, data[v])
	}
	signStr := Md5V1(str + "signStr")
	return signStr
}

func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ValidatePassword(password, hashedPassword string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false, errors.New("密码比对错误！")
	}
	return true, nil
}

// 随机字符串
func RandString(length int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
