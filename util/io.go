package util

import "os"

func IsExist(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

func Mkdir(p string) error {
	return os.Mkdir(p, 0755)
}
