package goutil

import "io/ioutil"

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
