
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
    return // branch not matching
  }

  shcmd := fmt.Sprintf("cd '%s';git pull '%s' '%s'", dir, remote, branch)
  cmd := exec.Command("sh", "-c", shcmd)
  cmd.Stdout = &out
  err = cmd.Run()
  if err != nil {
    log.Printf("pulling %s %s/%s : %s\n", dir, remote, branch, err)
  } else {
    log.Printf("pulling %s %s/%s : %s\n", dir, remote, branch, out.String())
  }
}

