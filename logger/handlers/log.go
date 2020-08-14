package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"logger/model"
	"logger/persistence"
	"net/http"
)

func Log(w http.ResponseWriter, r *http.Request) {
	req := model.LogRequest{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 9999999))
	if err != nil {
		log.Printf("Request Body Read Err: %v\n", err)
	}
	defer r.Body.Close()
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("Request Unmarshal Err: %v\n", err)
	}
	f, err := persistence.OpenFile()
	if err != nil {
		log.Printf("File Open Error: %v\n", err)
	}
	logResult, err := persistence.SaveRequestToFile(f, req)
	if err != nil {
		log.Printf("Log Save Error: %v\n", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(logResult)
	if err != nil {
		log.Printf("Json Encoding Error: %v\n", err)
	}
	_, err = persistence.CloseFile(f)
	if err != nil {
		log.Printf("Log Close Error: %v:\n", err)
	}
}
