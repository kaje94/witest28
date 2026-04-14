# Overview

The **witest28-6** project is a simple "Hello World" web application intended to serve as a foundational starting point or proof-of-concept for a web-based system. The application delivers a minimal, functional web page that displays a greeting message to any visitor who accesses it via a browser.

The primary target users are developers or stakeholders who need a working baseline web application — for example, to validate a deployment pipeline, test infrastructure, or establish a project scaffold. The system is intentionally lightweight, prioritizing simplicity, accessibility, and reliability over feature richness.

The application will be accessible over standard web protocols (HTTP/HTTPS), require no user authentication, and present a clear, readable greeting on a single web page. It must be stable, fast-loading, and work across all modern browsers.

---

# Capabilities

## Core Display

- The application must serve at least one web page accessible via a standard browser URL.
- The page must display a visible "Hello, World!" greeting message (or equivalent welcoming text) as its primary content.
- The greeting text must be human-readable and clearly distinguishable on the page (e.g., not hidden, invisible, or excessively small).
- The page must have a valid, descriptive HTML `<title>` element (e.g., "Hello World").

## Navigation & Routing

- The application must respond to requests at the root path (`/`).
- Any request to an undefined route must return a clear, user-friendly "Page Not Found" response (HTTP 404).
- The application must return the correct HTTP status code (`200 OK`) for successful page loads.

## Compatibility & Accessibility

- The page must render correctly in all major modern browsers (Chrome, Firefox, Safari, Edge — latest two versions each).
- The page must be usable on both desktop and mobile screen sizes without horizontal scrolling or broken layout.
- The page must meet WCAG 2.1 Level AA accessibility standards, including sufficient color contrast and meaningful text structure.
- The page must use valid, well-formed HTML.

## Performance

- The page must fully load within 2 seconds on a standard broadband connection (≥10 Mbps).
- The total page weight (HTML, CSS, assets) must not exceed 500 KB.
- The application must handle at least 100 concurrent users without degraded response times.

## Availability & Reliability

- The application must target 99.9% uptime (no more than ~8.7 hours downtime per year).
- The application must respond to health-check requests at a `/health` endpoint with an HTTP `200 OK` status and a simple status indicator (e.g., `{"status": "ok"}`).
- The application must recover automatically from crashes or restarts without manual intervention.

## Security

- The application must be served over HTTPS in any production environment.
- The application must not expose any server version, framework, or internal path information in HTTP response headers or page content.
- The application must include basic HTTP security headers (e.g., `Content-Security-Policy`, `X-Frame-Options`, `X-Content-Type-Options`).

## Deployment & Configuration

- The application must be deployable to a standard web hosting environment with no manual file edits required post-deployment.
- The application's listening port must be configurable via an environment variable, defaulting to port `8080`.
- The application must produce basic access logs (timestamp, request path, HTTP status code) to standard output.
