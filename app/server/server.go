package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	//Put item
	// tableName := "TestTable3"
	// // testId := "asdf1234"

	// // If a DynamoDB table has a partition key and a sort key, you can't use GetItem to get a single item in your table. You need to use something called Query.
	// result, err := svc.GetItem(&dynamodb.GetItemInput{
	// 	TableName: aws.String(tableName),
	// 	Key: map[string]*dynamodb.AttributeValue{
	// 		// M: aws.Map(mapQuery),
	// 		"UserId": {
	// 			S: aws.String("asdf1234"),
	// 		},
	// 	},
	// })

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// item := Item{}

	// err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	// if err != nil {
	// 	panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	// }

	// fmt.Println("Found item:")
	// fmt.Println("TestId:  ", item.UserId)
	// fmt.Println("Created: ", item.Username)

	// ---- Put item
	item := Item{
		UserId:   "aaaaa11111",
		Username: "1234567890",
	}

	fmt.Println(item)

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("length is: ", len(av))
	log.Println("b", av)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tableName := "TestTable3"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully added '" + item.UserId + "' (" + item.Username + ") to table " + tableName)

	//Other
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
