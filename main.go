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
	wg.Add(1)
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

// func Pause(w http.ResponseWriter, r *http.Request) {
// 	var req Request
// 	json.NewDecoder(r.Body).Decode(&req)
// 	command := make(chan string)
// 	response := make(chan string)
// 	fmt.Println(req.Command)
// 	command <- req.Command
// 	tmp := <-response
// 	fmt.Println(tmp)
// 	res := Response{
// 		Status: tmp,
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(res)
// }

func main() {
	fmt.Println("Go Docker Tutorial")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/start", getParameters)
	// router.HandleFunc("/pause", Pause)
	log.Fatal(http.ListenAndServe(":8000", router))
}
