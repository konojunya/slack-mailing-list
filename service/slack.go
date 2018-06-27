package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CredentialInfo struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	UserID      string `json:"user_id"`
	TeamName    string `json:"team_name"`
	TeamID      string `json:"team_id"`
}

type Profile struct {
	DisplayName   string `json:"display_name"`
	ImageOriginal string `json:"image_original"`
	Image192      string `json:"image_192"`
}

type User struct {
	ID        string  `json:"id"`
	TeamID    string  `json:"team_id"`
	Name      string  `json:"name"`
	RealName  string  `json:"real_name"`
	Profile   Profile `json:"profile"`
	IsBot     bool    `json:"is_bot"`
	IsAppUser bool    `json:"is_app_user"`
}

type UserList struct {
	Ok               bool             `json:"ok"`
	Members          []User           `json:"members"`
	ResponseMetadata ResponseMetadata `json:"response_metadata"`
}

type ResponseMetadata struct {
	NextCursor string `json:"next_cursor"`
}

var (
	Credential *CredentialInfo
)

func Authed() bool {
	return Credential != nil
}

func SetConfig(credential *CredentialInfo) {
	Credential = credential
}

func GetUsers(nextCursor string) (*UserList, error) {
	return Credential.getUsers(nextCursor)
}

func (cre *CredentialInfo) getUsers(nextCursor string) (*UserList, error) {
	token := cre.AccessToken
	resp, err := http.Get("https://slack.com/api/users.list?token=" + token + "&cursor=" + nextCursor)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	var userList UserList
	json.Unmarshal(bytes, &userList)

	return &userList, nil
}
