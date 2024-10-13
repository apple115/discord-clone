package github

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func GetClient() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     "Ov23liaNpf9LOe4Ov2Eo",
		ClientSecret: "eae47247e28ef13ca36ae36b2467715218f0d3a9",
		RedirectURL:  "http://127.0.0.1:8000/callback",
		Scopes:       []string{"user"},
		Endpoint:     github.Endpoint,
	}
}
