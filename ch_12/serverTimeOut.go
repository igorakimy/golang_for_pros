package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler3(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}
	fmt.Printf("Listening on http://0.0.0.0%s\n", PORT)

	m := http.NewServeMux()
	srv := &http.Server{
		Addr:    PORT,
		Handler: m,
		// Максимальная продолжительность операции чтения запроса, включая его тело.
		ReadTimeout: 3 * time.Second,
		// Максимально допустимое время ожидания ответа.
		WriteTimeout: 3 * time.Second,
	}

	m.HandleFunc("/time", timeHandler3)
	m.HandleFunc("/", myHandler3)

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
