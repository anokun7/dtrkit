package util

import (
  "net/http"
  "io/ioutil"
)

func ShowRepos() {
  client := &http.Client{}

  req, err := http.NewRequest("GET","https://dtr.noop.ga/api/v0/repositories", nil)
  req.SetBasicAuth("admin", "OrcaOrca")
  resp, err := client.Do(req)
  if err != nil {
  }
  json, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  println(string(json))
}
