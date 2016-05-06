package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mfine30/terminalboard/api"
)

func main() {
	concourseAPIPrefix := fmt.Sprintf("%s/api/v1/", os.Getenv("CONCOURSE_HOST"))
	concourseUsername := os.Getenv("CONCOURSE_USERNAME")
	concoursePassword := os.Getenv("CONCOURSE_PASSWORD")
	port := os.Getenv("PORT")

	checker := api.NewChecker(concourseAPIPrefix, concourseUsername, concoursePassword)
	router, err := api.NewRouter(checker)
	if err != nil {
		panic(err)
	}

	address := fmt.Sprintf(":%s", port)
	err = http.ListenAndServe(address, router)
	if err != nil {
		panic(err)
	}
}