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

	// Q := "printcombN"

	cmd := exec.Command("curl", "-L", "-H",
		"Accept: application/vnd.github.text-match+json",
		"-H",
		"Authorization: Bearer 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",
		"-H",
		"X-GitHub-Api-Version: 2022-11-28",
		"https://api.github.com/search/code?q="+CountCharacters(FILEPATH)+"+in:file+language:go",
	)

	result, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	var jsonU DataReceived
	err = json.Unmarshal(result, &jsonU)
	if err != nil {
		log.Fatal(err)
	}

	fragment := ""
	ghSearchContent := ""

	for _, v := range jsonU.Items {
		for _, e := range v.TextMatches {

			fragment = e.Fragment
			ghSearchContent = e.ObjectURL

		}
	}

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
