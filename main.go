package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const apiURL = "https://catfact.ninja/fact"

type profile struct {
	Status    string `json:"status"`
	User      user   `json:"user"`
	Timestamp string `json:"timestamp"`
	Fact     string `json:"fact"`
}

type user struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Stack string `json:"stack"`
}

type catFactResponse struct {
	Fact string `json:"fact"`
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(apiURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to make API request: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var respData catFactResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode API response: %v", err), http.StatusInternalServerError)
		return
	}

	profileData := profile{
		Status: "success",
		User: user{
			Email: "ebitegift235@gmail.com",
			Name:  "Godsgift Uloamaka Ebite",
			Stack: "Go/Gin",
		},
		Timestamp: time.Now().UTC().Truncate(time.Millisecond).Format("2006-01-02T15:04:05.000Z"),
		Fact:     respData.Fact,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profileData)
}

func main() {
	err := godotenv.Load()
    if err != nil {
    	fmt.Println("Error loading .env file")
    }

	http.HandleFunc("/me", getProfile)

	port := os.Getenv("PORT")
	fmt.Println("Port from env:", port)

	if port == "" {
		port = "8080" 
	}

	fmt.Println("Server is running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
