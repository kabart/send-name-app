package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", `{"status":"OK"}`)
}

func sendName(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.ServeFile(w, r, "templates/name.gohtml")
	} else {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("name")
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Ilość znaków w imieniu %s: %d\n", name, countChar(name))
	}
}

func main() {
	const port = 8080
	log.Printf("Starting send name application on port %d", port)

	http.HandleFunc("/health-check", healthCheck)
	http.HandleFunc("/send-name", sendName)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func countChar(name string) int {
	b := []rune(strings.ReplaceAll(name, " ", ""))
	return len(b)
}
