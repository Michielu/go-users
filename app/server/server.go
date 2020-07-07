package server

import (
	"log"
	"net/http"

	"github.com/Michielu/go-users/app/database"
	"github.com/Michielu/go-users/app/routes"
)

/**TODO:
- Include logger
- Enable HTTPS
- Remove encoding/josn -- have it be converted and returned in other packages
*/

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
	var db *database.Database
	var err error

	db, err = database.Connect()

	if err != nil {
		log.Fatal(err) //TODO: panic and recover
	}

	s.Handler = routes.LoadRoutes(db)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", s.Handler))
}
