package main

type Db interface {
	Read() ([]byte, error)
	Write([]byte) error
}

func main() {

}
