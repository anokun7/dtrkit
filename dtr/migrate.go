package dtr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	client   = &http.Client{}
	login    = "admin"
	password = "OrcaOrca"

	dtrapi  = "https://dtr.noop.ga/api/v0"
	enziapi = "https://dtr.noop.ga/enzi/v0"
	limits  = "limit=10000&pageSize=10000"

	accountsep     = "accounts"
	accountsfilter = "filter=users"
	orgsfilter     = "filter=orgs"
	accountsurl    = fmt.Sprintf("%s/%s?%s&%s", enziapi, accountsep, accountsfilter, limits)
	orgsurl        = fmt.Sprintf("%s/%s?%s&%s", enziapi, accountsep, orgsfilter, limits)

	reposep  = "repositories"
	reposurl = fmt.Sprintf("%s/%s?%s", dtrapi, reposep, limits)
	teamsep  = dtrapi + "/accounts/intranet/teams?limit=5000&pageSize=5000"
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

func getResource(url string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(login, password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("GetResource: URL: %s\nError: %s", url, err.Error)
	}
	resource, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error)
	}
	defer resp.Body.Close()
	return resource
}

func GetOrgs() []byte {
	return getResource(orgsurl)
}

func GetRepos() []byte {
	return getResource(reposurl)
}

func GetTeamsRepoAccess() {
	org := Account{}
	err := json.Unmarshal(GetOrgs(), &org)
	if err != nil {
		fmt.Println(err)
		return
	}
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
