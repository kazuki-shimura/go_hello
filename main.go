package main

import (
	"fmt"
	"net/http"
	"study"
)

func main() {
	fmt.Printf("Hello world\n")
	Part1()
	startWeb()
}

func startWeb() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello golang!!!")
}
