package reader

import (
	"bytes"
	"github.com/gustavoalochio/desafio-luizalabs/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	accessListRead = []entity.Access{
		entity.Access{
			Url:   "loja.google.com",
			Count: 100,
		},
	}
)

func TestAccessReader_Read(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		//// arrange
		var buffer bytes.Buffer
		buffer.WriteString("100,loja.google.com")
		domainReader := NewAccessReader()

		// action
		content, err := domainReader.Read(&buffer)
		if err != nil {
			t.Error("Failed to read csv data")
		}

		//// assert
		assert.NotNil(t, content)
		assert.Equal(t, content, accessListRead)
	})
}
