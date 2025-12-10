import { defineStore } from "pinia";
import type { BskyPost } from "~~/types";

export const useBskyStore = defineStore("bsky", {
  state: () => ({
    posts: [] as BskyPost[],
    loading: false,
    error: null as string | null,
  }),

  actions: {
    async fetchPosts(limit = 50) {
      this.loading = true;
      this.error = null;

      try {
        const { apiFetch } = useApi();
        this.posts = await apiFetch<BskyPost[]>(`/bsky/posts?limit=${limit}`);
      } catch (err) {
        this.error = err instanceof Error
          ? err.message
          : "Failed to fetch BlueSky posts";
      } finally {
        this.loading = false;
      }
    },

    async refreshPosts() {
      this.loading = true;
      this.error = null;

      try {
        const { apiFetch } = useApi();
        await apiFetch("/bsky/refresh", { method: "POST" });
        await this.fetchPosts();
      } catch (err) {
        this.error = err instanceof Error
          ? err.message
          : "Failed to refresh posts";
      } finally {
        this.loading = false;
      }
    },
  },
});
