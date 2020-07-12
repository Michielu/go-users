package account

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Michielu/go-users/app/database"
	"github.com/Michielu/go-users/app/modals"
	"github.com/Michielu/go-users/app/services/passwords"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gorilla/mux"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request, db *database.Database) {
	log.Println("routes.getUserHandler")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	item, _ := GetMasterUserFromDb(db, vars["userId"])

	j, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	jsonData := []byte(j)
	w.Write(jsonData)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request, db *database.Database) {
	w.Header().Set("Content-Type", "application/json")
	var p modals.MasterUser

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.CreatedAt = time.Now().Unix()
	p.Password = passwords.HashAndSalt([]byte(p.Password))
	item, err := db.CreateMasterUser(p)

	if err != nil {
		fmt.Fprintf(w, "Error: %+v", err)
	}

	j, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	jsonData := []byte(j)
	w.Write(jsonData)
}

func GetMasterUserFromDb(db *database.Database, userId string) (*modals.MasterUser, error) {
	result, err := db.GetMasterUser(userId)

	if err != nil {
		log.Println(err)
	}

	masterUser := modals.MasterUser{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &masterUser)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return &masterUser, nil
}
