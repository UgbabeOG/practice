// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"

// 	"golang.org/x/text/cases"
// 	"golang.org/x/text/language"
// )

// func main() {
// 	if len(os.Args) < 3 {
// 		fmt.Println("usage: go run main.go <input file> <output file>")
// 		os.Exit(1)
// 	}
// 	inputFile := os.Args[1]
// 	outputFile := os.Args[2]

// 	data, err := os.ReadFile(inputFile)
// 	if err != nil {
// 		fmt.Printf("Error reading file: %v\n", err)
// 		os.Exit(1)
// 	}

// 	content := string(data)

// 	words := strings.Fields(content)

// 	for i, word := range words {
// 		if word[0] == '(' && word[len(word)-1] == ')' {
// 			word = word[1 : len(word)-1]
// 		}
// 		if word == strings.ToLower("up") && i > 0 {
// 			words[i-1] = strings.ToUpper(words[i-1])
// 			words = append(words[:i], words[i+1:]...)
// 		}
// 		if word == strings.ToLower("down") && i > 0 {
// 			words[i-1] = strings.ToLower(words[i-1])
// 			words = append(words[:i], words[i+1:]...)
// 		}
// 		if word == strings.ToLower("cap") && i > 0 {
// 			words[i-1] = cases.Title(language.Und).String(words[i-1])
// 			words = append(words[:i], words[i+1:]...)
// 		}
// 		if word == strings.ToLower("hex") && i > 0 {
// 			decimalVal, err := strconv.ParseInt(words[i-1], 16, 64)
// 			if err == nil {
// 				words = append(words[:i], words[i+1:]...)
// 				words[i-1] = fmt.Sprintf("%d", decimalVal)
// 			}
// 		}
// 		if word == strings.ToLower("bin") && i > 0 {
// 			decimalVal, err := strconv.ParseInt(words[i-1], 2, 64)
// 			if err == nil {
// 				words = append(words[:i], words[i+1:]...)
// 				words[i-1] = fmt.Sprintf("%d", decimalVal)
// 			}
// 		}
// 	}

// 	result := strings.Join(words, " ")

// 	err = os.WriteFile(outputFile, []byte(result), 0644)
// 	if err != nil {
// 		fmt.Printf("Error writing file: %v\n", err)
// 		os.Exit(1)
// 	}
// }

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: go run main.go <input file> <output file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("unable to read file")
		os.Exit(1)
	}

	// Use Fields to handle multiple spaces/newlines automatically
	content := strings.Fields(string(data))

	for i := 0; i < len(content); i++ {
		word := content[i]
		cleanWord := strings.Trim(word, "()")

		// Single word transformations
		switch cleanWord {
		case "up":
			if i > 0 {
				content[i-1] = strings.ToUpper(content[i-1])
				content = append(content[:i], content[i+1:]...)
				i-- // Adjust index because we removed an element
			}
		case "low":
			if i > 0 {
				content[i-1] = strings.ToLower(content[i-1])
				content = append(content[:i], content[i+1:]...)
				i--
			}
		case "cap":
			if i > 0 {
				content[i-1] = capitalize(content[i-1])
				content = append(content[:i], content[i+1:]...)
				i--
			}
		case "hex":
			if i > 0 {
				val, err := strconv.ParseInt(content[i-1], 16, 64)
				if err == nil {
					content[i-1] = fmt.Sprintf("%d", val)
					content = append(content[:i], content[i+1:]...)
					i--
				}
			}
		case "bin":
			if i > 0 {
				val, err := strconv.ParseInt(content[i-1], 2, 64)
				if err == nil {
					content[i-1] = fmt.Sprintf("%d", val)
					content = append(content[:i], content[i+1:]...)
					i--
				}
			}
		}

		// Handle tags with arguments: (up, n), (low, n), (cap, n)
		if strings.Contains(cleanWord, ",") {
			parts := strings.Split(cleanWord, ",")
			if len(parts) == 2 {
				tag := parts[0]
				// Clean the number part of any remaining parentheses
				numStr := strings.TrimRight(parts[1], ")")
				n, err := strconv.Atoi(numStr)

				if err == nil && i > 0 {
					start := i - n
					if start < 0 {
						start = 0
					}

					for j := start; j < i; j++ {
						switch tag {
						case "up":
							content[j] = strings.ToUpper(content[j])
						case "low":
							content[j] = strings.ToLower(content[j])
						case "cap":
							content[j] = capitalize(content[j])
						}
					}
					// Remove the tag itself
					content = append(content[:i], content[i+1:]...)
					i--
				}
			}
		}
	}
	var vowels = "aei	ouhAEIOUH"
	for k := 0; k < len(content); k++ {
		word := content[k]
		nextWord := content[k+1]
		if strings.ContainsAny(string(nextWord[0]), vowels) {
			if word == "a" {
				content[k] = "an"
			} else if word == "A" {
				content[k] = "An"
			}

		}

	}

	result := strings.Join(content, " ")
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("unable to write file")
		os.Exit(1)
	}
}

// Helper to handle capitalization consistently
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	lowered := strings.ToLower(s)
	return strings.ToUpper(string(lowered[0])) + lowered[1:]
}
