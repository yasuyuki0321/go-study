package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v32/github"
	"github.com/tcnksm/go-gitconfig"
	"golang.org/x/oauth2"
	"gopkg.in/src-d/go-git.v4"
)

var (
	mode string
)

func init() {
	flag.StringVar(&mode, "mode", "current", "which to open PR current branch or all branch in local")
}

// 現在のブランチ名を取得する
func getCurrentBranch() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := git.PlainOpen(dir)
	if err != nil {
		log.Fatal(err)
	}

	head, err := repo.Head()
	if err != nil {
		log.Fatal(err)
	}

	return head.Name().Short()
}

// レポジトリ情報の取得
func repoRemoteInfo() (string, string, string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	r, err := git.PlainOpen(dir)
	if err != nil {
		log.Fatal(err)
	}

	remote, err := r.Remote("origin")
	if err != nil {
		log.Fatal(err)
	}

	if len(remote.Config().URLs) != 1 {
		log.Fatal("origin url is not only one!")
	}

	url := remote.Config().URLs[0]
	u := strings.Split(url, "@")[0]
	v := strings.SplitN(u, ":", 2)

	domain, path := v[0], v[1]
	v = strings.SplitN(path, "/", 2)
	org, repo := v[0], strings.TrimRight(v[1], ".git")

	return domain, org, repo
}

// URLの取得
func getURL(branch string) {
	domain, org, repo := repoRemoteInfo()

	token, err := gitconfig.GithubToken()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	pulls, _, err := client.PullRequests.List(ctx, org, repo, &github.PullRequestListOptions{
		State: "open",
		Head:  org + ":" + branch,
	})

	fmt.Println(pulls)
	fmt.Println(domain, org, repo)
}

func main() {
	fmt.Println("start")
	flag.Parse()
	fmt.Println("mode: " + mode)

	switch mode {
	case "current":
		fmt.Println("branch: " + getCurrentBranch())
		getURL(getCurrentBranch())

	case "all":
		fmt.Println("bbb")
	default:
		fmt.Println("ccc")
	}
}
