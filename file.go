package SHcommon

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	path = string(path[0:(strings.LastIndex(path, "/") + 1)])
	return path
}

func FileSaveString(data, filename string) (n int, err error) {
	fout, err := os.Create(filename)
	defer fout.Close()
	if err != nil {
		return -1, err
	}
	return fout.WriteString(data)
}

// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
