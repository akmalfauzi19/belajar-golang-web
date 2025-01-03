package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	println("Before")
	middleware.Handler.ServeHTTP(writer, request)
	println("After")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (ErrorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error: %s", err)
		}
	}()

	ErrorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Exec")
		fmt.Fprint(writer, "Hello Middleware")
	})

	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Foo Exec")
		fmt.Fprintf(writer, "Hello FOO")
	})

	LogMiddleware := &LogMiddleware{Handler: mux}

	ErrorHandler := &ErrorHandler{Handler: LogMiddleware}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: ErrorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
