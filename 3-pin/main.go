package main

import (
	"bin/bins"
	"bin/storage"
	"fmt"
)

func main() {
	binList := bins.NewBinListWithDb(storage.NewStorageDb("data.json"))

	fmt.Println(binList)
}
