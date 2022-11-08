package main

import (
	"fmt"
	//"html/template"
	"net/http"

	"github.com/akshayvidap128/go-course/templates/cmd/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("starting appication on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
