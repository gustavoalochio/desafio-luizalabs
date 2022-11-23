package reader	

import (
    "encoding/csv"
    "fmt"
    "os"
    "io"
    "testing"
    "reader"
)


func readFile(reader io.Reader) ([][]string, error) {
    r := csv.NewReader(reader)
    lines, err := r.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
    return lines, err
}

func TestReadFile(t *testing.T) {

    // arrange
    var buffer bytes.Buffer
    buffer.WriteString("10,loja.google.com")
    domainReader := reader.NewDomainReader()

    // action
    content, err := domainReader.Read(&buffer)
    if err != nil {
        t.Error("Failed to read csv data")
    }

    // assert
    t.Assert() // ...
}