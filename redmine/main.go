package main

import (
	"log"

	"github.com/jason0x43/go-redmine"
)

func main() {
	session, _ := redmine.NewSession(
		"https://projects.sitepen.com",
		"15c6ed92e21da85a8e2e8ddeb23c5256f1160c1c")

	issues, err := session.GetIssues(nil)
	if err != nil {
		log.Printf("Error: %s", err)
	} else {
		log.Printf("Issues: %s", issues)
	}
}
