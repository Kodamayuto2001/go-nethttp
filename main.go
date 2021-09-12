package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, middleware!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is middleware test!")
}

func middleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] middleware1")
		next.ServeHTTP(w, r)
		fmt.Println("[END] middleware1")
	}
}

func middleware2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] middleware2")
		next.ServeHTTP(w, r)
		fmt.Println("[END] middleware2")
	}
}

func middleware3(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] middleware3")
		next.ServeHTTP(w, r)
		fmt.Println("[END] middleware3")
	}
}

type middleware func(http.HandlerFunc) http.HandlerFunc

type mwStack struct {
	middlewares []middleware
}

func newMws(mws ...middleware) mwStack {
	return mwStack{append([]middleware(nil), mws...)}
}

func (m mwStack) then(h http.HandlerFunc) http.HandlerFunc {
	for i := range m.middlewares {
		h = m.middlewares[len(m.middlewares)-1-i](h)
	}
	return h 
}

func main() {
	middlewares := newMws(middleware1, middleware2, middleware3)

	mux := http.NewServeMux()
	mux.HandleFunc("/", middlewares.then(indexHandler))
	mux.HandleFunc("/about", middlewares.then(aboutHandler))

	http.ListenAndServe(":3000", mux)
}