package modals

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type MasterUser struct {
	Apps      []string
	CreatedAt int64
	Email     string
	UserId    string
	Username  string
	Password  string
}

type Database struct {
	Session *session.Session
	Service *dynamodb.DynamoDB
}

type SimpleResponse struct {
	Success      bool   `json:"success"`
	ShortMessage string `json:"shortMessage"`
}

// //TODO figu
// type ClientError struct {
// 	Type    int    `json:"code,omitempty"`
// 	Message string `json:"message,omitempty"`
// }

// func (e *ErrorResponse) Error() string {
// 	return fmt.Sprintf("Invalid Username")
// }
