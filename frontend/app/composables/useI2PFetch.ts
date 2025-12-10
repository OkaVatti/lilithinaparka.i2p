// composables/useI2PFetch.ts
import { useRuntimeConfig } from '#imports'

/**
 * Small composable to fetch from an I2P-hosted API.
 * Reads `runtimeConfig.public.i2pApiHost` (e.g. "http://profile.i2p:8080")
 *
 * Usage:
 *   const { fetchFromI2P, apiHost } = useI2PFetch()
 *   const data = await fetchFromI2P('/api/profile')
 */
export default function useI2PFetch() {
  const config = useRuntimeConfig()
  const apiHost = (config.public && config.public.i2pApiHost) || ''

  if (!apiHost) {
    // don't throw here — the UI will show errors — but warn in console
    // (You should set runtime config: public.i2pApiHost)
    // e.g. PUBLIC I2P host: "http://profile.i2p:8080" or the correct I2P HTTP proxy address.
    // In development you can also set an HTTP-to-I2P proxy URL.
    console.warn('[useI2PFetch] runtime public.i2pApiHost is not set.')
  }

  async function fetchFromI2P(path: string, opts: RequestInit = {}, timeoutMs = 8000) {
    if (!apiHost) throw new Error('I2P API host not configured (public.i2pApiHost)')
    const url = new URL(path, apiHost).toString()

    const controller = new AbortController()
    const id = setTimeout(() => controller.abort(), timeoutMs)
    try {
      const res = await fetch(url, { ...opts, signal: controller.signal })
      clearTimeout(id)

      if (!res.ok) {
        const text = await res.text().catch(() => '')
        throw new Error(`network ${res.status}: ${text || res.statusText}`)
      }

      const ct = res.headers.get('content-type') || ''
      if (ct.includes('application/json')) {
        return res.json()
      }
      // return text fallback
      return res.text()
    } catch (err: any) {
      if (err?.name === 'AbortError') throw new Error('request timed out')
      throw err
    }
  }

  return {
    fetchFromI2P,
    apiHost,
  }
}
