package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go Docker Tutorial")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Yellow Machaan!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
