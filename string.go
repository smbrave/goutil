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

/**
 * 驼峰转蛇形 snake string
 * @description XxYy to xx_yy , XxYY to xx_y_y
 * @date 2020/7/30
 * @param s 需要转换的字符串
 * @return string
 **/

func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

/**
 * 蛇形转驼峰
 * @description xx_yy to XxYx  xx_y_y to XxYY
 * @date 2020/7/30
 * @param s要转换的字符串
 * @return string
 **/

func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
