package config

// Need to figure out a way to test "golang.org/x/term.ReadPassword()" as the syscall is based on architecture and OS, might have to mock our own PasswordReader and inject this into the solution

// import (
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func setupUserInput(in []string) *os.File {
// 	// Hacky way to feed to stdin
// 	tmpfile, err := ioutil.TempFile("", "test")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Write content to tmp file
// 	for _, s := range in {
// 		if _, err := tmpfile.Write([]byte(s)); err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	// Set stdin to tmp file
// 	os.Stdin = tmpfile
// 	if _, err := tmpfile.Seek(0, 0); err != nil {
// 		log.Fatal(err)
// 	}

// 	return tmpfile
// }

// func cleanUpUserInput(oldIn *os.File, inFile *os.File) {
// 	defer func() { os.Stdin = oldIn }() // Set stdin back
// 	defer os.Remove(inFile.Name())      // Remove tmp file
// 	if err := inFile.Close(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// // Test that the required environment variables are set
// func TestSuppliedVars(t *testing.T) {
// 	oldStdin := os.Stdin

// 	str := []string{
// 		"test-github-organisation\n",
// 		"test-team-slug\n",
// 		"\n",
// 		"ado_user\n",
// 		"\n",
// 		"https://organisation@dev.azure.com/organisation/project/_git/repoone\n",
// 		"https://organisation@dev.azure.com/organisation/project/_git/repotwo\n",
// 		"\n",
// 		"yes\n",
// 	}
// 	in := setupUserInput(str)
// 	defer cleanUpUserInput(oldStdin, in)

// 	conf, _ := NewConfigFromUser()

// 	// Assertions
// 	assert := assert.New(t)
// 	assert.Equal("https://api.github.com", conf.GITHUB_BASE_URL, "The values should match")
// 	assert.Equal("test-github-organisation", conf.GITHUB_ORG_NAME, "The values should match")
// 	assert.Equal("test-team-slug", conf.GITHUB_TEAM_SLUG, "The values should match")
// 	assert.Equal("ghp_s3cr3tp4tt0k3n", conf.GITHUB_API_TOKEN, "The values should match")
// 	assert.Equal("ado_user", conf.AZURE_DEVOPS_USER, "The values should match")
// 	assert.Equal("ado_s3cur3t0k3n!", conf.AZURE_DEVOPS_PAT, "The values should match")
// 	assert.Contains(conf.REPOSITORY_URLS, "https://organisation@dev.azure.com/organisation/project/_git/repoone", "The URL should be in the list")
// 	assert.Contains(conf.REPOSITORY_URLS, "https://organisation@dev.azure.com/organisation/project/_git/repotwo", "The URL should be in the list")

// }

// func TestErrorThownUnsuppliedVars(t *testing.T) {
// 	oldStdin := os.Stdin

// 	str := []string{
// 		"\n",
// 	}
// 	in := setupUserInput(str)
// 	defer cleanUpUserInput(oldStdin, in)

// 	_, err := NewConfigFromUser()

// 	// Assertions
// 	assert.EqualError(t, err, "the GitHub Organisation Name is required", "The errors should match")

// }
