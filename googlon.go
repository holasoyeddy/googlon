package main

import (
	"fmt"
	"sort"
	"strings"
)

// Result structure groups the results that we want to print out to the console.
type Result struct {
	Prepositions  int64
	Verbs         int64
	SubVerbs      int64
	Vocabulary    []string
	PrettyNumbers int64
}

// Golang maps are implemented as hash tables. However,
// hash sets are not natively implemented in Golang. What we are doing here is
// implementing pseudo-constant hash sets using a struct{} as value. This will
// give us O(1) look up time, on average.
var foo = map[byte]struct{}{
	'u': struct{}{},
	'd': struct{}{},
	'x': struct{}{},
	's': struct{}{},
	'm': struct{}{},
	'p': struct{}{},
	'f': struct{}{},
}

// Another map that allows us to look up Googlon lexicographical order of characters.
var alphabeticalOrder = map[byte]int{
	's': 0,
	'x': 1,
	'o': 2,
	'c': 3,
	'q': 4,
	'n': 5,
	'm': 6,
	'w': 7,
	'p': 8,
	'f': 9,
	'y': 10,
	'h': 11,
	'e': 12,
	'l': 13,
	'j': 14,
	'r': 15,
	'd': 16,
	'g': 17,
	'u': 18,
	'i': 19,
}

// ProcessText validates text conditions and prints out the results to console.
func processText(text string) {

	result := Result{
		Prepositions:  0,
		Verbs:         0,
		SubVerbs:      0,
		PrettyNumbers: 0,
	}

	seen := make(map[string]struct{})

	words := strings.Split(text, " ")

	for _, word := range words {

		if isPreposition(word) {
			result.Prepositions++
		}

		if isVerb(word) {
			result.Verbs++

			if isSubjunctiveVerb(word) {
				result.SubVerbs++
			}
		}

		if _, seenBefore := seen[word]; !seenBefore {
			result.Vocabulary = append(result.Vocabulary, word)
			seen[word] = struct{}{}
		}
	}

	googlonSort(result.Vocabulary)

	for _, distinctWord := range result.Vocabulary {
		if isPrettyNumber(distinctWord) {
			result.PrettyNumbers++
		}
	}

	printResults(result)
}

// Validates if a word is a preposition:
// Time Complexity: O(n) (Due to strings.Contains method)
// Space Complexity: O(1)
func isPreposition(word string) bool {

	// Get the length of the word to avoid calculating again
	length := len(word)

	// Checks last letter in the word against foo letters map
	_, endsWithFoo := foo[word[length-1]]

	// Validate word length is 6, ends in a foo letter, and does not contain the letter u
	return length == 6 && endsWithFoo && !strings.Contains(word, "u")
}

// Validates if a word is a verb:
// Time Complexity: O(1)
// Space Complexity: O(1)
func isVerb(word string) bool {

	// Get the length of the word to avoid calculating again
	length := len(word)

	// Checks last letter in the word against foo letters map
	_, endsWithFoo := foo[word[length-1]]

	// Check that word length is 6 or more and ends with a bar letter
	return length >= 6 && !endsWithFoo
}

// Checks if the verb is subjunctive.
// Time Complexity: O(1)
// Space Complexity: O(1)
func isSubjunctiveVerb(verb string) bool {

	// Checks if the first letter of the word is foo letter
	_, startsWithFoo := foo[verb[0]]

	// Validate verb and validate that the first letter is bar letter
	return isVerb(verb) && !startsWithFoo
}

// Validates if a word is a pretty number.
// Time Complexity: O(N) (due to wordToNumber)
// Space complexity: O(1)
func isPrettyNumber(word string) bool {

	number := wordToNumber(word)
	return number >= 81827 && number%3 == 0
}

func wordToNumber(word string) int64 {

	var number int64
	var base int64

	number, base = 0, 1

	for index := range word {
		number += int64(alphabeticalOrder[word[index]]) * base
		base *= 20
	}

	return number
}

// GooglonSort receives a slice of words and orders them according to the Googlon
// lexicographical order.
// Time Complexity: O(N) (Average Case), but O(N*M) (Worst case where all words are the same)
// Space Complexity: O(1) since it is done in place.
func googlonSort(words []string) {

	// Define our custom iterator function
	var iterator func(i, j int) bool

	iterator = func(i, j int) bool {

		// define our aux variables
		wordA := words[i]
		wordB := words[j]
		letterIndex := 0

		// We're only iterating in case the first two letters of the words are the same.
		// This should be O(1) on average, but can get to O(N) in worse case scenarios.
		for letterIndex < len(wordA) && letterIndex < len(wordB) {

			// Check if first two letters are different.
			if wordA[letterIndex] != wordB[letterIndex] {
				// Return comparison of letters at index.
				return alphabeticalOrder[wordA[letterIndex]] < alphabeticalOrder[wordB[letterIndex]]
			}

			// Increment our letter index if the evaluated letters match and try to evaluate next letters.
			letterIndex++
		}

		// If all letters match, return the shortest of the two.
		return len(wordA) < len(wordB)
	}

	// Use the sort.Slice function with our custom iterator function.
	sort.Slice(words, iterator)
}

func printResults(result Result) {

	fmt.Printf("1.) There are %v prepositions in the text.\n", result.Prepositions)
	fmt.Printf("2.) There are %v verbs in the text.\n", result.Verbs)
	fmt.Printf("3.) There are %v subjunctive verbs in the text.\n", result.SubVerbs)

	fmt.Printf("4.) Vocabulary list: ")

	for _, word := range result.Vocabulary {
		fmt.Printf("%v ", word)
	}

	fmt.Printf("\n5.) There are %v distinct pretty numbers in the text.", result.PrettyNumbers)
}
