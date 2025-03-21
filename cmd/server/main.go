package main

import (
	"fmt"
	"log"
	"net/http"
	"studybud/internal/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)

	port := ":8080"
	fmt.Println("server is running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
