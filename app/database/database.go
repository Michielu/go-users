package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Michielu/go-users/app/modals"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Database struct {
	Session *session.Session
	Service *dynamodb.DynamoDB
}

type Item struct {
	UserId   string
	Username string
}

var DB_NAME string = "MasterUsers"

func Connect() (*Database, error) {
	log.Printf("Connecting to Database %s", DB_NAME)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	log.Println("Connected to Database")

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	return &Database{Session: sess, Service: svc}, nil
}

func (db *Database) GetMasterUser(userId string) (*modals.MasterUser, error) {
	svc := *db.Service

	// If a DynamoDB table has a partition key and a sort key, can't use GetItem to get a single item in table.
	// Need to use Query.
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(DB_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"UserId": {
				S: aws.String(userId),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	masterUser := modals.MasterUser{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &masterUser)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return &masterUser, nil
}

func (db *Database) CreateMasterUser(newMasterUser modals.MasterUser) (*modals.MasterUser, error) {
	svc := *db.Service

	//TODO remove hardcode
	item := modals.MasterUser{
		Apps:      []string{"CREDIT_CARD"},
		CreatedAt: 12312333333,
		Email:     "exampleEmail@anotheremail.com",
		UserId:    "exampleUserID2",
		Username:  "exampleUsername2",
		Password:  "exmaplePassword123",
	}

	// fmt.Println(item)

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling new master user:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(DB_NAME),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully added '" + item.UserId + "' (" + item.Username + ") to table " + DB_NAME)

	return &item, nil

}

// Close kills the current session and ends the Database connection
// func (d *Database) Close() {
// 	if d.session != nil {
// 		d.session.Close()
// 	}
// 	d.session = nil
// 	log.Println("Database closed")
// }
