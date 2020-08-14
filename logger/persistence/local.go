package persistence

import (
	"encoding/json"
	"fmt"
	"log"
	"log/model"
	"os"
	"time"
)

// open a file, keep it open, then save events to log

// OpenFile opens the file with the hostname and timestamp as a file
func OpenFile() (*os.File, error) {
	hostname, _ := os.Hostname()
	timestamp := time.Now().String()
	filename := fmt.Sprintf(hostname + "-" + timestamp + ".json")
	l, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// SaveToFile is pretty self explanatory
func SaveToRequestToFile(f *os.File, req model.LogRequest) (bool, error) {
	reqm, err := json.Marshal(req)
	_, err = f.Write(reqm)
	if err != nil {
		log.Println("Request File Write Error: %v", err)
		return false, err
	}
	return true, nil
}

// SaveToFile is pretty self explanatory
func SaveToResponseToFile(f *os.File, res model.LogResponse) (bool, error) {
	resm, err := json.Marshal(res)
	_, err = f.Write(resm)
	if err != nil {
		log.Println("Request File Write Error: %v", err)
		return false, err
	}
	return true, nil
}

// SaveToFile is pretty self explanatory
func SaveToFile(f *os.File, req model.LogRequest, res model.LogResponse) (bool, error) {
	reqResult, err := SaveToRequestToFile(f, req)
	resResult, err := SaveToResponseToFile(f, res)
	if reqResult && resResult == true {
		return true, nil
	}
	return false, err
}

// CloseFile do you need a dictionary?
func CloseFile(f *os.File) (bool, error) {
	hostname, _ := os.Hostname()
	timestamp := time.Now().String()
	filename := fmt.Sprintf(hostname + "-" + timestamp + ".json")
	if f.Name() != filename {
		return true, f.Close()
	}
	return false, nil
}
