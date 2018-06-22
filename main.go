package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/konojunya/slack-oauth/service"
)

var (
	clientID     string
	clientSecret string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
}

const COOKIE_NAME = "slack-mailing-list-authed"

func main() {
	r := gin.Default()

	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")
	r.LoadHTMLGlob("views/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://slack.com/oauth/authorize?client_id="+clientID+"&scope=chat:write:user%20users:read")
	})
	r.GET("/oauth", func(c *gin.Context) {
		code := c.Request.URL.Query().Get("code")

		values := url.Values{}
		values.Add("code", code)
		values.Add("client_id", clientID)
		values.Add("client_secret", clientSecret)

		req, err := http.NewRequest(
			"POST",
			"https://slack.com/api/oauth.access",
			strings.NewReader(values.Encode()),
		)
		if err != nil {
			panic(err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)

		var cre *service.CredentialInfo
		json.Unmarshal(byteArray, &cre)
		service.SetConfig(cre)

		c.SetCookie(
			COOKIE_NAME,
			"true",
			1000*60,
			"/",
			"127.0.0.1",
			false,
			false,
		)
		c.Redirect(http.StatusTemporaryRedirect, "/")

	})

	r.Run(":7000")
}
