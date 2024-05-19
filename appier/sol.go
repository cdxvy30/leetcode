/**
* Usage of this file.
*	1. Ensure there is golang environment installed. I use Go 1.20.3.
* 2. Execute command: `go run sol.go "string to be tested."`, e.g. `go run sol.go "Appier is the best technical corporation in the world." `
* 3. Evaluate the results.
 */

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func FindFrequentWord(text string) string {
	required := "appier"
	wordPattern := regexp.MustCompile(`\b\w+\b`) // This should help identify only string and number.
	words := wordPattern.FindAllString(strings.ToLower(text), -1)

	count := make(map[string]int)
	max := 0
	mostFreqWord := ""

	for _, word := range words {
		for _, char := range required {
			if strings.ContainsRune(word, char) {
				count[word]++
				if count[word] > max {
					max = count[word]
					mostFreqWord = word
				}
				break
			}
		}
	}

	return mostFreqWord
}

func main() {
	// Use default text if there is no command line input provided.
	inputText := "Appier is a software-as-a-service (SaaS) company that uses artificial intelligence (AI) to power business decision-making. Founded in 2012 with a vision of democratizing AI, Appier's mission is turning AI into ROI by making software intelligent. Appier now has 17 offices across APAC, Europe and U.S., and is listed on the Tokyo Stock Exchange (Ticker number: 4180)."
	if len(os.Args) > 1 {
		inputText = strings.Join(os.Args[1:], " ")
	} else {
		fmt.Println("No command line input found, using default text.")
	}

	freqWord := FindFrequentWord(inputText)
	fmt.Printf("The most frequent word containing any character from 'appier' is: %s\n", freqWord)
}
