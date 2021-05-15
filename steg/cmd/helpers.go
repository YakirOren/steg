package cmd

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
)

// checks if the err is nil
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// reads data from the file and checks from error
func ReadData(fileName string) []byte {
	data, err := ioutil.ReadFile(fileName)

	Check(err)
	return data
}

// delete_empty removes empty slices
func delete_empty(s [][]byte) [][]byte {
	var r [][]byte

	for _, file := range s {
		if !bytes.Equal(file, []byte{}) {
			r = append(r, file)
		}
	}
	return r
}

// CopyTemplate
func CopyTemplate(fileName string) string {

	fileName = strings.Split(fileName, ".")[0] + "1.gif"

	ioutil.WriteFile(fileName, ReadData("base.gif"), 0644)

	return fileName

}
