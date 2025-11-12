package storage

import (
	"bin/bins"
	"bin/file"
	"encoding/json"
	"fmt"
)

func saveBin(bin *bins.Bin) {
	data, err := json.Marshal(bin)

	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}

	err = file.WriteJson(data, "data.json")
	if err != nil {
		return
	}
}

func readBin(fileName string) (*bins.Bin, error) {
	data, err := file.ReadJson(fileName)
	if err != nil {
		return nil, err
	}

	var bin bins.Bin
	err = json.Unmarshal(data, &bin)
	if err != nil {
		return nil, err
	}
	return &bin, nil
}
