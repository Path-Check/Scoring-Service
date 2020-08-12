package logger

import (
	"log"
	"net/http"
	"os"
	"time"
)

type Event struct {
	Time          string
	Method        string
	RequestURI    string
	Name          string
	ExecutionTime string
	Pid           int
	Host          string
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		inner.ServeHTTP(w, r)

		e := Event{
			Time:          time.Now().String(),
			Method:        r.Method,
			RequestURI:    r.RequestURI,
			Name:          name,
			ExecutionTime: (time.Since(start)).String(),
			Pid:           os.Getpid(),
			Host:          r.Host,
		}

		log.Printf(
			"%s\t%s\t%s\t%s\t%s\t%v\t%s",
			e.Time,
			e.Method,
			e.RequestURI,
			e.Name,
			e.ExecutionTime,
			e.Pid,
			e.Host,
		)
	})
}
