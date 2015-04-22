package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	Listen string
}

type Repo struct {
	Folder		string
	Remote		string
	Branch		string
	Before		string
	After		string
	Bitbucket	bool
}

type Repos []Repo

type Handler struct {
	Config Config
	Repos  map[string]Repos
}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var pushInfos PushInfos

	if req.Method != "POST" {
		res.WriteHeader(405)
		res.Write([]byte("method not allowed"))
		log.Println("wrong method", req.Method, req.URL)
		return
	}

	err := ReadInfos(req.Body, &pushInfos)
	if err != nil {
		log.Println("couldn't read body")
		res.WriteHeader(400)
		res.Write([]byte("couldn't read body"))
	}

	repoName := pushInfos.Repository.Name
	if _, ok := h.Repos[repoName]; ok == false {
		log.Println("repo not found", repoName)
		res.WriteHeader(404)
		res.Write([]byte("not found"))
		return
	}

	for _, repo := range h.Repos[repoName] {
		h.PullRepo(repo, pushInfos)
	}
	res.WriteHeader(200)
	res.Write([]byte(""))
}
