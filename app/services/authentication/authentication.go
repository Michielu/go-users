package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Michielu/go-users/app/database"
	"github.com/Michielu/go-users/app/modals"
	"github.com/Michielu/go-users/app/services/account"
	"github.com/Michielu/go-users/app/services/passwords"
)

func Login(w http.ResponseWriter, r *http.Request, db *database.Database) {
	w.Header().Set("Content-Type", "application/json")
	var p modals.MasterUser

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Thing: %+v", p)

	masterUser, err := account.GetMasterUserFromDb(db, p.Username)

	item := modals.SimpleResponse{}

	if err != nil {
		item = modals.SimpleResponse{Success: false, ShortMessage: "Unable to find user with username: " + p.Username}
		j, err := json.Marshal(item)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		jsonData := []byte(j)
		w.Write(jsonData)
	}

	if passwords.ComparePasswords(masterUser.Password, []byte(p.Password)) {
		item = modals.SimpleResponse{Success: true, ShortMessage: "Login username:password: " + p.Username + ":" + p.Password}

	}

	j, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	jsonData := []byte(j)
	w.Write(jsonData)
}
