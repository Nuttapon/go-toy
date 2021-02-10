package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FundContainer struct {
	ClientReferenceID string `json:"id"`
	TypeOfFund        string `json:"type"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to the Homepage!</h1>")
}

func getAllFundContainers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(FundContainers)
}

func handleRequest() {
	http.HandleFunc("/", getAllFundContainers)
	http.ListenAndServe(":8080", nil)
}

// FundContainers is exportable
var FundContainers []FundContainer

func main() {
	FundContainers = []FundContainer{
		{ClientReferenceID: "1", TypeOfFund: "Truemoney"},
		{ClientReferenceID: "2", TypeOfFund: "Truemoney"},
	}

	handleRequest()
}
