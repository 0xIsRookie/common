package file

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

// TraverseFiles 遍历指定目录下的所有文件，并返回所有文件的路径列表
//
// 参数：
//   - rootPath: 待遍历的根目录路径
//
// 返回值：
//   - []string: 所有文件的路径列表
//   - error: 遍历过程中的错误信息，如果没有错误，则为 nil
func TraverseFiles(rootPath string) ([]string, error) {
	var result []string
	err := filepath.Walk(rootPath,
		func(path string, info os.FileInfo, err error) error {
			// 如果遍历过程中出现错误，则直接返回
			if err != nil {
				return err
			}

			// 如果当前路径不是一个文件，则忽略
			if !info.Mode().IsRegular() {
				return nil
			}

			// 将文件路径添加到结果列表中
			asbPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			result = append(result, asbPath)

			return nil
		})
	return result, err
}

// ReadAllFilesByLine 读取指定目录下所有文件的内容，并按行合并到一个字符串列表中
//
// 参数：
// - rootPath: 待读取文件的根目录路径
//
// 返回值：
// - []string: 所有文件的内容列表，每个元素为文件的一行文本
// - error: 读取过程中的错误信息，如果没有错误，则为 nil
func ReadAllFilesByLine(rootPath string) ([]string, error) {
	filePaths, err := TraverseFiles(rootPath)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, path := range filePaths {
		contents, err := ReadFileByLine(path)
		if err != nil {
			continue
		}
		result = append(result, contents...)
	}
	return result, err
}

// ReadFileByLine 按行读取指定文件的内容，并将内容以字符串切片的形式返回
//
// 参数：
// - path：要读取的文件路径
//
// 返回值：
// - []string：文件内容按行分割后的字符串切片
// - error：读取文件过程中出现的错误信息，如果没有错误则为 nil
func ReadFileByLine(path string) ([]string, error) {
	var result []string
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		return result, err
	}
	defer file.Close()

	// 逐行读取文件内容
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return result, err
		}
		content := string(line)
		result = append(result, content)
	}
	return result, nil
}
