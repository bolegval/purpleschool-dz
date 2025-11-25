package bins

import (
	"errors"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

func NewBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	if name == "" {
		return nil, errors.New("name is required")
	}

	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}, nil
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBinList(bin *Bin) *BinList {
	var bins []Bin
	bins = append(bins, *bin)
	return &BinList{
		Bins: bins,
	}
}
