package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ikawaha/kagome/tokenizer"
)

func main() {
	t := tokenizer.New()
	scanner := bufio.NewScanner(os.Stdin)
	words := map[string]struct{}{}
	for scanner.Scan() {
		for _, token := range t.Tokenize(scanner.Text()) {
			if token.Class != tokenizer.DUMMY {
				if _, ok := words[token.Surface]; !ok {
					fmt.Println(token.Surface)
					words[token.Surface] = struct{}{}
				}
			}
		}
	}
}
