package goutil

import (
	"io/ioutil"
	"os"
)

func FileList(dirPath string) ([]string, error) {
	lis, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	fileList := make([]string, 0)
	for _, f := range lis {
		if f.IsDir() {
			continue
		}
		fileList = append(fileList, f.Name())
	}
	return fileList, nil
}

// HasDir 判断文件夹是否存在
func ExistDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

// CreateDir 创建文件夹
func CreateDir(path string) error {
	_exist, _err := ExistDir(path)
	if _err != nil {
		return _err
	}
	if !_exist {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
