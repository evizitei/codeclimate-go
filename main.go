package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"code.google.com/p/gcfg"
)

type Config struct {
	Codeclimate struct {
		Token string
		Repo  string
	}
}
type Repository struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	LastSnapshot struct {
		Id  string  `json:"id"`
		Sha string  `json:"commit_sha"`
		Gpa float64 `json:"gpa"`
	} `json:"last_snapshot"`
}

func QueryCodeClimate(token string, repoId string) (Repository, error) {
	url := "https://codeclimate.com/api/repos/" + repoId + "?api_token=" + token
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in API request")
		return Repository{}, err
	}
	defer resp.Body.Close()
	var repo Repository
	jsonErr := json.NewDecoder(resp.Body).Decode(&repo)

	if jsonErr != nil {
		fmt.Println("Error in JSON decoding: ", jsonErr)
		return Repository{}, err
	}
	return repo, nil
}

func LoadCodeclimateConfig() (string, string) {
	var config Config
	cnfErr := gcfg.ReadFileInto(&config, "codeclimate.gcfg")
	if cnfErr != nil {
		fmt.Println("Configuration error...", cnfErr)
	}
	token := config.Codeclimate.Token
	repoId := config.Codeclimate.Repo
	return token, repoId
}

func main() {
	token, repoId := LoadCodeclimateConfig()
	repo, err := QueryCodeClimate(token, repoId)
	if err != nil {
		fmt.Println("Oops, bailing...", err)
	}
	fmt.Println("Current Score Is: ", repo.LastSnapshot.Gpa)
}
