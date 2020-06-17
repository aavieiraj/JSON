package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	jsonFile, err := os.Open(`/home/aavieiraj/go/src/examples_json/json/arquivo.json`)

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValueJSON, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	transactionFile := Transaction{}

	err = json.Unmarshal(byteValueJSON, &transactionFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(transactionFile)
}
