import { defineStore } from "pinia";
import type { BlogPost } from "../types";

export const useBlogStore = defineStore("blog", {
  state: () => ({
    posts: [] as BlogPost[],
    currentPost: null as BlogPost | null,
    loading: false,
    error: null as string | null,
  }),

  actions: {
    async fetchPosts(includeDrafts = false) {
      this.loading = true;
      this.error = null;

      try {
        const config = useRuntimeConfig();
        const params = includeDrafts ? "?drafts=true" : "";
        const response = await fetch(
          `${config.public.apiBase}/blog/posts${params}`,
        );

        if (!response.ok) {
          throw new Error("Failed to fetch blog posts");
        }

        this.posts = await response.json();
      } catch (e) {
        this.error = e instanceof Error ? e.message : "Unknown error";
        console.error("Error fetching blog posts:", e);
      } finally {
        this.loading = false;
      }
    },

    async fetchPostBySlug(slug: string) {
      this.loading = true;
      this.error = null;

      try {
        const config = useRuntimeConfig();
        const response = await fetch(
          `${config.public.apiBase}/blog/posts/${slug}`,
        );

        if (!response.ok) {
          throw new Error("Failed to fetch blog post");
        }

        this.currentPost = await response.json();
        return this.currentPost;
      } catch (e) {
        this.error = e instanceof Error ? e.message : "Unknown error";
        console.error("Error fetching blog post:", e);
        return null;
      } finally {
        this.loading = false;
      }
    },

    async fetchPostsByCategory(category: string, includeDrafts = false) {
      this.loading = true;
      this.error = null;

      try {
        const config = useRuntimeConfig();
        const params = includeDrafts ? "?drafts=true" : "";
        const response = await fetch(
          `${config.public.apiBase}/blog/category/${category}${params}`,
        );

        if (!response.ok) {
          throw new Error("Failed to fetch blog posts");
        }

        this.posts = await response.json();
      } catch (e) {
        this.error = e instanceof Error ? e.message : "Unknown error";
        console.error("Error fetching blog posts by category:", e);
      } finally {
        this.loading = false;
      }
    },

    async fetchPostsByTag(tag: string, includeDrafts = false) {
      this.loading = true;
      this.error = null;

      try {
        const config = useRuntimeConfig();
        const params = includeDrafts ? "?drafts=true" : "";
        const response = await fetch(
          `${config.public.apiBase}/blog/tag/${tag}${params}`,
        );

        if (!response.ok) {
          throw new Error("Failed to fetch blog posts");
        }

        this.posts = await response.json();
      } catch (e) {
        this.error = e instanceof Error ? e.message : "Unknown error";
        console.error("Error fetching blog posts by tag:", e);
      } finally {
        this.loading = false;
      }
    },
  },

  getters: {
    publishedPosts: (state) => state.posts.filter((post) => !post.draft),
    postsByCategory: (state) => (category: string) => {
      return state.posts.filter((post) => {
        try {
          const categories = JSON.parse(post.categories);
          return categories.includes(category);
        } catch {
          return false;
        }
      });
    },
    postsByTag: (state) => (tag: string) => {
      return state.posts.filter((post) => {
        try {
          const tags = JSON.parse(post.tags);
          return tags.includes(tag);
        } catch {
          return false;
        }
      });
    },
  },
});
