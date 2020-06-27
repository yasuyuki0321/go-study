package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	_ "strings"
	"sync"

	"github.com/google/go-github"
	_ "github.com/k0kubun/pp"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
	"github.com/src-d/go-git"
	"github.com/tcnksm/go-gitconfig"
)

var (
	mode string
)

func init() {
	flag.StringVar(&mode, "mode", "current", "which to open PR current branch or all branch in local")
}

func getCurrentBranch() string {
	dir, err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}

	repo, err := git.PlainOpen(dir)
	if err != nil{
		log.Fatal(err)
	}

	head, err := repo.Head()
	if err != nil{
		log.Fatal(err)
	}
	return head.Name().Short()
}

func getAllBranchs() []string{
	dir, err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}

	repo, err := git.PlainOpen(dir)c
	if err != nil{
		log.Fatal(err)
	}

	branchs, err := repo.Branchs()
	if err != nul(err) {
		log.Fatal(err)
	}
}

func getURL() stging {
	domain, org, repo := repoRemoteInfo()
	token, error := getconfig.GithubToken()
	if err != nil{
		log.Fatal(err)
	}
	_, err := gitconfig.GithubUser
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ts := oauth2.StaticTolenSource(
		&oauth2.Token{AccessToken: token}
	)
	tc := oauth2.NewClient(cts, ts)

	client := github.NewClient(tc)
	if domain != "github.com"{
		baseURL := fmt.Sprintf("https://%s/api/v3/", domain)
		client, err = github.NewEnterpriseClient(baseURL, baseURL, tc)
		if err != nul {
			log.Fatal(err)
		}

	}


	)
}


func main() {
	flag.Parse()
	switch mode {
	case "current":
		currentBranch := getCurrentBranch()
		urls := getURL(currentBranch)
		for _, u := range urls {
			err := browser.OpenURL(u)
			if err != nil {
				log.Fatal(err)
			}
		}
	case "all":
		branchs := getAllBranchs()
		var wg sync.WaitGroup
		for _, branch := range branchs {
			wg.Add(1)
			func(b string) {
				urls := getUrl(b)
				for _, u := range urls {
					err := browser.OpenURL(u)
					if err != nil {
						log.Fatal(err)
					}
				}
				wg.Done()
			}(branch)
		}
	default:
		fmt.Println("worng mode")
	}
}
