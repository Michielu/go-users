package database

import (
	"errors"
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

func (db *Database) GetMasterUser(userId string) (*dynamodb.GetItemOutput, error) {
	svc := *db.Service

	// If a DynamoDB table has a partition key and a sort key, can't use GetItem to get a single item in table.
	// Need to use Query.
	return svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(DB_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"UserId": {
				S: aws.String(userId),
			},
		},
	})
}

func (db *Database) CreateMasterUser(newMasterUser modals.MasterUser) (*modals.MasterUser, error) {
	svc := *db.Service

	//TODO check to see if username exists
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(DB_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"UserId": {
				S: aws.String(newMasterUser.Username),
			},
		},
	})

	_, err = checkValidUsername(result, newMasterUser.Username)

	if err != nil {
		fmt.Println("Username of ", newMasterUser.Username, " exists")
		return nil, err
	}

	av, err := dynamodbattribute.MarshalMap(newMasterUser)
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

	fmt.Println("Successfully added '" + newMasterUser.UserId + "' (" + newMasterUser.Username + ") to table " + DB_NAME)

	return &newMasterUser, nil

}

func checkValidUsername(result *dynamodb.GetItemOutput, username string) (bool, error) {
	if len(result.Item) != 0 {
		fmt.Println("Username of ", username, " exists")
		return false, errors.New("Invalid Username")
	}
	return true, nil
}

// Close kills the current session and ends the Database connection
// func (d *Database) Close() {
// 	if d.session != nil {
// 		d.session.Close()
// 	}
// 	d.session = nil
// 	log.Println("Database closed")
// }
