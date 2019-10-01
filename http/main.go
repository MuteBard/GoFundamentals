package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//make a byte slice with empty 999999 values inside
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Print("Response: ", string(bs))
}
