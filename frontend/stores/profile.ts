import { defineStore } from "pinia";
import type { Profile } from "~~/types";

export const useProfileStore = defineStore("profile", {
  state: () => ({
    profile: null as Profile | null,
    loading: false,
    error: null as string | null,
  }),

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

  actions: {
    async fetchProfile() {
      this.loading = true;
      this.error = null;

      try {
        const { apiFetch } = useApi();
        this.profile = await apiFetch<Profile>("/profile");
      } catch (err) {
        this.error = err instanceof Error
          ? err.message
          : "Failed to fetch profile";
      } finally {
        this.loading = false;
      }
    },

    async refreshProfile() {
      this.loading = true;
      this.error = null;

      try {
        const { apiFetch } = useApi();
        await apiFetch("/profile/refresh", { method: "POST" });
        await this.fetchProfile();
      } catch (err) {
        this.error = err instanceof Error
          ? err.message
          : "Failed to refresh profile";
      } finally {
        this.loading = false;
      }
    },
  },
});
