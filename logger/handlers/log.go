package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"log/model"
	"log/persistence"
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
		log.Println("Request Unmarshal Err %v", err)
	}
	f, err := persistence.OpenFile()
	// result, err := persistence.SaveRequestToFile(f, req)
	
}
