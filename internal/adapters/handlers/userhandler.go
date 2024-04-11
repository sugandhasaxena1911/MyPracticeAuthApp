package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sugandhasaxena1911/MyPracticeAuthApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/dto"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/ports/service"
)

func GetTestAuth(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello Test Authorization Page"))

}

type UserHandler struct {
	Usrservice service.UserService
}

func (usrhandler UserHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	log.Println("inside Handler PostUser")
	userdto := dto.User{}
	e := json.NewDecoder(r.Body).Decode(&userdto)
	if e != nil {
		w.Header().Add("Content-Type", "application/json")
		log.Println(e.Error())
		er := error.NewBadRequestAppError("Invalid request")
		w.WriteHeader(er.Code)
		json.NewEncoder(w).Encode(er.Getmessage())
		return
	}
	log.Println("calling service register user ")

	userdto, err := usrhandler.Usrservice.RegisterUser(userdto)
	if err != nil {
		log.Println(err.Message)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err.Getmessage())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userdto)

}
