package utils

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

// ReadTextFileToSlice TXT => Slice
func ReadTextFileToSlice(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// GetSuffixFiles 寻找目录下指定的后缀文件名
func GetSuffixFiles(path string, suffix string) ([]string, error) {
	var suFiles []string
	files, err := os.ReadDir(path)
	for _, file := range files {
		if strings.Contains(file.Name(), suffix) {
			filename := path + "/" + file.Name()
			suFiles = append(suFiles, filename)
		}
	}
	return suFiles, err
}

// GetCsvColumn 读取 csv 文件的某列
func GetCsvColumn(filename string, column int) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return []string{}
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return []string{}
	}
	var datas []string
	for i := 1; i < len(data); i++ {
		datas = append(datas, data[i][column-1])
	}
	return datas
}

// Deduplication Slice 去重
func Deduplication(lines []string) []string {
	m := make(map[string]bool)
	var result []string
	for _, line := range lines {
		if !m[line] {
			result = append(result, line)
			m[line] = true
		}
	}
	return result
}

// SliceWriter 将 Slice 写入文件
func SliceWriter(path string, lines []string) error {
	// 打开文件,没有的话就创建,有的话清空文件进行写入
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		writer.WriteString(line + "\n")
	}
	// 刷新缓存
	writer.Flush()
	return err
}

// RemoveDirectory 删除目录
func RemoveDirectory(dirPath string) error {
	return os.RemoveAll(dirPath)
}

// CreateDirectory 创建目录
func CreateDirectory(dirPath string) error {
	return os.MkdirAll(dirPath, os.ModePerm)
}

// DirectoryExists 判断目录是否存在
func DirectoryExists(dirPath string) bool {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// Dir 目录初始化操作
//  1. 存在则删除后创建
//  2. 不存在直接创建
func Dir(dirPath string) {
	if DirectoryExists(dirPath) {
		RemoveDirectory(dirPath)
		CreateDirectory(dirPath)
	}
	CreateDirectory(dirPath)
}
func AppendToFile(filePath string, textToAppend []string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, text := range textToAppend {
		file.WriteString(text + "\n")
	}
	return nil
}
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatalf("os.Open() : %v", err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("os.OpenFile() : %v", err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
