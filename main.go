package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	playpause "playpauseapi/playplause"
	"sync"

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

func getParameters(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	var wg sync.WaitGroup
	input := req.Command
	command := make(chan string)
	response := make(chan string)

	if req.Command == "Pause" {
		command <- req.Command
	} else {
		fmt.Println("input", input)
		go playpause.Routine(command, response, &wg)
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
