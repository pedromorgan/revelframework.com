package controllers

import(
	//"fmt"
	"github.com/google/go-github/github"
)

/*
basic version that fetched all revel stuff from API
TODO need caching etc and below proof of concept..

 */

type Repo struct {


}

func GetReposList() ([]github.Repository, error) {
	client := github.NewClient(nil)
	opt := &github.RepositoryListByOrgOptions{}
	repos, _, err := client.Repositories.ListByOrg("revel", opt)
	//fmt.Println(repos)
	return repos, err
}
