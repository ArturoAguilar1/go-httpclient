package main

import (
	"fmt"
	"github.com/ArturoAguilar1/go-httpclient/gohttp"
	"io/ioutil"
)

var (
	githubHttpClient = getGithubClient()
)

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func getGithubClient() gohttp.Client {
	client := gohttp.NewBuilder().
		DisableTimeouts(true).
		SetMaxIdleConnections(5).
		Build()

	return client
}

func main() {
	getUrls()
}

func getUrls() {
	url := "https://api.github.com"
	response, err := githubHttpClient.Get(url,nil)

	if err != nil{
		panic(err)
	}
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))
}

func createUser(user User) {
	url := "https://api.github.com"
	response, err := githubHttpClient.Post(url, nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}