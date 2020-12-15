package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	playpause "playpauseapi/playplause"

	"github.com/gorilla/mux"
)

//Request array
type Request struct {
	Command string
}

//Response
type Response struct {
	Status string
}

var (
	command  chan string
	response chan string
	// runningChans = map[string]chan string
)

func init() {
}

func getParameters(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	input := req.Command
	if req.Command == "Pause" || req.Command == "Stop" {
		command <- req.Command
	} else {
		fmt.Println("input", input)
		response = make(chan string)
		command = make(chan string)
		// runningChans["test_name"] = command
		go playpause.Routine(command, response)
	}
	tmp := <-response
	fmt.Println(tmp)
	res := Response{
		Status: tmp,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	fmt.Println("Server Running")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/start", getParameters)
	log.Fatal(http.ListenAndServe(":8000", router))
}
