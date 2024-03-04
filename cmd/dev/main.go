package main

import (
	"fmt"

	"github.com/book/help/cmd/templ"
	"github.com/book/help/internal"
	"github.com/leapkit/core/gloves"
)

func main() {
	templ.Generate()

	err := gloves.Start(
		"cmd/app/main.go",
		internal.GlovesOptions...,
	)

	if err != nil {
		fmt.Println(err)
	}

}
