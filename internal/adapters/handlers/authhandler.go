package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sugandhasaxena1911/MyPracticeAuthApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/dto"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/service"
)

type AuthHandler struct {
	Authservice service.LoginCoreService
}

func (authhandler AuthHandler) FetchLogindetails(w http.ResponseWriter, r *http.Request) {
	var logindto dto.LoginDto
	err := json.NewDecoder(r.Body).Decode(&logindto)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		er := error.NewBadRequestAppError("Invalid Request")
		json.NewEncoder(w).Encode(er.Getmessage())
		return
	}
	authtoken, e := authhandler.Authservice.FetchLoginDetails(logindto)
	if e != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(e.Code)
		json.NewEncoder(w).Encode(e.Getmessage())
		return

	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authtoken)
}
