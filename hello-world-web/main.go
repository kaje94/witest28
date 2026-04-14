package main

import (
	"log"
	"net/http"
)

const homepageHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Hello World</title>
  <style>
    *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
    body {
      font-family: system-ui, sans-serif;
      background-color: #ffffff;
      color: #1a1a1a;
      min-height: 100vh;
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 1rem;
    }
    main {
      text-align: center;
    }
    h1 {
      font-size: clamp(2rem, 8vw, 4rem);
      font-weight: 700;
      letter-spacing: -0.02em;
    }
  </style>
</head>
<body>
  <main>
    <h1>Hello, World!</h1>
  </main>
</body>
</html>`

const notFoundHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Page Not Found – Hello World</title>
  <style>
    *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
    body {
      font-family: system-ui, sans-serif;
      background-color: #ffffff;
      color: #1a1a1a;
      min-height: 100vh;
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 1rem;
    }
    main { text-align: center; }
    h1 { font-size: 2rem; font-weight: 700; margin-bottom: 1rem; }
    p { margin-bottom: 1.5rem; }
    a { color: #0057b8; text-decoration: underline; }
    a:hover { color: #003d82; }
  </style>
</head>
<body>
  <main>
    <h1>404 – Page Not Found</h1>
    <p>The page you are looking for does not exist.</p>
    <a href="/">Go back to the homepage</a>
  </main>
</body>
</html>`

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'unsafe-inline'")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte(notFoundHTML)); err != nil {
			log.Println("error writing 404 response:", err)
		}
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := w.Write([]byte(homepageHTML)); err != nil {
		log.Println("error writing homepage response:", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"status":"ok"}`)); err != nil {
		log.Println("error writing health response:", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/", homeHandler)

	handler := securityHeaders(mux)

	log.Println("server listening on :9090")
	if err := http.ListenAndServe("0.0.0.0:9090", handler); err != nil {
		log.Fatal("server error:", err)
	}
}
