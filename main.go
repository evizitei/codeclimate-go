package main

import "fmt"

func main() {
	token, repoId, cnfErr := LoadConfig()
	if cnfErr != nil {
		fmt.Println("Configuration error:", cnfErr)
		return
	}
	repo, err := QueryCodeClimate(token, repoId)
	if err != nil {
		fmt.Println("Oops, bailing...", err)
		return
	}
	fmt.Println("Current Score Is: ", repo.LastSnapshot.Gpa)
}
