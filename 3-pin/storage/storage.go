package storage

import (
	"bin/bins"
	"bin/file"
	"encoding/json"
	"fmt"
)

type StorageDb struct {
	fileName string
}

func NewStorageDb(name string) *StorageDb {
	return &StorageDb{
		fileName: name,
	}
}

func (db *StorageDb) SaveBin(bin *bins.Bin) {
	data, err := json.Marshal(bin)

	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}

	err = file.WriteJson(data, db.fileName)
	if err != nil {
		return
	}
}

func (db *StorageDb) ReadBin() (*bins.BinList, error) {
	data, err := file.ReadJson(db.fileName)
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
