package main

import (
	"fmt"

	"github.com/somersbmatthews/surf"
)

// const x509UnknownAuthorityError = "Get \"https://digitalmint.exchange/\": x509: certificate signed by unknown authority"

func main() {
	bow := surf.NewBrowser()
	// err := bow.Open("https://digitalmint.io") // works fine
	err := bow.Open("https://digitalmint.exchange") // returns x509UnknownAuthorityError
	if err != nil {
		panic(err)
	}

	fmt.Println(bow.Title())
}
