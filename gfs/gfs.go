package gfs

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Path Exists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	} else {
		return true
	}
}

// File Exists
func FileExists(fileUri string) bool {
	fileInfo, err := os.Stat(fileUri)
	if err != nil {
		return false
	}
	if fileInfo.IsDir() {
		return false
	}
	return true
}

// Dir Exists
func DirExists(dir string) bool {
	fileInfo, err := os.Stat(dir)
	if err != nil {
		return false
	}
	if fileInfo.IsDir() {
		return true
	}
	return false
}

// Copy File
func CopyFile(srcUri string, dstUri string) error {
	reader, err := os.Open(srcUri)
	if err != nil {
		return err
	}
	defer reader.Close()
	dst, errCopyOpen := os.OpenFile(dstUri, os.O_CREATE|os.O_RDWR, 0777)
	if errCopyOpen != nil {
		return errCopyOpen
	}
	defer dst.Close()
	_, errCopy := io.Copy(dst, reader)
	if errCopy != nil {
		return errCopy
	} else {
		return nil
	}
}

// Copy Dir
func CopyDir(src string, dst string) error {
	fileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return CopyFile(src, dst)
	}
	MakeDir(dst)
	list, err := os.ReadDir(src)
	if err == nil {
		for _, file := range list {
			CopyDir(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()))
		}
	}
	return nil
}

// Make Dir
func MakeDir(dir string) error {
	err := os.MkdirAll(dir, 0777)
	return err
}

// Remove Dir
func RemoveDir(dir string) error {
	err := os.RemoveAll(dir)
	return err
}

// Dir Size
func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// ModifyTime
func ModifyTime(path string) (int64, error) {
	f, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return f.ModTime().Unix(), nil
}

// Read file
func ReadFile(fileUrl string) (string, error) {
	content, err := os.ReadFile(fileUrl)
	if err != nil {
		return "", err
	} else {
		return string(content), nil
	}
}

// Write Content To File
func WriteContentToFile(content string, dir string, fileUrl string) error {
	if !DirExists(dir) {
		MakeDir(dir)
	}
	return os.WriteFile(dir+fileUrl, []byte(content), 0777)
}

// Append Content To File
func AppendContentToFile(content string, fileUrl string) error {
	file, err := os.OpenFile(fileUrl, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

// Prepend Content To File
func PrependContentToFile(content string, fileUrl string) error {
	contentOld, err := ReadFile(fileUrl)
	if err != nil {
		return err
	} else {
		content = content + contentOld
		return os.WriteFile(fileUrl, []byte(content), 0777)
	}
}

// Get Extension
func GetExtension(fileName string) string {
	fileNames := strings.Split(fileName, ".")
	return strings.ToLower(fileNames[len(fileNames)-1])
}

// GET File Name
func GetFileName(filePath, Splitor string) string {
	fileNames := strings.Split(filePath, Splitor)
	return fileNames[len(fileNames)-1]
}
