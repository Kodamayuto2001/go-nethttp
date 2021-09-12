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

func main() {
	mux := http.NewServeMux()
    mux.HandleFunc("/", middleware1(middleware2(middleware3(indexHandler))))
    mux.HandleFunc("/about", middleware1(middleware2(middleware3(aboutHandler))))

    http.ListenAndServe(":3000", mux)
}