package authentication

import (
	"encoding/json"
	"net/http"

	"github.com/Michielu/go-users/app/database"
	"github.com/Michielu/go-users/app/modals"
	"github.com/Michielu/go-users/app/services/account"
	"github.com/Michielu/go-users/app/services/passwords"
)

func Login(w http.ResponseWriter, r *http.Request, db *database.Database) {
	w.Header().Set("Content-Type", "application/json")
	var p modals.MasterUser
	var item modals.SimpleResponse

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	masterUser, err := account.GetMasterUserFromDb(db, p.Username)

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
		item = modals.SimpleResponse{Success: true, ShortMessage: "Login successful for user: " + p.Username}
	} else {
		item = modals.SimpleResponse{Success: false, ShortMessage: "Incorrect password"}
	}

	j, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	jsonData := []byte(j)
	w.Write(jsonData)
}
