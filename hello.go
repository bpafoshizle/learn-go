package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Printf("quote.Go(): %s\n", quote.Go())
	fmt.Printf("quote.Glass(): %s\n", quote.Glass())
	fmt.Printf("quote.Hello(): %s\n", quote.Hello())
	fmt.Printf("quote.Opt(): %s\n", quote.Opt())
}
