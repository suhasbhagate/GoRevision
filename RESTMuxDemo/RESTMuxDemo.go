package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	UserId   int
	UserName string
}

var userList = []User{{11, "Saksham"}, {12, "Suhas"}}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", servehome).Methods("GET")
	r.HandleFunc("/getallusers", getallusers).Methods("GET")
	r.HandleFunc("/getuser/{id}", getuser).Methods("GET")
	r.HandleFunc("/getuserqp", getuserqp).Methods("GET")
	r.HandleFunc("/adduser", adduser).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:2100", r))
}

func servehome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome")
	//w.Write([]byte("Welcome"))
}

func getallusers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(userList)
}

func getuser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idstr := param["id"]
	id, _ := strconv.Atoi(idstr)

	for _, v := range userList {
		if id == v.UserId {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode("Element Not Found")
}

func getuserqp(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idstr)

	for _, v := range userList {
		if id == v.UserId {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode("Element Not Found")
}

func adduser(w http.ResponseWriter, r *http.Request) {
	var usr User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		log.Println(err)
		return
	}
	userList = append(userList, usr)
	json.NewEncoder(w).Encode(userList)
}
