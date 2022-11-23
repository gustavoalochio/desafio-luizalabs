package reader

import (
	"encoding/csv"
	"fmt"
	"github.com/gustavoalochio/desafio-luizalabs/entity"
	"io"
	"strconv"
)

type AccessReader interface {
	Read(reader io.Reader) ([]entity.Access, error)
}

type accessReader struct{}

func NewAccessReader() AccessReader { return &accessReader{} }

func (d *accessReader) Read(reader io.Reader) ([]entity.Access, error) {

	if d == nil {
		return nil, fmt.Errorf("not implemented")
	}

	// read csv values using csv.Reader
	csvReader := csv.NewReader(reader)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var accessList []entity.Access

	for i := 0; i < len(data); i++ {
		count, _ := strconv.Atoi(data[i][0])
		accessList = append(accessList, entity.NewAccess(data[i][1], count))
	}

	return accessList, nil // TO-DO unmarshal csv -> entity
}
