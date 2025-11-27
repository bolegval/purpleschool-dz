package main

import (
	"bin/bins"
	"bin/storage"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Не удалось прочитать env файл")
	}
	binList := bins.NewBinListWithDb(storage.NewStorageDb("data.json"))

	fmt.Println(binList)
}
