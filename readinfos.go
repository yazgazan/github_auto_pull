package main

import (
	"io"
	"fmt"
	"bytes"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

type RepoType int

const (
	Github RepoType = iota
	Bitbucket
)

type RepoInfos struct {
	Name string `json:"name"`
}

type PushInfos struct {
	Ref		string    `json:"ref"`
	Repository	RepoInfos `json:"repository"`
	Type		RepoType `json:"-"`
}

type CommitBitbucket struct {
	Branch	string	`json:"branch"`
}

type PushInfosBitbucket struct {
	Repository struct {
		Name	string	`json:"name"`
	} `json:"repository"`
	Commits []CommitBitbucket `json:"commits"`
}

func ReadInfos(body io.ReadCloser, pushInfos *PushInfos) error {
	raw, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, &pushInfos)
	if err != nil {
		var bitbucket PushInfosBitbucket

		err = nil
		ss := bytes.SplitN(raw, []byte("="), 2)
		if len(ss) != 2 {
			return fmt.Errorf("couldn't parse body")
		}
		s, err := url.QueryUnescape(string(ss[1]))
		err = json.Unmarshal([]byte(s), &bitbucket)
		if err != nil {
			return err
		}
		commit := bitbucket.Commits[len(bitbucket.Commits)-1]
		pushInfos.Ref = fmt.Sprintf("refs/heads/%s", commit.Branch)
		pushInfos.Repository.Name = bitbucket.Repository.Name
		pushInfos.Type = Bitbucket
	} else {
		pushInfos.Type = Github
	}
	return err
}
