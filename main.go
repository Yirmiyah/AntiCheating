package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const Access_Token = "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"

const FILEPATH = "../../../piscine-go/piscine/printcombn.go"

func main() {

	// GetGitHub(Access_Token)

	// SearchTextMatch()

	// SearchAll()

	// CountCharacters(FILEPATH)
	SearchAll()

}

func GetGitHub(tokenReceive string) {

	// bText, err := os.ReadFile("../../../piscine-go/piscine/printcombn.go")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// codeNeeded := string(bText)

	token := os.Getenv(tokenReceive)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	// bText, err := os.ReadFile("../../../piscine-go/piscine/printcombn.go")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Recherche du code sur Github
	query := "print"
	opts := &github.SearchOptions{
		Sort: "indexed",
	}
	result, _, err := client.Search.Code(ctx, query, opts)
	if err != nil {
		fmt.Printf("Erreur lors de la recherche du code: %v\n", err)
		return
	}

	// Vérification de la correspondance
	if result.GetTotal() > 0 {
		fmt.Printf("Le code existe sur Github.\n")
	} else {
		fmt.Printf("Le code n'existe pas sur Github.\n")
	}

}

func SearchAll() {

	// bText, err := os.ReadFile("../../../piscine-go/piscine/printcombn.go")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// codeNeeded := string(bText)

	Q := "Itoabase"

	cmd := exec.Command("curl", "-L", "-H",
		"Accept: application/vnd.github.text-match+json",
		"-H",
		"Authorization: Bearer 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",
		"-H",
		"X-GitHub-Api-Version: 2022-11-28",
		"https://api.github.com/search/code?q="+Q+"+in:file+language:golang",
	)

	result, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("string(result): %v\n", string(result))

	var jsonU DataReceived
	err = json.Unmarshal(result, &jsonU)
	if err != nil {
		log.Fatal(err)
	}

	var fragment []string
	var longest string
	ghSearchContent := ""

	for _, v := range jsonU.Items {
		for i, e := range v.TextMatches {

			fragment = append(fragment, v.TextMatches[i].Fragment)
			longest = v.TextMatches[0].Fragment

			ghSearchContent = e.ObjectURL

		}
	}

	for i := 0; i < len(fragment); i++ {

		if len(fragment[i]) >= len(longest) {
			longest = fragment[i]
		}
	}

	fmt.Printf("longest: %v\n", longest)
	fmt.Printf("fragment: %v\n", fragment)
	responseFromGH, err := http.Get(ghSearchContent)
	if err != nil {
		log.Fatal(err)
	}

	searchContent, err := io.ReadAll(responseFromGH.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string(searchContent): %v\n", string(searchContent))

	responseFromGH.Body.Close()
}

func SearchTextMatch() {

	cmd2 := exec.Command("curl", "-L",
		"Accept: application/vnd.github.text-match+json",
		"https://api.github.com/search/code?q=printcombn+in:file+language:go+repo:printcomb+language:golang+state:open&sort=created&order=asc",
	)
	result2, err := cmd2.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string(result2): %v\n", string(result2))

}
