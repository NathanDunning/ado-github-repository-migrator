package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/term"
)

type Config struct {
	GITHUB_BASE_URL   string `required:"true"`
	GITHUB_TEAM_SLUG  string
	GITHUB_ORG_NAME   string   `required:"true"`
	GITHUB_API_TOKEN  string   `required:"true"`
	AZURE_DEVOPS_USER string   `required:"true"`
	AZURE_DEVOPS_PAT  string   `required:"true"`
	REPOSITORY_URLS   []string `required:"true"`
}

var (
	GithubBaseUrl string = "https://api.github.com"

// // Implement when working on pipeline version
// requiredVars         = []string{
// 	"GITHUB_BASE_URL",
// 	"GITHUB_ORG_NAME",
// 	"GITHUB_API_TOKEN",
// 	"AZURE_DEVOPS_USER",
// 	"AZURE_DEVOPS_PAT",
// 	"REPOSITORY_URLS",
// }
)

func (c Config) print() {
	fmt.Println("GITHUB_TEAM_SLUG:", c.GITHUB_TEAM_SLUG)
	fmt.Println("GITHUB_ORG_NAME:", c.GITHUB_ORG_NAME)
	fmt.Println("GITHUB_API_TOKEN: ", "masked")
	fmt.Println("AZURE_DEVOPS_USER:", c.AZURE_DEVOPS_USER)
	fmt.Println("AZURE_DEVOPS_PAT:", "masked")
	fmt.Println("REPOSITORY_URLS:\n", c.REPOSITORY_URLS)
}

// // Implement when working on pipeline version
// func checkRequiredEnvVars() {
// 	// Check env vars are set
// 	err := []string{}
// 	for _, v := range requiredVars {
// 		if os.Getenv(v) == "" {
// 			err = append(err, v)
// 		}
// 	}
// 	if len(err) > 0 {
// 		fmt.Println("Missing required environment variables:")
// 		for _, e := range err {
// 			fmt.Println(e)
// 		}
// 		os.Exit(1)
// 	}
// }

// // Implement when working on pipeline version
// func NewConfigFromEnv() (*Config, error) {
// 	checkRequiredEnvVars()

// 	return &Config{
// 		GITHUB_BASE_URL:  GithubBaseUrl,
// 		GITHUB_API_TOKEN: os.Getenv("GITHUB_API_TOKEN"),
// 		AZURE_DEVOPS_PAT: os.Getenv("AZURE_DEVOPS_PAT"),
// 	}, nil
// }

func NewConfigFromUser() (*Config, error) {
	// GitHub Org Name
	fmt.Println("Enter your GitHub Organisation Slug/ID:")
	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	ghorg := strings.TrimSuffix(in, "\n")
	if ghorg == "" {
		return nil, fmt.Errorf("the GitHub Organisation Slug/ID is required")
	}

	// GitHub Team Slug
	fmt.Println("Enter your GitHub Team Slug:")
	in, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	ghts := strings.TrimSuffix(in, "\n")

	// GitHub PAT Token
	fmt.Println("Enter your GitHub Personal Access Token:")
	passwd, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}
	ghp := strings.TrimSuffix(string(passwd), "\n")
	if ghp == "" {
		return nil, fmt.Errorf("the GitHub Personal Access Token is required")
	}

	// Azure DevOps Username
	fmt.Println("Enter your Azure DevOps Username:")
	in, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	adou := strings.TrimSuffix(in, "\n")
	if adou == "" {
		return nil, fmt.Errorf("the Azure DevOps Username is required")
	}

	// Azure DevOps PAT Token
	fmt.Println("Enter your Azure DevOps Token:")
	passwd, err = term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}
	adop := strings.TrimSuffix(string(passwd), "\n")
	if adop == "" {
		return nil, fmt.Errorf("the Azure DevOps Token is required")
	}

	// Azure DevOps Repository URLs
	fmt.Println("Paste your Azure DevOps Clone URL, separate with new line:")
	var adorepos []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		adorepos = append(adorepos, strings.TrimSuffix(line, "\n"))
	}
	if len(adorepos) == 0 {
		return nil, fmt.Errorf("the Azure DevOps Clone URL is required")
	}

	c := &Config{
		GITHUB_BASE_URL:   GithubBaseUrl,
		GITHUB_ORG_NAME:   ghorg,
		GITHUB_TEAM_SLUG:  ghts,
		GITHUB_API_TOKEN:  ghp,
		AZURE_DEVOPS_USER: adou,
		AZURE_DEVOPS_PAT:  adop,
		REPOSITORY_URLS:   adorepos,
	}

	c.print()

	// Confirm
	fmt.Println("Is the above correct? Type 'yes' to proceed:")
	in, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	resp := strings.TrimSuffix(in, "\n")
	if resp != "yes" {
		os.Exit(1)
	}

	return c, nil
}
