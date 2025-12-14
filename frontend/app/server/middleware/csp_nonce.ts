// server/middleware/csp-nonce.ts
// Nitro middleware: create per-request nonce, attach to event.context, and set CSP headers using that nonce.

import { defineEventHandler, setResponseHeader } from "h3";
import { randomBytes } from "crypto";

function makeNonce(len = 16) {
  return randomBytes(len).toString("base64");
}

export default defineEventHandler((event) => {
  // create a per-request nonce
  const nonce = makeNonce();
  // attach to context so later server plugins/components can access it
  event.context.cspNonce = nonce;

  // Build a CSP that uses the nonce for inline style/script and otherwise is restrictive.
  // style-src and script-src use the generated nonce (safer than 'unsafe-inline').
  const styleNonce = `'nonce-${nonce}'`;
  const scriptNonce = `'nonce-${nonce}'`;

  const csp = [
    "default-src 'none'",
    "base-uri 'self'",
    `connect-src 'self' https: http:`,
    "img-src 'self' data:",
    `style-src 'self' ${styleNonce}`,
    "font-src 'self' data:",
    `script-src 'self' ${scriptNonce}`,
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
