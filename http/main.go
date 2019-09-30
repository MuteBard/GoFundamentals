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

	//DONT DO THIS, THIS IS BAD. DONT BE BAD
	//this gives you no response body to work with
	//interfaces will help us deal with this
	fmt.Println(resp)
}
