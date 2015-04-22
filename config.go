package main

import (
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

func ReadConfigs(handler *Handler) {
	reposConfig, err := ioutil.ReadFile("repos.toml")
	if err != nil {
		log.Fatal("couldn't read repos config file ", err)
	}

	config, err := ioutil.ReadFile("config.toml")
	if err != nil {
		log.Fatal("couldn't read config file ", err)
	}

	_, err = toml.Decode(string(reposConfig), &handler.Repos)
	if err != nil {
		log.Fatal("couldn't read repos config file ", err)
	}
	if _, ok := handler.Repos["config"]; ok == true {
		delete(handler.Repos, "config")
	}
	_, err = toml.Decode(string(config), &handler.Config)
	if err != nil {
		log.Fatal("couldn't read config file ", err)
	}
}
