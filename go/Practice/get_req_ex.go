package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	PerformGetRequest()
}

func PerformGetRequest() {
	const url string = "http://localhost:8000/get"

	response, errHttp1 := http.Get(url)
	printErr(errHttp1)

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content Len:", response.ContentLength)

	//c, _ := ioutil.ReadAll(response.Body)
	//fmt.Println("Content", string(c))

	// USING STRING BUILDER

	var responseStrBldr strings.Builder

	content, errHttp2 := ioutil.ReadAll(response.Body)
	printErr(errHttp2)
	byte_count, errResBldr := responseStrBldr.Write(content)
	printErr(errResBldr)

	fmt.Println("ByteCount:", byte_count)
	fmt.Println("Content:", responseStrBldr.String())
}
func printErr(err error) {
	if err != nil {
		panic(err)
	}
}
