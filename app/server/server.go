package server

import (
	"encoding/json"
	"log"
	"net/http"
)

/**TODO:
- Include logger
- Enable HTTPS
- Remove encoding/josn -- have it be converted and returned in other packages
*/

type Thing struct {
	Field1 int    `json:"field1"`
	Field2 string `json:"field2"`
}

// Server represents the server which listens for connections when started
type Server struct {
	Hostname  string `json:"hostname"`  // Server name
	UseHTTP   bool   `json:"UseHTTP"`   // Listen on HTTP
	UseHTTPS  bool   `json:"UseHTTPS"`  // Listen on HTTPS
	HTTPPort  int    `json:"HTTPPort"`  // HTTP port
	HTTPSPort int    `json:"HTTPSPort"` // HTTPS port
	CertFile  string `json:"CertFile"`  // HTTPS certificate
	KeyFile   string `json:"KeyFile"`   // HTTPS private key
	Handler   http.Handler
}

func (s *Server) Start() {
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
