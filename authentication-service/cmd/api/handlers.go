package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)


func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request)  {
	
	var requestPayload struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w,r,&requestPayload)  

	if err != nil {
		app.errorJSON(w,err,http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	user, err := app.Models.User.GetByEmail(requestPayload.Email)

	if err != nil {
		app.errorJSON(w,errors.New("invalid email credentials"),http.StatusUnauthorized)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)

	if err != nil || !valid {
		app.errorJSON(w,errors.New("invalid credentials"),http.StatusUnauthorized)
		return
	}

	payload := jsonResponse {
		Error: false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
	}
	//log authentication:
	err = app.logRequest("authentication", fmt.Sprintf("%s logged in",user.Email))
	if err != nil {
		app.errorJSON(w,err)
		return
	}
	app.writeJSON(w,http.StatusAccepted,payload)
}


func (app *Config) logRequest(name string, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry,"","\t")
	loggerURL := "http://logger-service/log"
	
	request, err := http.NewRequest("POST",loggerURL,bytes.NewBuffer(jsonData))
	if err != nil {
		
		return err
	}
	
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()


	return nil
}