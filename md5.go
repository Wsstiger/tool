package tool

import "fmt"
import (
	"crypto/md5"
	"os"
	"io/ioutil"
)

func GetMD5String(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GetFileMd5String(file string) (string, error) {
	fd, err := os.OpenFile(file, os.O_RDONLY, 0400)
	if err != nil {
		return "", err
	}
	defer fd.Close()

	bytes, err := ioutil.ReadAll(fd)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", md5.Sum(bytes)), nil
}
