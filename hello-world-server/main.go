package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const helloHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Hello World</title>
  <style>
    *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
    body {
      font-family: system-ui, sans-serif;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      min-height: 100vh;
      background-color: #ffffff;
      color: #1a1a1a;
      padding: 1rem;
    }
    h1 {
      font-size: clamp(2rem, 8vw, 4rem);
      font-weight: 700;
      text-align: center;
    }
    p {
      margin-top: 1rem;
      font-size: 1.125rem;
      color: #4a4a4a;
      text-align: center;
    }
  </style>
</head>
<body>
  <main>
    <h1>Hello, World!</h1>
    <p>Welcome to witest28-6.</p>
  </main>
</body>
</html>`

const notFoundHTML = `<!DOCTYPE html>
<html lang="en">
<head><meta charset="UTF-8"/><title>Page Not Found</title></head>
<body><h1>404 – Page Not Found</h1><p><a href="/">Go to Home</a></p></body>
</html>`

// responseWriterWrapper captures the status code written by downstream handlers.
type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterWrapper) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// securityHeaders sets security response headers on every response.
func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Server", "")
		next.ServeHTTP(w, r)
	})
}

// accessLog emits structured access logs to stdout.
func accessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrapped, r)
		fmt.Printf("%s %s %s %d\n",
			time.Now().UTC().Format(time.RFC3339),
			r.Method,
			r.URL.Path,
			wrapped.statusCode,
		)
	})
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, notFoundHTML)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, helloHTML)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"ok"}`)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/", rootHandler)

	handler := securityHeaders(accessLog(mux))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
