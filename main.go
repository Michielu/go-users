//The first statement in a Go source file must be package name.
//Executable commands must always use package main.package main
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Thing struct {
	Field1 int    `json:"field1"`
	Field2 string `json:"field2"`
}

func main() {
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Add some logger
	t := Thing{42, r.URL.Path[1:]}
	j, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	jsonData := []byte(j)
	w.Write(jsonData)

	// fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
