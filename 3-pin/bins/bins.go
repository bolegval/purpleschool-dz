package bins

import (
	"errors"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
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
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}, nil
}

type BinList struct {
	Bins []Bin
}

func NewBinList(bin *Bin) *BinList {
	var bins []Bin
	bins = append(bins, *bin)
	return &BinList{
		Bins: bins,
	}
}
