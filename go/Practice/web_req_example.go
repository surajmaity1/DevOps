package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	response, errHttp := http.Get(url)

	readErr(errHttp)

	fmt.Printf("Response type: %T\n\n", response)

	defer response.Body.Close()

	dataByte, errioutil := ioutil.ReadAll(response.Body)
	readErr(errioutil)

	fmt.Println("Body code:")
	fmt.Println(string(dataByte))
}

func readErr(err error) {
	if err != nil {
		panic(err)
	}
}
