import { defineStore } from "pinia";
import type { Game, GameScore } from "~/types";

export const useGamesStore = defineStore("games", {
  state: () => ({
    games: [] as Game[],
    currentGame: null as Game | null,
    leaderboard: [] as GameScore[],
    loading: false,
    error: null as string | null,
  }),

  actions: {
    async fetchGames() {
      this.loading = true;
      this.error = null;

      try {
        const { apiFetch } = useApi();
        this.games = await apiFetch<Game[]>("/games");
      } catch (err) {
        this.error = err instanceof Error
          ? err.message
          : "Failed to fetch games";
      } finally {
        this.loading = false;
      }
    },

    async fetchGameBySlug(slug: string) {
      this.loading = true;
      this.error = null;

      try {
        const { apiFetch } = useApi();
        this.currentGame = await apiFetch<Game>(`/games/${slug}`);
      } catch (err) {
        this.error = err instanceof Error
          ? err.message
          : "Failed to fetch game";
        this.currentGame = null;
      } finally {
        this.loading = false;
      }
    },

    async fetchLeaderboard(slug: string, limit = 100) {
      this.loading = true;
      this.error = null;

      try {
        const { apiFetch } = useApi();
        this.leaderboard = await apiFetch<GameScore[]>(
          `/games/${slug}/leaderboard?limit=${limit}`,
        );
      } catch (err) {
        this.error = err instanceof Error
          ? err.message
          : "Failed to fetch leaderboard";
      } finally {
        this.loading = false;
      }
    },

    async submitScore(
      slug: string,
      alias: string,
      score: number,
      level = 1,
      data = "{}",
    ) {
      this.loading = true;
      this.error = null;

      try {
        const { apiFetch } = useApi();
        await apiFetch(`/games/${slug}/score`, {
          method: "POST",
          body: JSON.stringify({ alias, score, level, data }),
        });
      } catch (err) {
        this.error = err instanceof Error
          ? err.message
          : "Failed to submit score";
        throw err;
      } finally {
        this.loading = false;
      }
    },
  },
});
