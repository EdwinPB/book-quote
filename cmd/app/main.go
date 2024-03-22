package main

import (
	"fmt"
	"github.com/book-quote/internal"
	"github.com/leapkit/core/envor"
	"net/http"
	"os"

	"github.com/leapkit/core/server"
)

func main() {
	s := server.New(
		server.WithHost(envor.Get("HOST", "0.0.0.0")),
		server.WithPort(envor.Get("PORT", "3000")),
	)

	if err := internal.AddRoutes(s); err != nil {
		os.Exit(1)
	}

	fmt.Println("Server started at", s.Addr())
	err := http.ListenAndServe(s.Addr(), s.Handler())
	if err != nil {
		fmt.Println(err)
	}
}
