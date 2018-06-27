package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/CA19Creators/slack-mailing-list/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func main() {
	r := gin.Default()

	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")
	r.Static("/images", "./public/images")
	r.LoadHTMLGlob("views/*")

	r.GET("/", func(c *gin.Context) {
		if service.Authed() {
			c.HTML(http.StatusOK, "index.html", nil)
			return
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.GET("/auth", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://slack.com/oauth/authorize?client_id="+clientID+"&scope=chat:write:user users:read channels:read")
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

		c.Redirect(http.StatusTemporaryRedirect, "/")

	})
	r.GET("/api/users", func(c *gin.Context) {
		nextCursor := c.Param("next_cursor")
		userList, err := service.GetUsers(nextCursor)
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		if userList.Ok {
			c.JSON(http.StatusOK, userList)
		} else {
			log.Println(userList.Error)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	})
	r.GET("/api/channels", func(c *gin.Context) {
		nextCursor := c.Param("next_cursor")
		channelList, err := service.GetChannels(nextCursor)
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		if channelList.Ok {
			c.JSON(http.StatusOK, channelList)
		} else {
			log.Println(channelList.Error)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	})
	r.POST("/api/message", func(c *gin.Context) {
		var req struct {
			Text string   `json:"text"`
			IDs  []string `json:"ids"`
		}
		if err := c.BindJSON(&req); err == nil {
			for _, id := range req.IDs {
				_ = service.PostMessage(req.Text, id)
			}
			c.AbortWithStatus(http.StatusOK)
			return
			// if apiError != nil {
			// 	log.Println(apiError)
			// 	c.AbortWithStatus(http.StatusInternalServerError)
			// 	return
			// } else {
			// 	c.AbortWithStatus(http.StatusCreated)
			// 	return
			// }
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
		}
	})

	r.Run(":7000")
}
