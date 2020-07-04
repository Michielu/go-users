package db

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Database struct {
	session *session.Session
	// service *dynamodb
}

// Datastore represents a store for the data (Database session)
// type Datastore struct {
// 	session *session.Session
// 	name    string
// }

func Connect() (*Database, error) {
	log.Printf("Connecting to Database")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	log.Println("Connected to Database")
	return &Database{session: sess}, nil
}

// Close kills the current session and ends the Database connection
// func (d *Database) Close() {
// 	if d.session != nil {
// 		d.session.Close()
// 	}
// 	d.session = nil
// 	log.Println("Database closed")
// }
