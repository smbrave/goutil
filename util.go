package goutil

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// format bytes number friendly
func BytesToTips(bytes uint64) string {
	switch {
	case bytes < 1024:
		return fmt.Sprintf("%dB", bytes)
	case bytes < 1024*1024:
		return fmt.Sprintf("%.2fK", float64(bytes)/1024)
	case bytes < 1024*1024*1024:
		return fmt.Sprintf("%.2fM", float64(bytes)/1024/1024)
	default:
		return fmt.Sprintf("%.2fG", float64(bytes)/1024/1024/1024)
	}
}

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func CopyStruct(dst interface{}, src interface{}) {

	dtype := reflect.TypeOf(dst)
	stype := reflect.TypeOf(src)

	if stype.Kind() != reflect.Ptr || stype.Kind() != dtype.Kind() {
		panic(errors.New("src/dst must ptr"))
	}
	if reflect.ValueOf(dst).IsNil() || reflect.ValueOf(src).IsNil() {
		panic(errors.New("src/dst is nil"))
	}

	dval := reflect.ValueOf(dst).Elem()
	sval := reflect.ValueOf(src).Elem()

	for i := 0; i < sval.NumField(); i++ {
		sValue := sval.Field(i)

		dValue := dval.FieldByName(sval.Type().Field(i).Name)
		if sValue.IsZero() || dValue.IsValid() == false || !dValue.CanSet() {
			continue
		}
		if sValue.Kind() != dValue.Kind() {
			continue
		}

		switch sValue.Type().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			dValue.SetInt(sValue.Int())

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			dValue.SetUint(sValue.Uint())

		case reflect.Float32, reflect.Float64:
			dValue.SetFloat(sValue.Float())

		case reflect.String:
			dValue.SetString(sValue.String())

		case reflect.Bool:
			dValue.SetBool(sValue.Bool())

		case reflect.Ptr:
			CopyStruct(dValue.Interface(), sValue.Interface())
		}
	}

}

func FormatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return strconv.FormatInt(bytes, 10) + " B"
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f%cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func HtmlStrip(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{1,}")
	src = re.ReplaceAllString(src, "")

	//去除&#12345;这类字符
	//re, _ = regexp.Compile("&#\\d*;")
	//src = re.ReplaceAllString(src, "")

	src = strings.ReplaceAll(src, "&nbsp;", "")
	src = strings.ReplaceAll(src, "nbsp;", "")
	src = strings.ReplaceAll(src, "& nbsp;", "")
	src = strings.ReplaceAll(src, "&nbsp", "")
	return strings.TrimSpace(src)
}
