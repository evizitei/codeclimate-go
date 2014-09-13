package main

import "code.google.com/p/gcfg"

type Config struct {
	Codeclimate struct {
		Token string
		Repo  string
	}
}

func LoadConfig() (string, string, error) {
	var config Config
	cnfErr := gcfg.ReadFileInto(&config, "codeclimate.gcfg")
	if cnfErr != nil {
		return "", "", cnfErr
	}
	token := config.Codeclimate.Token
	repoId := config.Codeclimate.Repo
	return token, repoId, nil
}
