/* Learning stuff:
1. Go Programming Langunage Specification: https://go.dev/ref/spec
2. Getting started with Sockets Concurrency in GoLang: https://dev.to/alicewilliamstech/getting-started-with-sockets-in-golang-2j66
3. Example: https://gobyexample.com
4. Tour: https://go.dev/tour/basics/1
5. Effective Go: https://go.dev/doc/effective_go
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
