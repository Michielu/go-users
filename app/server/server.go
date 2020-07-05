package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Michielu/go-users/app/database"
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

type MyUserItem struct {
	TestId    string
	CreatedAt int
	Username  string
}

//Fields of Struct has to start as capital
type Item struct {
	UserId   string
	Username string
}

func (s *Server) Start() {
	var db *database.Database
	var err error

	db, err = database.Connect()

	if err != nil {
		log.Fatal(err) //TODO: panic and recover
	}

	db.GetUser("asdf1234")

	userId := "aaaaa11112"
	username := "testuername"

	db.PostUser(userId, username)

	// sess, err := session.NewSession(&aws.Config{
	// 	Region: aws.String("us-west-2")},
	// )

	// Create DynamoDB client
	// svc := dynamodb.New(*db.session)

	//Other
	// creds, err := *db.Session.Config.Credentials.Get()
	// if err != nil {
	// 	log.Fatal("a", err) //TODO: panic and recover
	// }
	log.Println("a", *db.Service)
	http.HandleFunc("/", HelloServer)
	// log.Fatal(creds)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// fmt.Println(w, "Hi %d", time.Now())
	//Add some logger
	t := Thing{42, r.URL.Path[1:]}
	log.Println(r.URL.Path[1:])
	j, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	jsonData := []byte(j)
	w.Write(jsonData)
}
