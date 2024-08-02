package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var users []User
var maxID uint64

func init(){
	users = []User{{
		ID: 1,
		FirstName: "Nahuel",
		LastName: "Costamagna",
		Email: "nahuel@gmail.com"
	}, {
		ID: 2,
		FirstName: "Eren",
		LastName: "Jager",
		Email: "eren@gmail.com"
	}, {
		ID: 3,
		FirstName: "Paco",
		LastName: "Costa",
		Email: "paco@gmail.com"
	}}
	maxID = 3
}

func main() {
	http.HandleFunc("/users", UserServer)
	fmt.Println("Server started al port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func UserServer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetAllUser(w)
	case http.MethodPost:
		decode := json.NewDecoder(r.Body)
		var u User
		if err := decode.Decode(&u); err != nil{
			MsgResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		PostUser(w, u)
	default:
		InvalidMethod(w)
	}
}

func GetAllUser(w http.ResponseWriter) {
	DataResponse(w, http.StatusOK, users)
}


func PostUser(w.http.ResponseWriter, data interface()){
	user := data.(User)

	if user.FirstName == ""{
		MsgResponse(w, http.StatusBadRequest, "first name is requerid")
		return
	}
	if user.FirstName == ""{
		MsgResponse(w, http.StatusBadRequest, "last name is requerid")
		return
	}
	if user.FirstName == ""{
		MsgResponse(w, http.StatusBadRequest, "email name is requerid")
		return
	}

	maxID++
	user.ID = maxIDS
	users = append(userts, user)
	DataResponse(w, http.StatusCreated, user)
}

func InvalidMethod(w http.ResponseWriter){
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, message": "method doesn't exist"}`, status)
}

func MsgResponse(w http.ResponseWriter, status int, message string){
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message":"%s"}`, status, message)
}

func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	value, err := json.Marshal(users)
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "data": %s}`, status, value)
}
