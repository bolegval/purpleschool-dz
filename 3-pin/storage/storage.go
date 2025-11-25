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

func readBin(fileName string) (*bins.BinList, error) {
	data, err := file.ReadJson(fileName)
	if err != nil {
		return nil, err
	}

	var binList bins.BinList
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return nil, err
	}
	return &binList, nil
}
