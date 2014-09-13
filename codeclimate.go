package main

import (
	"encoding/json"
	"net/http"
)

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
		return Repository{}, err
	}
	defer resp.Body.Close()
	var repo Repository
	jsonErr := json.NewDecoder(resp.Body).Decode(&repo)

	if jsonErr != nil {
		return Repository{}, err
	}
	return repo, nil
}
