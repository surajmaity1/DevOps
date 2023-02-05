package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//PerformGetRequest()
	//PerformPostJsonRequest()
	PerformPostFormRequest()
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
func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price": 0,
			"platform":"learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	printErr(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
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
