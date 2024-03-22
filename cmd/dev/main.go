package main

import (
	"fmt"
	"github.com/book-quote/internal"

	"github.com/leapkit/core/gloves"
)

func main() {
	err := gloves.Start(
		"cmd/app/main.go",
		internal.GlovesOptions...,
	)

	if err != nil {
		fmt.Println(err)
	}

}
