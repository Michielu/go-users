package database

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Database struct {
	Session *session.Session
	Service *dynamodb.DynamoDB
	// service *dynamodb
}

type Item struct {
	UserId   string
	Username string
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
	svc := dynamodb.New(sess)
	// log.Printf(svc)

	return &Database{Session: sess, Service: svc}, nil
}

func (db *Database) GetUser(userId string) {
	tableName := "TestTable3"

	svc := *db.Service

	// If a DynamoDB table has a partition key and a sort key, you can't use GetItem to get a single item in your table. You need to use something called Query.
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			// M: aws.Map(mapQuery),
			"UserId": {
				S: aws.String(userId),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		// return
	}

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	fmt.Println("TestId:  ", item.UserId)
	fmt.Println("Created: ", item.Username)
}

func (db *Database) PostUser(userId string, username string) {
	svc := *db.Service

	item := Item{
		UserId:   userId,
		Username: username,
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
}

// Close kills the current session and ends the Database connection
// func (d *Database) Close() {
// 	if d.session != nil {
// 		d.session.Close()
// 	}
// 	d.session = nil
// 	log.Println("Database closed")
// }
