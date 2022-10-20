package main

import (
	"fmt"
	"net/http"
	)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Hello World")
}

func main() {
	http.HandleFunc("/", handler)
	
	fmt.Println("listening at port: 8080")
	err := http.ListenAndServe(":8080", nil)
	if (err != nil) {
		panic(err)
	}
}
