package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		uid := uuid.New().String()
		fmt.Printf("All good 👍 %v\n", uid)
		fmt.Fprintf(w, "All good 👍 %v\n", uid)
	})
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
