package main

import (
	"bin/api"
	"bin/bins"
	"bin/config"
	"bin/storage"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Не удалось прочитать env файл")
		fmt.Println(err.Error())
	}
	api := api.NewApi(*config.NewConfig())
	fmt.Println(api)
	binList := bins.NewBinListWithDb(storage.NewStorageDb("data.json"))

	fmt.Println(binList)
}
