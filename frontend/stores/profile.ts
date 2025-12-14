import { defineStore } from "pinia";
import type { Profile } from "~~/types";

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
        const { apiFetch } = useApi();
        this.profile = await apiFetch<Profile>("/profile");
      } catch (err) {
        this.error = err instanceof Error
          ? err.message
          : "Failed to fetch profile";
        this.profile = null;
      } finally {
        this.loading = false;
      }
    },

    async updateProfile(data: Partial<Profile>) {
      this.loading = true;
      this.error = null;

      try {
        const { apiFetch } = useApi();
        this.profile = await apiFetch<Profile>("/profile", {
          method: "PUT",
          body: data,
        });
      } catch (err) {
        this.error = err instanceof Error
          ? err.message
          : "Failed to update profile";
      } finally {
        this.loading = false;
      }
    },
  },
});
