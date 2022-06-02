package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	pb "github.com/LordRadamanthys/grpc_profile_github/pb"
)

type Server struct{}

type Profile struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Login      string `json:"login"`
	AvatarURL  string `json:"avatar_url"`
	Location   string `json:"location"`
	Followers  int64  `json:"followers"`
	Following  int64  `json:"following"`
	Repos      int64  `json:"public_repos"`
	Gists      int64  `json:"public_gists"`
	URL        string `json:"url"`
	StarredURL string `json:"starred_url"`
	ReposURL   string `json:"repos_url"`
}

//same function in the proto file
func (s *Server) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("Receive message from client: %s", in.Username)

	// make a request to github api
	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%v", in.Username))

	if err != nil {
		log.Fatal(err)
	}

	// try to parse body in to bytes
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	githubUser := Profile{}
	// try to parse bytes from body in struct
	if jsonErr := json.Unmarshal(body, &githubUser); jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return &pb.UserResponse{
		Id:        githubUser.ID,
		Name:      githubUser.Name,
		Username:  githubUser.Username,
		Avatarurl: githubUser.AvatarURL,
		Location:  githubUser.Location,
		Statistics: &pb.Statistics{
			Followers: githubUser.Followers,
			Following: githubUser.Following,
			Repos:     githubUser.Repos,
			Gists:     githubUser.Gists,
		},
		ListURLs: []string{githubUser.URL, githubUser.StarredURL, githubUser.ReposURL},
	}, nil
}
