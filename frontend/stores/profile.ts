import { defineStore } from "pinia";
import type { Profile } from "../types";

export const useProfileStore = defineStore("profile", {
  state: () => ({
    profile: null as Profile | null,
    loading: false,
    error: null as string | null,
  }),

  actions: {
    async fetchProfile() {
      this.loading = true;
      this.error = null;

      try {
        const config = useRuntimeConfig();
        const response = await fetch(`${config.public.apiBase}/profile`);

        if (!response.ok) {
          throw new Error("Failed to fetch profile");
        }

        this.profile = await response.json();
      } catch (e) {
        this.error = e instanceof Error ? e.message : "Unknown error";
        console.error("Error fetching profile:", e);
      } finally {
        this.loading = false;
      }
    },
  },

  getters: {
    interests: (state) => {
      if (!state.profile?.interests) return [];
      try {
        return JSON.parse(state.profile.interests);
      } catch {
        return [];
      }
    },
  },
});
