package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"
func main(){

	fmt.Println(myurl)

	// PARSING
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	// FETCH ALL PARAM	
	qparams := result.Query()

	fmt.Printf("qparam's type: %T\n", qparams)
	for key, val := range qparams{
		fmt.Println("Key:", key, " Val:", val)
	}


	// CONSTRUCT A URL

	diff_parts_url := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawQuery: "user=suraj",
	}

	construct_url := diff_parts_url.String()
	fmt.Println(construct_url)
}