package main

import (
	"fmt"
	"github.com/bmizerany/pat"
	"net/http"
)

func main() {
	m := pat.New()

	m.Post("/action", middleware(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("POST Action!")
		w.Write([]byte("POST ACTION!"))
	})))

	m.Get("/action", moreMiddleware(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("GET Action!")
		w.Write([]byte("GET ACTION!"))
	})))

	fmt.Println("Starting server on port 3000")
	http.ListenAndServe(":3000", m)
}

func moreMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("More Middleware!")
		next.ServeHTTP(w, req)
	}
}

func middleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Middleware!")
		next.ServeHTTP(w, req)
	}
}
