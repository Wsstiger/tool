package tool

import (
	"os"
	"path/filepath"
)

/*
	级联创建文件夹：不存在，新建
*/
func MakeDir(path string) bool {
	if !IsDirExist(path) {
		err := os.MkdirAll(path, 0777)
		return err == nil
	}
	return true
}

// 检查文件夹是否存在
func IsDirExist(path string) bool {
	f, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return false
	}
	return f.IsDir()
}

// 新建文件夹、级联创建
func NewDirPath(path string) bool {
	err := os.MkdirAll(path, 0777)
	return err == nil
}

// 检查文件是否存在
func IsFileExist(path string) bool {
	f, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return false
	}
	return !f.IsDir()
}

// 删除所有文件
func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

// 获取文件件夹
func GetDirs(dirpath string) ([]string, error) {
	var dir_list []string
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				dir_list = append(dir_list, path)
				return nil
			}

			return nil
		})
	return dir_list, dir_err
}
