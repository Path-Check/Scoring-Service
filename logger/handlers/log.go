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
		log.Println("Request Body Read Err: %v", err)
	}
	defer r.Body.Close()
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("Request Unmarshal Err: %v", err)
	}
	f, err := persistence.OpenFile()
	if err != nil {
		log.Println("File Open Error: %v", err)
	}
	logResult, err := persistence.SaveRequestToFile(f, req)
	if err != nil {
		log.Println("Log Save Error: %v", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(logResult)
	if err != nil {
		log.Println("Json Encoding Error: %v", err)
	}
	_, err = persistence.CloseFile(f)
	if err != nil {
		log.Println("Log Close Error: %v:", err)
	}
}
