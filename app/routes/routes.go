package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Michielu/go-users/app/database"
	"github.com/Michielu/go-users/app/modals"

	"github.com/gorilla/mux"
)

func LoadRoutes(db *database.Database) *mux.Router {
	router := mux.NewRouter()

	// HTTP Errors
	// router.NotFound = http.HandleFunc("/", NotFound)
	// router.MethodNotAllowed = http.HandleFunc("/", MethodNotAllowed)

	//Users
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		getUserHandler(w, r, db)
	})

	router.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		postUserHandler(w, r, db)
	})

	router.HandleFunc("/{anything}", DefaultHandler)

	return router
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	t := modals.SimpleResponse{Success: false, ShortMessage: "invalid url: " + vars["anything"]}
	j, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	jsonData := []byte(j)
	w.Write(jsonData)
}

func getUserHandler(w http.ResponseWriter, r *http.Request, db *database.Database) {
	log.Println("routes.getUserHandler")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	// log.Println("Vars: ", vars, vars["id"])

	item, _ := db.GetMasterUser(vars["id"])

	j, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	jsonData := []byte(j)
	w.Write(jsonData)
}

func postUserHandler(w http.ResponseWriter, r *http.Request, db *database.Database) {
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

// func (env *Env) NotFound(w http.ResponseWriter, req *http.Request) {
//
// }

// // MethodNotAllowed is the route for a 405 Method Not Allowed Error
// // It returns {meta: {error: true}, result: {error: NotAllowed}}
// func (env *Env) MethodNotAllowed(w http.ResponseWriter, req *http.Request) {
// }
