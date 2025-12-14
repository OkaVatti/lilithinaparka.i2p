// server/middleware/csp.ts
// Nitro middleware to add privacy-first CSP headers for the site.
// Placed in server/middleware so Nitro picks it up automatically.

import { defineEventHandler, setResponseHeader } from "h3";

export default defineEventHandler((event) => {
  // Very strict but pragmatic CSP:
  // - default-src 'none' (deny everything by default)
  // - allow same-origin for navigation, images from data: and self,
  // - allow styles from self and inline (inline is used here for simple themes; remove if you change to nonce-based styles)
  // - disallow external scripts & trackers
  // Adjust to taste for I2P environment if necessary.
  const csp = [
    "default-src 'none'",
    "base-uri 'self'",
    "connect-src 'self' http: https:",
    "img-src 'self' data:",
    "style-src 'self' 'unsafe-inline'",
    "font-src 'self' data:",
    "script-src 'self'",
    "frame-ancestors 'none'",
    "block-all-mixed-content",
    "upgrade-insecure-requests",
  ].join("; ");

  setResponseHeader(event, "Content-Security-Policy", csp);
  setResponseHeader(event, "Referrer-Policy", "no-referrer");
  setResponseHeader(event, "X-Content-Type-Options", "nosniff");
  setResponseHeader(event, "X-Frame-Options", "DENY");
  setResponseHeader(event, "X-XSS-Protection", "1; mode=block");
});
