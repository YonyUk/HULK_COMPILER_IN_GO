package filesystem

import (
	"errors"
	"os"
)

type FileReader struct {
	file_path string
}

func NewFileReader(file_path string) (*FileReader, error) {
	fileinfo, err := os.Stat(file_path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("The specified path doesn't exists")
		}
		return nil, err
	}
	if fileinfo.IsDir() {
		return nil, errors.New("The specified path is not a file")
	}
	return &FileReader{
		file_path: file_path,
	}, nil
}

func (file *FileReader) ReadFile() (string, error) {
	data, err := os.ReadFile(file.file_path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
