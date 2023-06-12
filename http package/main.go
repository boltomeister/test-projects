package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

const keyServerAddr = "serverAdrr"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Println("got / request")
	io.WriteString(w, fmt.Sprintf("addr: %s\n Root page.\n", ctx.Value(keyServerAddr)))
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Println("got /hello request")
	io.WriteString(w, fmt.Sprintf("addr: %s\n Hello World!\n", ctx.Value(keyServerAddr)))
}

func main() {
	router := http.NewServeMux()

	ctx := context.Background()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	router.HandleFunc("/", getRoot)
	router.HandleFunc("/hello", getHello)

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server is closed")
	} else if err != nil {
		fmt.Println("unknown error")
		os.Exit(1)
	}
}
