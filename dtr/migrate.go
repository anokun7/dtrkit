package dtr

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

var (
  client = &http.Client{}
  dtrapi = "https://dtr.noop.ga/api/v0"
  enziapi = "https://dtr.noop.ga/enzi/v0"
)

type Account struct {
  Accounts []struct {
    Name     string `json:"name"     description:"Name of the account"`
    ID       string `json:"id"       description:"ID of the account"`
    FullName string `json:"fullName" description:"Full Name of the account"`
    IsOrg    bool   `json:"isOrg"    description:"Whether the account is an organization (or user)"`

    // Fields for users only.
    IsAdmin    *bool `json:"isAdmin,omitempty"    description:"Whether the user is a system admin (users only)"`
    IsActive   *bool `json:"isActive,omitempty"   description:"Whether the user is active and can login (users only)"`
    IsImported *bool `json:"isImported,omitempty" description:"Whether the user was imported from an upstream identity provider"`

    // Fields for orgs only.
    MembersCount *int `json:"membersCount,omitempty" description:"The number of members of the organization"`
  } `json:"accounts"`
}

type Team struct {
  Teams []struct {
    OrgID        string `json:"orgID"        description:"ID of the organization to which this team belongs"`
    Name         string `json:"name"         description:"Name of the team"`
    ID           string `json:"id"           description:"ID of the team"`
    Description  string `json:"description"  description:"Description of the team"`
    MembersCount int    `json:"membersCount" description:"The number of members of the team"`
  } `json:"teams"`
}

func GetOrgs() []byte {
  endpoint := enziapi+"/accounts?filter=orgs&limit=10000&pageSize=10000"
  req, _ := http.NewRequest("GET", endpoint, nil)
  req.SetBasicAuth("admin", "OrcaOrca")
  resp, err := client.Do(req)
  if err != nil {
    fmt.Printf("GetOrgs: %s", err)
  }
  orgs, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  return orgs
}

func GetRepos() {
  endpoint := dtrapi+"/repositories"
  req, _ := http.NewRequest("GET", endpoint, nil)
  req.SetBasicAuth("admin", "OrcaOrca")
  resp, err := client.Do(req)
  if err != nil {
    fmt.Printf("GetRepos: %s", err)
  }
  json, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  println(string(json))
}

func GetTeamsRepoAccess() {
  orgs := GetOrgs()
  println(string(orgs))
  org := Account{}
  err := json.Unmarshal(orgs, &org)
  if err != nil {
    fmt.Println(err)
    return
  }

  //  if err != nil {
  //    fmt.Println("====== In Get Teams Repo Access =====")
  //    println(org.Account[0].Name)
  //  } else {
  //    println("GetTeamsRepoAccess: "+ err.Error())
  //  }
  fmt.Println(len(org.Accounts))
  println(org.Accounts[0].Name)
}

// 
// func GetMembers(dtr URL) {
// }
// 
// func GetNamespaces(dtr URL) {
// }
// 
// func GetTagsPerRepo(dtr URL) {
// }
// 
// func GetImageTags(dtr URL) {
// }
// 
// func PullImages(dtr URL) {
// }
// 
// func CreateNameSpaces(dtr URL) {
// }
// 
// func CreateOrgs(dtr URL) {
// }
// 
// func CreateTeams(dtr URL) {
// }
// 
// func CreateMembers(dtr URL) {
// }
// 
// func CreateRepos(dtr URL) {
// }
// 
// func PushImages(dtr URL) {
// }
