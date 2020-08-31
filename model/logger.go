package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

func SaveJSONFile(request *ExposureNotificationRequest) (string, error) {
	file, _ := json.MarshalIndent(request, "", " ")
	now := int(time.Now().Unix())

	// How should we include time + another ID into our log filename?
	fname := rootDir() + "/logs/" + strconv.Itoa(now) + ".json"

	err := ioutil.WriteFile(fname, file, 0644)
	if err != nil {
		log.Println(err)
	}
	return "Saved ExposureNotificationRequest in JSON file", err
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
