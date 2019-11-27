package main

import(
	"fmt"
	"docmatcher/pkg"
	"docmatcher/internal"
)

func main() {
	corpus := dictionary.MakeCorpus()
	corpus.AddSample(utils.PascalToWords("HelloWorldThisIsAProgram"), nil)
	match := corpus.FindMatch(utils.PascalToWords("HelloWorld"))
	fmt.Println(match)
}