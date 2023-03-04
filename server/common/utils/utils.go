package utils

import (
	"fmt"
	"os"
)

// FileExists 判断所给路径文件/文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func MkdirP(path string) error {
	if !FileExists(path) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
		fmt.Printf("创建文件夹: %v \n", path)
	}
	return nil
}
