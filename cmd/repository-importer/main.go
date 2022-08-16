package main

import (
	"fmt"
	"log"
	"strings"
	"technology-services-and-platforms-accnz/repository-importer/internal/auth"
	"technology-services-and-platforms-accnz/repository-importer/internal/config"
	"technology-services-and-platforms-accnz/repository-importer/internal/github"
)

func main() {
	// Create configuration from user input and initialise client
	conf, err := config.NewConfigFromUser()
	if err != nil {
		log.Fatal(err)
	}
	gh := github.NewGithubClient(auth.NewAuthClient(conf.GITHUB_API_TOKEN))

	for _, r := range conf.REPOSITORY_URLS {
		// Retrieve ADO repo name
		s := strings.Split(r, "/")
		name := s[len(s)-1]
		fmt.Printf("attempting to import repository '%s'\n", name)

		// Ensure repo does not already exist in GitHub
		if !gh.RepositoryExists(conf.GITHUB_ORG_NAME, name) {
			// If team slug is provided, create a team repo, else create regular repo
			if conf.GITHUB_TEAM_SLUG != "" {
				teamId, err := gh.GetTeam(conf.GITHUB_ORG_NAME, conf.GITHUB_TEAM_SLUG)
				if err != nil {
					fmt.Printf("%s, moving on...\n", err)
					continue
				}
				_, err = gh.CreateTeamRepository(conf.GITHUB_ORG_NAME, name, "internal", teamId)
				if err != nil {
					fmt.Printf("%s, moving on...\n", err)
					continue
				}
				fmt.Printf("created repository %s in team %s... SUCCESS\n", name, conf.GITHUB_TEAM_SLUG)

				// Import existing contents from Azure DevOps
				_, err = gh.StartImport(conf.GITHUB_ORG_NAME, name, r, conf.AZURE_DEVOPS_USER, conf.AZURE_DEVOPS_PAT)
				if err != nil {
					// Stop if import fails
					log.Fatal(err)
				}
				fmt.Printf("import repository %s in team %s... SUCCESS\n", name, conf.GITHUB_TEAM_SLUG)

			} else {
				// No team slug provided
				_, err := gh.CreateRepository(conf.GITHUB_ORG_NAME, name, "internal")
				if err != nil {
					fmt.Printf("%s, moving on...\n", err)
					continue
				}
				fmt.Printf("created repository %s... SUCCESS\n", name)

				_, err = gh.StartImport(conf.GITHUB_ORG_NAME, name, r, conf.AZURE_DEVOPS_USER, conf.AZURE_DEVOPS_PAT)
				if err != nil {
					// Stop if import fails
					log.Fatal(err)
				}
				fmt.Printf("import repository %s... SUCCESS\n", name)
			}
		} else {
			fmt.Printf("repository '%s' already exists, moving on...\n", name)
			continue
		}
	}

	fmt.Println("Finished importing repositories, do note that the import may take some time to complete in the background")
}
