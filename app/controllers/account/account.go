package account

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Michielu/go-users/app/database"
	"github.com/Michielu/go-users/app/modals"
	"github.com/gorilla/mux"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request, db *database.Database) {
	log.Println("routes.getUserHandler")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	item, _ := db.GetMasterUser(vars["id"])

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

	fmt.Fprintf(w, "Thing: %+v", p)

	item, _ := db.CreateMasterUser(p)

	j, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	jsonData := []byte(j)
	w.Write(jsonData)
}
