package filex

import (
	"bufio"
	"os"
)

// ReadLines 读取 file 中的所有行
func ReadLines(file *os.File) ([]string, error) {
	r := bufio.NewReader(file)

	var lines []string
	for {
		line, err := readLine(r)
		if err != nil {
			break
		}
		lines = append(lines, line)
	}
	return lines, nil
}

// readLine 读取 r 中的下一行
func readLine(r *bufio.Reader) (string, error) {
	var (
		isPrefix = true
		err      error
		line     []byte
		ln       []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}

	return string(ln), err
}

// CreateFile 创建一个没有任何内容的全新文件
// 如果文件已经存在，则清空该文件的内容
func CreateFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
}

// OpenFile 打开一个文件
// 如果文件不存在，则创建一个
func OpenFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

// MakeDir 创建单层目录或多层目录
// 如果目录已经存在，则什么都不干
func MakeDir(name string) error {
	return os.MkdirAll(name, os.ModeDir)
}

// ClearDir 清空目录下的所有内容
func ClearDir(name string) error {
	err := os.RemoveAll(name)
	if err != nil {
		return err
	}

	return os.MkdirAll(name, os.ModeDir)
}
