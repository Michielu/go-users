//The first statement in a Go source file must be package name.
//Executable commands must always use package main.package main
package main

import (
	"github.com/Michielu/go-users/app/server"
)

type Thing struct {
	Field1 int    `json:"field1"`
	Field2 string `json:"field2"`
}

// Use all CPU cores
// runtime.GOMAXPROCS(runtime.NumCPU())
func main() {
	// http.HandleFunc("/", HelloServer)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	s := &server.Server{
		UseHTTP:  true,
		HTTPPort: 4444,
	}

	s.Start()
}
