package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var MainRouter *mux.Router

func init() {
	log.Println("Initialized")
	MainRouter = mux.NewRouter()
	MainRouter.PathPrefix("/").Handler(http.HandlerFunc(allHandler))
}

func main() {
	var err error

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}

	log.Println("Server started at port:", port)
	err = http.ListenAndServe(
		port,
		MainRouter,
	)

	if err != nil {
		log.Fatal("There was error when starting HTTP server", err)
	}

}

func allHandler(w http.ResponseWriter, r *http.Request) {


	clientIp := r.Header.Get("X-Forwarded-For")

	if clientIp == "" {
		clientIp = r.Header.Get("X-Real-Ip")
	}

	contentLength := r.Header.Get("Content-Length")
	if contentLength != "" {
		contentSize, errContentSize := strconv.Atoi(contentLength)
		if errContentSize != nil {
			log.Println(contentSize)
		}
	}

	addr := r.RemoteAddr
	if i := strings.LastIndex(addr, ":"); i != -1 {
		addr = addr[:i]
	}

	var datab []byte
	var data string

	if r.Method == "POST" {
		var request map[string]interface{}

		err := json.NewDecoder(r.Body).Decode(&request)

		if err == nil {
			datab, err = json.Marshal(request)
			data = string(datab)
		}
	}

	fmt.Printf(`
========================================
	addr:     %s
	clientIp: %s
	method:   %s
	URL:      %s
	data:     %s
****************************************`,
	addr, clientIp, r.Method, r.URL, data)

	w.WriteHeader(http.StatusAccepted)
}
