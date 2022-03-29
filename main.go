package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Hello world\n")
	startWeb()
}

func startWeb() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello golang!!!")
}
