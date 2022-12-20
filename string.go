package goutil

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func StrInField(key string, str string) bool {
	str = strings.ReplaceAll(str, "，", ",")
	s1 := strings.Split(str, ",")
	for _, v := range s1 {
		if v == key {
			return true
		}
	}
	return false
}

func StrInSlice(key string, slice []string) bool {
	if len(slice) == 0 {
		return false
	}

	for _, v := range slice {
		if v == key {
			return true
		}
	}

	return false
}

func StrInPrefix(s string, arr []string) bool {
	for _, a := range arr {
		if strings.HasPrefix(s, a) {
			return true
		}
	}
	return false
}

// ArrayUnique 数组去重
func StrUnique(arr []string) []string {
	set := make(map[string]bool)
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = true
		arr[j] = v
		j++
	}
	return arr[:j]
}

// 生成随机字符串
func RandomStr(length int64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := int64(0); i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
