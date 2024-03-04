package main

import (
	"fmt"
	"os"

	"github.com/book/help/cmd/templ"
	"github.com/book/help/internal"
	"github.com/leapkit/core/server"
)

func main() {
	s := server.New("Book Help")

	templ.Generate()

	if err := internal.AddRoutes(s); err != nil {
		os.Exit(1)
	}

	if err := s.Start(); err != nil {
		fmt.Println(err)
	}
}
