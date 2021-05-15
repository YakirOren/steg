package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// turns out you cant have const byte slice in golang,
// so I created this funciton to return the gif header.
func getGIFHeader() []byte {
	return []byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x3B}
}

// Split extracts the stored files out of the combined file.
func Split(fileName string) {

	data := ReadData(fileName)

	// the dirName is the same as the fileName just with no ending.
	dirName := strings.Split(fileName, ".")[0]

	files := bytes.Split(data, getGIFHeader()) // get the stored files

	os.Mkdir(dirName, 0755) // create a new dir to store the files.

	// loop over the files and write them to the new folder.
	for i, file := range delete_empty(files) {

		path := dirName + "/" + dirName + strconv.Itoa(i)

		Check(ioutil.WriteFile(path, file, 0644))

	}

}

// prints the file content
// in hex form
// string form
// and the total size of the file.
func PrintFileContent(filename string) error {

	data := ReadData(filename)

	fmt.Printf("Data as hex: %x\n", data)
	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))

	return nil
}

// appends the given data to the end of the file.
func AppendToFile(filename string, data []byte) error {

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {

		return errors.Wrap(err, "open file failed")
	}

	if _, err := f.Write(data); err != nil {
		f.Close() // ignore error; Write error takes precedence

		return errors.Wrap(err, "writing to file failed")
	}

	// add the gif header to the end, so we could split it later.
	if _, err := f.Write(getGIFHeader()); err != nil {

		f.Close() // ignore error; Write error takes precedence

		return errors.Wrap(err, "adding base failed")
	}

	if err := f.Close(); err != nil {
		return errors.Wrap(err, "cant close file")
	}

	return nil

}
