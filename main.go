package main

import (
	"fmt"

	"github.com/google/go-github/github"
)

func main() {
	fmt.Println("hello")

	client := github.NewClient(nil)
	orgs, _, err := client.Repositories.List("justyntemme", nil)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(orgs)

}
