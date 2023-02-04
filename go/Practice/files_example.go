package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("")
	text := "Hey this is text. We will write in a file"

	file, errorOs := os.Create("./myfile.txt")
	checkNilErr(errorOs)

	length, errorIo := io.WriteString(file, text)
	checkNilErr(errorIo)

	fmt.Println("Length is : ", length)
	defer file.Close()

	readFile("./myfile.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("Text Data in Byte format: ", dataByte)
	fmt.Println("\n\n\nText Data in Text format: ", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
