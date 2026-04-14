package main

import (
	"fmt"
	"log"
	"net/http"
)

const helloPage = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Hello World</title>
<style>
body {
  background-color: #ffffff;
  color: #111111;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem;
}
</style>
</head>
<body>
<h1>Hello, World!</h1>
</body>
</html>`

const notFoundPage = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>404 – Page Not Found</title>
<style>
body {
  background-color: #ffffff;
  color: #111111;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem;
}
</style>
</head>
<body>
<h1>404 – Page Not Found</h1>
</body>
</html>`

// stripServerHeader wraps an http.Handler to remove the Server header from responses.
type stripServerHeader struct {
	handler http.Handler
}

type serverHeaderWriter struct {
	http.ResponseWriter
	wroteHeader bool
}

func (w *serverHeaderWriter) WriteHeader(code int) {
	w.Header().Del("Server")
	w.wroteHeader = true
	w.ResponseWriter.WriteHeader(code)
}

func (w *serverHeaderWriter) Write(b []byte) (int, error) {
	if !w.wroteHeader {
		w.Header().Del("Server")
		w.wroteHeader = true
	}
	return w.ResponseWriter.Write(b)
}

func (s *stripServerHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(&serverHeaderWriter{ResponseWriter: w}, r)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handleNotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, helloPage)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"ok"}`)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, notFoundPage)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/health", handleHealth)

	server := &http.Server{
		Addr:    ":9090",
		Handler: &stripServerHeader{handler: mux},
	}

	log.Println("Starting server on :9090")
	log.Fatal(server.ListenAndServe())
}
