//The first statement in a Go source file must be package name.
//Executable commands must always use package main.package main
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
