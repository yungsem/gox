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

// OpenFile 判断文件是否存在，存在则 OpenFile 不存在则 Create
func OpenFile(fileName string) (*os.File, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return os.Create(fileName)
	}
	return os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
}
