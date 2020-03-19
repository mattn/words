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
	for scanner.Scan() {
		for _, token := range t.Tokenize(scanner.Text()) {
			if token.Class != tokenizer.DUMMY {
				fmt.Println(token.Surface)
			}
		}
	}
}
