package main

import (
	"fmt"
	"log"
	"net/http"
)

const helloHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Hello World</title>
  <style>
    *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
    body {
      display: flex;
      align-items: center;
      justify-content: center;
      min-height: 100vh;
      background-color: #ffffff;
      font-family: system-ui, sans-serif;
      overflow-x: hidden;
    }
    h1 {
      font-size: clamp(2rem, 8vw, 5rem);
      color: #1a1a1a;
      text-align: center;
      padding: 1rem;
    }
  </style>
</head>
<body>
  <h1>Hello World</h1>
</body>
</html>`

const notFoundHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>404 Not Found</title>
  <style>
    *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
    body {
      display: flex;
      align-items: center;
      justify-content: center;
      min-height: 100vh;
      background-color: #ffffff;
      font-family: system-ui, sans-serif;
      overflow-x: hidden;
    }
    h1 {
      font-size: clamp(1.5rem, 5vw, 3rem);
      color: #1a1a1a;
      text-align: center;
      padding: 1rem;
    }
  </style>
</head>
<body>
  <h1>404 Not Found</h1>
</body>
</html>`

type secureWriter struct {
	http.ResponseWriter
}

func (sw secureWriter) WriteHeader(code int) {
	sw.Header().Del("Server")
	sw.ResponseWriter.WriteHeader(code)
}

func (sw secureWriter) Write(b []byte) (int, error) {
	sw.Header().Del("Server")
	return sw.ResponseWriter.Write(b)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sw := secureWriter{w}

		defer func() {
			if err := recover(); err != nil {
				sw.Header().Set("Content-Type", "text/plain; charset=UTF-8")
				sw.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(sw, "Internal Server Error")
			}
		}()

		switch r.URL.Path {
		case "/":
			sw.Header().Set("Content-Type", "text/html; charset=UTF-8")
			sw.WriteHeader(http.StatusOK)
			fmt.Fprint(sw, helloHTML)
		case "/health":
			sw.Header().Set("Content-Type", "application/json")
			sw.WriteHeader(http.StatusOK)
			fmt.Fprint(sw, `{"status":"ok"}`)
		default:
			sw.Header().Set("Content-Type", "text/html; charset=UTF-8")
			sw.WriteHeader(http.StatusNotFound)
			fmt.Fprint(sw, notFoundHTML)
		}
	})

	log.Println("Starting server on :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
