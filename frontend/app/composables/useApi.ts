// app/composables/useApi.ts
import { ref } from 'vue'

export function useApi() {
  const config = useRuntimeConfig()
  // NUXT_PUBLIC_API_BASE expected to include /api (e.g. http://localhost:8080/api)
  const base = (config.public.apiBase || '').replace(/\/$/, '')

  const getToken = () => {
    try {
      return process.client ? localStorage.getItem('auth:token') : null
    } catch {
      return null
    }
  }

  const setToken = (t: string | null) => {
    try {
      if (!process.client) return
      if (t) localStorage.setItem('auth:token', t)
      else localStorage.removeItem('auth:token')
    } catch {}
  }

  function buildUrl(path: string) {
    if (!path) return base || '/api'
    if (/^https?:\/\//.test(path)) return path
    // allow calling with or without leading slash
    if (path.startsWith('/')) {
      return `${base}${path}`
    } else {
      return `${base}/${path}`
    }
  }

  async function apiFetch<T = any>(path: string, opts: RequestInit = {}): Promise<T> {
    const url = buildUrl(path)
    const headers: Record<string, string> = (opts.headers as any) || {}

    // attach auth token
    const token = getToken()
    if (token) headers['Authorization'] = `Bearer ${token}`

    // don't override Content-Type for FormData (browser sets it)
    if (opts.body && !(opts.body instanceof FormData) && !headers['Content-Type']) {
      headers['Content-Type'] = 'application/json'
    }

    const final: RequestInit = {
      method: opts.method || (opts.body ? 'POST' : 'GET'),
      ...opts,
      headers,
    }

    // auto-stringify JSON bodies
    if (final.body && typeof final.body !== 'string' && !(final.body instanceof FormData)) {
      try {
        final.body = JSON.stringify(final.body)
      } catch (e) {
        // leave as-is
      }
    }

    const res = await fetch(url, final)
    const contentType = res.headers.get('content-type') || ''

    // handle non-JSON responses gracefully
    if (!res.ok) {
      let bodyText = ''
      try {
        if (contentType.includes('application/json')) {
          const json = await res.json()
          bodyText = JSON.stringify(json)
        } else {
          bodyText = await res.text()
        }
      } catch (e) {
        bodyText = res.statusText || 'error'
      }

      // clear token on 401
      if (res.status === 401) setToken(null)

      throw new Error(`API error ${res.status}: ${bodyText}`)
    }

    if (contentType.includes('application/json')) {
      return res.json()
    } else {
      // fallback: return text
      return (await res.text()) as unknown as T
    }
  }

  // Convenience: multipart upload (expects FormData)
  async function upload<T = any>(path: string, form: FormData, opts: RequestInit = {}): Promise<T> {
    const headers = (opts.headers as any) || {}
    // ensure auth header
    const token = getToken()
    if (token) headers['Authorization'] = `Bearer ${token}`

    // DO NOT set Content-Type for FormData
    const res = await fetch(buildUrl(path), {
      method: 'POST',
      body: form,
      ...opts,
      headers,
    })

    if (!res.ok) {
      const body = await res.text().catch(() => '')
      throw new Error(`Upload failed ${res.status}: ${body}`)
    }
    return (await res.json()) as T
  }

  // Server-Sent Events helper: returns a function to close the connection
  function sse(path: string, onMessage: (ev: MessageEvent) => void, onOpen?: () => void, onError?: (err: any) => void) {
    const url = buildUrl(path).replace(/^http:/, location.protocol === 'https:' ? 'https:' : 'http:')
    // For local dev, EventSource must be same-origin or CORS-supporting on server
    const es = new EventSource(url)
    es.onmessage = onMessage
    es.onopen = () => {
      if (onOpen) onOpen()
    }
    es.onerror = (err) => {
      if (onError) onError(err)
    }
    return () => es.close()
  }

  return {
    apiFetch,
    upload,
    sse,
    getToken,
    setToken
  }
}
