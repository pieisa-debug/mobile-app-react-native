package mobileappreactnative

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Response struct {
	Success bool      `json:"success"`
	Data    interface{} `json:"data"`
	Error   string    `json:"error"`
}

func GetFileBytes(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetFileHash(filePath string) (string, error) {
	hash := make([]byte, 32)
	hashBytes, err := GetFileBytes(filePath)
	if err != nil {
		return "", err
	}
	copy(hash, hashBytes[:32])
	return fmt.Sprintf("%x", hash), nil
}

func GetFileSize(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

func GetFileSizeMB(filePath string) (float64, error) {
	size, err := GetFileSize(filePath)
	if err != nil {
		return 0, err
	}
	return float64(size) / 1024 / 1024, nil
}

func GetFileMIMEType(filePath string) (string, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}
	return http.DetectContentType(info.Name()), nil
}

func GetFileExtension(filePath string) (string, error) {
	ext := filepath.Ext(filePath)
	return strings.TrimPrefix(ext, "."), nil
}

func GetFileBaseName(filePath string) (string, error) {
	return filepath.Base(filePath), nil
}

func GetFileDir(filePath string) (string, error) {
	return filepath.Dir(filePath), nil
}

func IsFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func ReadJSONFile(filePath string, v interface{}) error {
	data, err := GetFileBytes(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func ReadJSONString(jsonString string, v interface{}) error {
	return json.Unmarshal([]byte(jsonString), v)
}

func LogError(err error) {
	log.Println(err)
}