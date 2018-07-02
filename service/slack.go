package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
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
	Error            string           `json:"error"`
}

type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ChannelList struct {
	Ok               bool             `json:"ok"`
	Channels         []Channel        `json:"channels"`
	ResponseMetadata ResponseMetadata `json:"response_metadata"`
	Error            string           `json:"error"`
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

func GetChannels(nextCursor string) (*ChannelList, error) {
	return Credential.getChannels(nextCursor)
}

func PostMessage(text, channel string) error {
	return Credential.postMessage(text, channel)
}

func (cre *CredentialInfo) postMessage(text, channel string) error {
	token := cre.AccessToken
	values := url.Values{}

	values.Add("token", token)
	values.Add("channel", channel)
	values.Add("text", text)
	values.Add("as_user", "true")
	req, err := http.NewRequest(
		"POST",
		"https://slack.com/api/chat.postMessage",
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (cre *CredentialInfo) getChannels(nextCursor string) (*ChannelList, error) {
	token := cre.AccessToken
	resp, err := http.Get(fmt.Sprintf("https://slack.com/api/channels.list?token=%s&cursor=%s", token, nextCursor))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	var channelList ChannelList
	json.Unmarshal(bytes, &channelList)

	return &channelList, nil
}

func (cre *CredentialInfo) getUsers(nextCursor string) (*UserList, error) {
	token := cre.AccessToken
	resp, err := http.Get(fmt.Sprintf("https://slack.com/api/users.list?token=%s&cursor=%s", token, nextCursor))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	var userList UserList
	json.Unmarshal(bytes, &userList)

	return &userList, nil
}
