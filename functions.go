package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func CountCharacters(filePath string) string {
	stringFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string(stringFile): %v\n", string(stringFile))

	fileContent := string(stringFile)

	contentSplit := strings.Split(fileContent, "")

	fmt.Printf("len(contentSplit): %v\n", len(contentSplit))

	return fileContent

}

// var contentSplitWithoutSpaces []string

// for _, v := range contentSplit {
// 	if v != " " {
// 		contentSplitWithoutSpaces = append(contentSplitWithoutSpaces, v)
// 	}
// }

// fmt.Printf("contentSplitWithoutSpaces: %v\n", contentSplitWithoutSpaces)
// fmt.Printf("len(contentSplitWithoutSpaces): %v\n", len(contentSplitWithoutSpaces))
