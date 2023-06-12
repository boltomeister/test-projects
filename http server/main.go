package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

const keyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}

	fmt.Printf("%s: got / request. first(%t)=%s, second(%t)=%s, body:\n%s\n", ctx.Value(keyServerAddr), hasFirst, first, hasSecond, second, body)
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))

	myName := r.PostFormValue("myName")
	if myName == "" {
		w.Header().Set("x-missing-field", "myName")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	io.WriteString(w, fmt.Sprintf("Hello, %s!\n", myName))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	ctx := context.Background()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx := context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server is shut down")
	} else if err != nil {
		fmt.Println("error listening the server")
	}
}

// Use this commands to check if the server is working:
// curl http://localhost:8080/
// curl http://localhost:8080/hello
// curl -X POST -d 'This is a body' 'http://localhost:8080?first=20&second10'
// curl -X POST -F 'myName=Ravshan' 'http://localhost:8080/hello'
