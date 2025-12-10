import { defineStore } from "pinia";
import type { BlogPost } from "~~/types";

export const useBlogStore = defineStore("blog", {
  state: () => ({
    posts: [] as BlogPost[],
    currentPost: null as BlogPost | null,
    loading: false,
    error: null as string | null,
  }),

  getters: {
    publishedPosts: (state) => state.posts.filter((post) => !post.draft),

    postsByCategory: (state) => (category: string) =>
      state.posts.filter((post) =>
        post.categories.includes(category) && !post.draft
      ),

    postsByTag: (state) => (tag: string) =>
      state.posts.filter((post) => post.tags.includes(tag) && !post.draft),
  },

  actions: {
    async fetchPosts(includeDrafts = false) {
      this.loading = true;
      this.error = null;

      try {
        const { apiFetch } = useApi();
        const query = includeDrafts ? "?drafts=true" : "";
        this.posts = await apiFetch<BlogPost[]>(`/blog/posts${query}`);
      } catch (err) {
        this.error = err instanceof Error
          ? err.message
          : "Failed to fetch posts";
      } finally {
        this.loading = false;
      }
    },

    async fetchPostBySlug(slug: string) {
      this.loading = true;
      this.error = null;

      try {
        const { apiFetch } = useApi();
        this.currentPost = await apiFetch<BlogPost>(`/blog/posts/${slug}`);
      } catch (err) {
        this.error = err instanceof Error
          ? err.message
          : "Failed to fetch post";
        this.currentPost = null;
      } finally {
        this.loading = false;
      }
    },

    async fetchPostsByCategory(category: string, includeDrafts = false) {
      this.loading = true;
      this.error = null;

      try {
        const { apiFetch } = useApi();
        const query = includeDrafts ? "?drafts=true" : "";
        this.posts = await apiFetch<BlogPost[]>(
          `/blog/category/${category}${query}`,
        );
      } catch (err) {
        this.error = err instanceof Error
          ? err.message
          : "Failed to fetch posts";
      } finally {
        this.loading = false;
      }
    },

    async fetchPostsByTag(tag: string, includeDrafts = false) {
      this.loading = true;
      this.error = null;

      try {
        const { apiFetch } = useApi();
        const query = includeDrafts ? "?drafts=true" : "";
        this.posts = await apiFetch<BlogPost[]>(`/blog/tag/${tag}${query}`);
      } catch (err) {
        this.error = err instanceof Error
          ? err.message
          : "Failed to fetch posts";
      } finally {
        this.loading = false;
      }
    },
  },
});
