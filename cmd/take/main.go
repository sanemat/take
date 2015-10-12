package main

import (
	"flag"
	"fmt"
	"github.com/jbai/take"
	"math"
	"os"
)

func invalidArguments(message string) {
	fmt.Fprintf(os.Stderr, "Invalid arguments: %v\n", message)
	os.Exit(1)
}

func main() {
	fromPtr := flag.Int("f", 0, "From column")
	toPtr := flag.Int("t", math.MinInt16, "To column")
	delimiterPtr := flag.String("d", " ", "Delimiter used to split each line into columns.")
	flag.Parse()

	if *toPtr == math.MinInt16 {
		*toPtr = *fromPtr
	}
	take.Process(*fromPtr, *toPtr, *delimiterPtr)

}
