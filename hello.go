package main

import (
	"fmt"

	"rsc.io/quote"

	"github.com/dev001hajipro/greetings"
)

func main() {
	fmt.Println("こんにちは世界!")

	fmt.Println(quote.Go())

	message := greetings.Hello("john")
	fmt.Println(message)
}
