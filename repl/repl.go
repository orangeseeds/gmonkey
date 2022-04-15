package repl

// Read Eval Print Loop

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/orangeseeds/gmonkey/lexer"
	"github.com/orangeseeds/gmonkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf("%v", PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if line == "exit" {
			os.Exit(1)
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
