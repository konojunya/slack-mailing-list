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
}

type User struct {
	ID       string  `json:"id"`
	TeamID   string  `json:"team_id"`
	Name     string  `json:"name"`
	RealName string  `json:"real_name"`
	Profile  Profile `json:"profile"`
}

type UserInfo struct {
	User User `json:"user"`
}

var (
	Credential *CredentialInfo
)

func SetConfig(credential *CredentialInfo) {
	Credential = credential
}

func GetUserInfo() UserInfo {
	return Credential.getUserInfo()
}

func (cre *CredentialInfo) getUserInfo() UserInfo {
	token := cre.AccessToken
	userID := cre.UserID
	resp, err := http.Get("https://slack.com/api/users.info?token=" + token + "&user=" + userID)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	var info UserInfo
	json.Unmarshal(bytes, &info)
	return info
}
