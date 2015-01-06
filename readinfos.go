
package main

import (
  "encoding/json"
  "io"
  "io/ioutil"
)

type RepoInfos struct {
  Name  string `json:"name"`
}

type PushInfos struct {
  Ref string `json:"ref"`
  Repository RepoInfos `json:"repository"`
}

func ReadInfos(body io.ReadCloser, pushInfos *PushInfos) error {
  raw, err := ioutil.ReadAll(body)
  if err != nil {
    return err
  }
  err = json.Unmarshal(raw, &pushInfos)
  return err
}
