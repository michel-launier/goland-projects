package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
)

var serverPort int = 10000
var serverPortStr = strconv.Itoa(serverPort)

var page = "<html>\n  <head>\n    <title>Static Website</title>\n  </head>\n  <body>\n    <h2>Static Website</h2>\n  </body>\n</html>"

func homePage(w http.ResponseWriter, _ *http.Request) {
	//	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Fprintf(w, page)
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":"+serverPortStr, nil))
}

func main() {
	myOS, myArch := runtime.GOOS, runtime.GOARCH
	fmt.Println("Hello World")
	fmt.Printf("Running on %s/%s\n", myOS, myArch)
	fmt.Println("Starting server on port " + serverPortStr)
	handleRequests()
}
