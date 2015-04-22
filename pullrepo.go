
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func (h Handler) PullRepo(repo Repo, pushInfos PushInfos) {
	var out bytes.Buffer

	dir := repo.Folder
	remote := repo.Remote
	branch := repo.Branch

	if repo.Bitbucket == true && pushInfos.Type != Bitbucket {
		log.Println("repo types don't match")
		return
	}
	if repo.Bitbucket == false && pushInfos.Type != Github {
		log.Println("repo types don't match")
		return
	}
	info, err := os.Stat(dir)
	if err != nil {
		log.Println("folder", dir, "not found", err)
		return
	}
	if info.IsDir() == false {
		log.Println("folder", dir, "is not a directory")
		return
	}

	if pushInfos.Ref != fmt.Sprintf("refs/heads/%s", branch) {
		log.Println("branch", pushInfos.Ref, "does not match", branch)
		return
	}

	beforecmd := fmt.Sprintf("cd '%s';%s", dir, repo.Before)
	aftercmd := fmt.Sprintf("cd '%s';%s", dir, repo.After)
	if len(repo.Before) != 0 {
		cmdbefore := exec.Command("sh", "-c", beforecmd)
		cmdbefore.Stdout = &out
		err = cmdbefore.Run()
		if err != nil {
			log.Printf("pre-script %s %s/%s : %s\n", dir, remote, branch, err)
			return
		}
	}
	shcmd := fmt.Sprintf("cd '%s';git pull '%s' '%s'", dir, remote, branch)
	cmd := exec.Command("sh", "-c", shcmd)
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Printf("pulling %s %s/%s : %s\n", dir, remote, branch, err)
		return
	}
	if len(repo.After) != 0 {
		cmdafter := exec.Command("sh", "-c", aftercmd)
		cmdafter.Stdout = &out
		err = cmdafter.Run()
		if err != nil {
			log.Printf("post-script %s %s/%s : %s\n", dir, remote, branch, err)
			return
		}
	}
	log.Printf("pulling %s %s/%s : %s\n", dir, remote, branch, out.String())
}

