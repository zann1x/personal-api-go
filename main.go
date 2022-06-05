package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handleIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type IndexResponse struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Pronouns  string `json:"pronouns"`
	}

	res := IndexResponse{
		FirstName: "lukas",
		LastName:  "zanner",
		Pronouns:  "he/him",
	}
	json_res, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json_res)
}

func main() {
	router := httprouter.New()
	router.GET("/", handleIndex)

	log.Fatal(http.ListenAndServe(":8080", router))
}
