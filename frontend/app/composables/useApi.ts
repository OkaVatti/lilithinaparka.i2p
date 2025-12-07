// frontend/app/composables/useApi.ts
import type { BlogPost, BskyPost, Profile } from '../../types'

interface ApiError {
  error: string
  message?: string
}

interface ApiResponse<T> {
  data: T | null
  error: ApiError | null
  loading: boolean
}

export const useApi = () => {
  const config = useRuntimeConfig()
  const baseURL = config.public.apiBase

  const fetchWithRetry = async <T>(
    url: string,
    options: RequestInit = {},
    retries = 3,
    delay = 1000
  ): Promise<T> => {
    for (let i = 0; i < retries; i++) {
      try {
        const response = await fetch(`${baseURL}${url}`, {
          ...options,
          headers: {
            'Content-Type': 'application/json',
            ...options.headers,
          },
        })

        if (!response.ok) {
          const errorData = await response.json().catch(() => ({
            error: `HTTP ${response.status}: ${response.statusText}`,
          }))
          throw new Error(errorData.error || `Request failed with status ${response.status}`)
        }

        return await response.json()
      } catch (error) {
        if (i === retries - 1) throw error
        await new Promise(resolve => setTimeout(resolve, delay * (i + 1)))
      }
    }
    throw new Error('Max retries exceeded')
  }

  // Blog API
  const blog = {
    async getPosts(includeDrafts = false): Promise<BlogPost[]> {
      const params = includeDrafts ? '?drafts=true' : ''
      return fetchWithRetry<BlogPost[]>(`/blog/posts${params}`)
    },

    async getPostBySlug(slug: string): Promise<BlogPost> {
      return fetchWithRetry<BlogPost>(`/blog/posts/${slug}`)
    },

    async getPostsByCategory(category: string, includeDrafts = false): Promise<BlogPost[]> {
      const params = includeDrafts ? '?drafts=true' : ''
      return fetchWithRetry<BlogPost[]>(`/blog/category/${category}${params}`)
    },

    async getPostsByTag(tag: string, includeDrafts = false): Promise<BlogPost[]> {
      const params = includeDrafts ? '?drafts=true' : ''
      return fetchWithRetry<BlogPost[]>(`/blog/tag/${tag}${params}`)
    },

    async rescanPosts(): Promise<{ message: string }> {
      return fetchWithRetry<{ message: string }>('/blog/rescan', {
        method: 'POST',
      })
    },
  }

  // BlueSky API
  const bsky = {
    async getPosts(limit = 50): Promise<BskyPost[]> {
      return fetchWithRetry<BskyPost[]>(`/bsky/posts?limit=${limit}`)
    },

    async getPostByURI(uri: string): Promise<BskyPost> {
      return fetchWithRetry<BskyPost>(`/bsky/post?uri=${encodeURIComponent(uri)}`)
    },

    async refreshPosts(): Promise<{ message: string }> {
      return fetchWithRetry<{ message: string }>('/bsky/refresh', {
        method: 'POST',
      })
    },
  }

  // Profile API
  const profile = {
    async getProfile(): Promise<Profile> {
      return fetchWithRetry<Profile>('/profile')
    },

    async refreshProfile(): Promise<{ message: string }> {
      return fetchWithRetry<{ message: string }>('/profile/refresh', {
        method: 'POST',
      })
    },
  }

  // Health check
  const health = {
    async check(): Promise<{ status: string; time: string }> {
      return fetchWithRetry<{ status: string; time: string }>('/health')
    },
  }

  return {
    blog,
    bsky,
    profile,
    health,
    fetchWithRetry,
  }
}