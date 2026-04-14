# Overview

The **witest28-2** project is a simple Hello World web application intended to serve as a foundational starting point or proof-of-concept for a web-based system. The application will display a "Hello, World!" message to users who visit it via a web browser, confirming that the system is correctly set up and accessible.

The target users are developers, stakeholders, or evaluators who need to verify that a working web application can be deployed and reached over the internet. The application follows a straightforward request-response model: a user visits a URL and receives a rendered page containing the greeting message.

Despite its simplicity, the application will meet a baseline set of quality and reliability standards to ensure it is production-ready in terms of availability, security, and performance — establishing good practices for any future expansion of the project.

---

# Capabilities

## Core Functionality

- The application must display the text "Hello, World!" prominently on the main page when a user navigates to the root URL (`/`).
- The page must render correctly in all modern web browsers (Chrome, Firefox, Safari, Edge — latest two major versions each).
- The application must return an HTTP `200 OK` status code for successful requests to the root URL.
- The application must return an appropriate HTTP `404 Not Found` response for any unrecognized routes or URLs.

## User Interface

- The page must have a valid HTML structure including a `<head>` with a descriptive `<title>` tag and a `<body>` containing the greeting.
- The page title displayed in the browser tab must reflect the application name or greeting (e.g., "Hello World").
- The page must be readable and usable on both desktop and mobile screen sizes (responsive layout).
- The page must meet basic accessibility standards: sufficient color contrast and semantic HTML elements.

## Availability & Performance

- The application must be accessible via a public URL (domain or IP address).
- The application must respond to requests within **500 milliseconds** under normal load conditions.
- The application must maintain **99% uptime** measured on a monthly basis.

## Security

- The application must be served over **HTTPS**, with HTTP requests automatically redirected to HTTPS.
- The application must include appropriate HTTP security headers (e.g., `X-Content-Type-Options`, `X-Frame-Options`, `Content-Security-Policy`).
- The application must not expose server version information or internal error details in responses.

## Error Handling

- The application must display a user-friendly error page for `404 Not Found` and `500 Internal Server Error` responses.
- Error pages must not expose stack traces, file paths, or any internal system information to the end user.
- All unhandled server errors must be logged internally for diagnostic purposes.

## Observability & Health

- The application must expose a `/health` endpoint that returns an HTTP `200 OK` response with a simple status indicator (e.g., `{ "status": "ok" }`).
- Application errors and access events must be logged with timestamps and request metadata (method, path, status code).
