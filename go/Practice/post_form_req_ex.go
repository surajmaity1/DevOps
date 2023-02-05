package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	PerformPostFormRequest()
}



func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "suraj")
	data.Add("lastname", "maity")
	data.Add("email", "surajxyz@xyz.com")

	response, errHttpPF := http.PostForm(myurl, data)

	if errHttpPF != nil {
		panic(errHttpPF)
	}

	defer response.Body.Close()

	content, errio := ioutil.ReadAll(response.Body)

	if errio != nil {
		panic(errio)
	}

	fmt.Println(string(content))
}
func printErr(err error) {
	if err != nil {
		panic(err)
	}
}
