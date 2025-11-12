package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadJson(fileName string) ([]byte, error) {
	if !strings.EqualFold(filepath.Ext(fileName), ".json") {
		return nil, errors.New(fmt.Sprintf("Файл %s не json", fileName))
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteJson(data []byte, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return errors.New("Не удалось создать файл")
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
