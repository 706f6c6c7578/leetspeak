package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var leetMap = map[rune]rune{
	'A': '4', 'B': '8', 'C': 'C', 'D': 'D', 'E': '3', 'F': 'F', 'G': '6', 'H': 'H',
	'I': '1', 'J': 'J', 'K': 'K', 'L': '1', 'M': 'M', 'N': 'N', 'O': '0', 'P': 'P',
	'Q': '9', 'R': 'R', 'S': '5', 'T': '7', 'U': 'U', 'V': 'V', 'W': 'W', 'X': 'X',
	'Y': 'Y', 'Z': '2',
	'a': '4', 'b': '8', 'c': 'c', 'd': 'd', 'e': '3', 'f': 'f', 'g': '6', 'h': 'h',
	'i': '1', 'j': 'j', 'k': 'k', 'l': '1', 'm': 'm', 'n': 'n', 'o': '0', 'p': 'p',
	'q': '9', 'r': 'r', 's': '5', 't': '7', 'u': 'u', 'v': 'v', 'w': 'w', 'x': 'x',
	'y': 'y', 'z': '2',
}

var reverseLeetMap map[rune]rune

func init() {
	reverseLeetMap = make(map[rune]rune)
	for k, v := range leetMap {
		if v != rune(k) {
			reverseLeetMap[v] = k
		}
	}
}

func encode(input string) string {
	var result strings.Builder
	for _, char := range input {
		if leet, ok := leetMap[char]; ok {
			result.WriteRune(leet)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func decode(input string) string {
	var result strings.Builder
	for _, char := range input {
		if normal, ok := reverseLeetMap[char]; ok {
			result.WriteRune(normal)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func main() {
	decodeFlag := flag.Bool("d", false, "Decode mode")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var result string
		if *decodeFlag {
			result = decode(line)
		} else {
			result = encode(line)
		}
		fmt.Println(result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
