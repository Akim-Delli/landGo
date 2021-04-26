package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const GITHUB_URL = "https://api.github.com/users/akim-delli"

type User struct {
	Name       string `json:"name"`
	PublicRepo int    `json:"public_repos"`
}

func main() {
	dec, err := userInfo("akim-delli")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	log.Printf("%+v\n", dec)
}

func userInfo(login string) (*User, error) {
	resp, err := http.Get(GITHUB_URL)
	if err != nil {
		log.Fatalf("Error Fetching %s : %v", GITHUB_URL, err)
	}

	defer resp.Body.Close()

	// Decode JSON
	user := &User{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(user); err != nil {
		return nil, err
	}

	return user, nil

}
