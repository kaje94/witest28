# Overview

The **witest28-7** project is a simple "Hello World" web application intended to serve as a foundational starting point for web development. The system will display a greeting message to users who visit the application via a web browser, confirming that the application is live and functioning correctly.

The target users are developers and stakeholders who need to verify that a basic web environment is correctly set up and operational. The application is intentionally minimal, focusing on delivering a clear, accessible, and reliable user-facing page with no unnecessary complexity.

The application will be publicly accessible via a standard web browser, load quickly, and require no user authentication or data input. It serves as a baseline from which more complex features can be built.

---

# Capabilities

## Core Display

- The application must render a webpage that displays the text "Hello, World!" prominently on the page.
- The greeting message must be visible without any scrolling on standard desktop and mobile screen sizes.
- The page must have a descriptive and meaningful `<title>` tag (e.g., "Hello World") visible in the browser tab.
- The page must include a clear visual hierarchy with the greeting as the primary heading element.

## Accessibility

- The page must meet WCAG 2.1 Level AA accessibility standards.
- All text must have a minimum contrast ratio of 4.5:1 against its background.
- The page must be navigable and readable by screen readers.
- The page must use semantic HTML markup (e.g., proper heading tags).

## Compatibility

- The application must render correctly in the latest two major versions of Chrome, Firefox, Safari, and Edge.
- The application must be fully functional and readable on mobile devices with screen widths as small as 320px.
- The application must not rely on browser plugins or extensions to function.

## Performance

- The page must fully load within 2 seconds on a standard broadband connection (≥10 Mbps).
- The total page weight (HTML, CSS, assets) must not exceed 500 KB.
- The application must achieve a Google Lighthouse performance score of 90 or above.

## Availability & Reliability

- The application must be available at a stable, publicly accessible URL.
- The application must return an HTTP 200 status code for requests to the root URL (`/`).
- The application must have an uptime of at least 99% measured on a monthly basis.

## Security

- The application must be served over HTTPS.
- The application must include appropriate HTTP security headers (e.g., `Content-Security-Policy`, `X-Frame-Options`, `X-Content-Type-Options`).
- The application must not expose any server-side error details or stack traces to the end user.

## Error Handling

- If a user navigates to a non-existent route, the application must return a user-friendly 404 page.
- The 404 page must include a link back to the main page.
