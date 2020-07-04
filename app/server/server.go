package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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

type Item struct {
	userId    string
	createdAt int
	username  string
}

func (s *Server) Start() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	tableName := "MyUsers"
	getUser := "1a2b3c4d"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(getUser),
			},
			// "userId": aws.String(getUser),
		},
	})

	log.Println("b", result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	creds, err := sess.Config.Credentials.Get()
	if err != nil {
		log.Fatal("a", err) //TODO: panic and recover
	}
	log.Println("a", creds)
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
