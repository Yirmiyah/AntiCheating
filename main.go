package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

const Access_Token = "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"

func main() {

	GetGitHub(Access_Token)

}

func GetGitHub(tokenReceive string) {

	// cmd := exec.Command("curl", "-L", "")

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
	query := "printcombn"
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

// package main

// import (
// 	"bytes"
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/exec"
// 	"path/filepath"
// 	"strings"
// 	"time"

// 	"github.com/sashabaranov/go-openai"
// )

// const EXERCISES_NAMES = "all.txt"
// const PATH = "../../../piscine-go/piscine/"
// const NAME_PATH = "all.txt"
// const EXTENSION = ".go"
// const FILE_SOLUTION = "/home/student07/projet_perso/Big_Projects/Cheating/excercise_solution.txt"
// const TEST_FILE = "/test.txt"

// var countResponse int

// func main() {

// 	gpt3Response := "nothing"
// 	cmd0 := exec.Command("xsel", "-b", "-d")
// 	err := cmd0.Run()
// 	if err != nil {
// 		fmt.Println("Erreur exec.Command delete primary buffer!")
// 	}
// 	GetNameFile()
// 	for {
// 		time.Sleep(time.Second)
// 		if gpt3Response != WriteCtrlV() {
// 			gpt3Response = WriteCtrlV()

// 			fmt.Printf("gpt3Response: %v\n", gpt3Response)

// 		} else {
// 			cmd0 := exec.Command("xsel", "-b", "-d")
// 			err := cmd0.Run()
// 			if err != nil {
// 				fmt.Println("Erreur exec.Command delete primary buffer!")
// 			}
// 			time.Sleep(time.Second)
// 			gpt3Response = "nothing"
// 		}
// 	}
// }

// func WriteCtrlV() string {

// 	cmd2 := exec.Command("xsel", "-b", "-o")
// 	var out1 bytes.Buffer
// 	cmd2.Stdout = &out1
// 	err := cmd2.Run()
// 	if err != nil {
// 		fmt.Println("erreur copy content into the primary buffer")

// 	}
// 	clipboard := out1.String()

// 	// fmt.Printf("clipboard content: %v\n", clipboard)

// 	if strings.Contains(clipboard, "-gpt3") {

// 		client := openai.NewClient("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")
// 		resp, err := client.CreateChatCompletion(
// 			context.Background(),
// 			openai.ChatCompletionRequest{
// 				Model: openai.GPT3Dot5Turbo,
// 				Messages: []openai.ChatCompletionMessage{
// 					{
// 						Role:    openai.ChatMessageRoleUser,
// 						Content: clipboard,
// 					},
// 				},
// 			},
// 		)
// 		if err != nil {
// 			fmt.Println("Error GPT3 Response...")
// 			log.Fatal(err)
// 		}

// 		gpt3Response := resp.Choices[0].Message.Content
// 		countResponse++
// 		fmt.Printf("countResponse: %v\n", countResponse)

// 		if gpt3Response != "" {
// 			fmt.Printf("gpt3Response: %v\n", gpt3Response)
// 			gpt3 := strings.NewReader(gpt3Response)
// 			cmd2 = exec.Command("xsel", "-i", "-b")
// 			cmd2.Stdin = gpt3
// 			if err := cmd2.Run(); err != nil {
// 				fmt.Println("erreur cmd2.Run()")
// 			}
// 			return gpt3Response
// 		} else {
// 			return "nothing..."
// 		}
// 	}
// 	return "nothing..."
// }

// func GetNameFile() {
// 	dir := PATH // replace with the path to your folder
// 	files, err := os.ReadDir(dir)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	f, err := os.Create("all.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer f.Close()
// 	for _, file := range files {
// 		if !file.IsDir() {
// 			filename := filepath.Base(file.Name())
// 			name := strings.Split(filename, ".")
// 			fmt.Fprintln(f, name[0])
// 		}
// 	}
// 	fmt.Println("File names written to all.txt")
// }
