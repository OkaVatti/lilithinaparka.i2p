import { defineStore } from 'pinia'
import type { BskyPost } from '../types'

export const useBskyStore = defineStore('bsky', {
  state: () => ({
    posts: [] as BskyPost[],
    loading: false,
    error: null as string | null
  }),
  
  actions: {
    async fetchPosts(limit = 50) {
      this.loading = true
      this.error = null
      
      try {
        const config = useRuntimeConfig()
        const response = await fetch(`${config.public.apiBase}/bsky/posts?limit=${limit}`)
        
        if (!response.ok) {
          throw new Error('Failed to fetch BlueSky posts')
        }
        
        this.posts = await response.json()
      } catch (e) {
        this.error = e instanceof Error ? e.message : 'Unknown error'
        console.error('Error fetching BlueSky posts:', e)
      } finally {
        this.loading = false
      }
    }
  }
})