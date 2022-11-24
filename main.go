package main

import (
	"fmt"
	"github.com/gustavoalochio/desafio-luizalabs/counter"
	"github.com/gustavoalochio/desafio-luizalabs/reader"
	"os"
)

func main() {

	// open file
	f, err := os.Open("data.csv")
	if err != nil {
		return
	}

	// to close the file at the end of the program
	defer f.Close()

	// receive a csv file to generate a access struct list
	accessReader := reader.NewAccessReader()
	accessList, err := accessReader.Read(f)

	if err != nil {
		fmt.Printf("Error reading data from csv file: %v", err)
		return
	}

	//receive access struct list to create and count the domains
	counterDomain := counter.NewCounterDomain()
	countDomainList := counterDomain.CountDomains(accessList)

	// print the domains struct
	counterDomain.PrintDomainsSorted(countDomainList)
}
