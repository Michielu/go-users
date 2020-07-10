package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Michielu/go-users/app/services/account"
	"github.com/Michielu/go-users/app/services/authentication"

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
	router.HandleFunc("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		account.GetUserHandler(w, r, db)
	})

	router.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		account.PostUserHandler(w, r, db)
	})

	router.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		authentication.Login(w, r, db)
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

// func (env *Env) NotFound(w http.ResponseWriter, req *http.Request) {
//
// }

// // MethodNotAllowed is the route for a 405 Method Not Allowed Error
// // It returns {meta: {error: true}, result: {error: NotAllowed}}
// func (env *Env) MethodNotAllowed(w http.ResponseWriter, req *http.Request) {
// }
