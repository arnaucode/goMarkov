package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readTxt(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		//Do something
	}
	dataClean := strings.Replace(string(data), "\n", " ", -1)
	content := string(dataClean)
	return content, err
}

func main() {
	fmt.Print("entry the first word: ")
	newcommand := bufio.NewReader(os.Stdin)
	firstWord, _ := newcommand.ReadString('\n')
	firstWord = strings.TrimSpace(firstWord)
	fmt.Print("first word: ")
	fmt.Println(firstWord)

	fmt.Println("how many words you want on the text?")
	newcommand = bufio.NewReader(os.Stdin)
	answer, _ := newcommand.ReadString('\n')
	answer = strings.TrimSpace(answer)
	fmt.Print("Number of words on text to generate: ")
	fmt.Println(answer)
	count, err := strconv.Atoi(answer)
	if err != nil {
		fmt.Println("incorrect entry, need a positive number")
	}

	text, _ := readTxt("text.txt")

	fmt.Println("generating markov chains")
	states := markov.train(text)
	fmt.Println("generating text")
	generatedText := markov.generateText(states, firstWord, count)
	fmt.Println("")
	fmt.Println(generatedText)
}
