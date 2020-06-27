package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	// "github.com/marcelfw/mgit/repository"
	// "gopkg.in/src-d/go-git.v4/internal/revision"
	// "github.com/go-git/go-git"
)

var (
	mode string
)

func init() {
	flag.StringVar(&mode, "mode", "current", "which to open PR current branch or all branch in local")
}

func getCurrentBranch() string {
	dir, err := os.Getwd()
	if err != nil  {
		log.Fatal(err)
	}
	return dir
}

func main() {
	fmt.Println("start")
	flag.Parse()
	fmt.Println(mode)

	switch mode {
	case "current":
		fmt.Println(getCurrentBranch())
	case "all":
		fmt.Println("bbb")
	default:
		fmt.Println("ccc")
	}
}
