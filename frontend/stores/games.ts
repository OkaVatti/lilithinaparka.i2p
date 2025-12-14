// stores/games.ts
import { defineStore } from "pinia";
import { ref } from "vue";
import { useApi } from "~/composables/useApi";
import type { Game, GameScore } from "~~/types";

export const useGamesStore = defineStore("games", () => {
  const games = ref<Game[]>([]);
  const currentGame = ref<Game | null>(null);
  const loading = ref(false);
  const error = ref<string | null>(null);

  const { apiFetch } = useApi();

  async function loadAll() {
    loading.value = true;
    error.value = null;
    try {
      const res = await apiFetch<Game[]>("/games");
      games.value = res || [];
    } catch (e: any) {
      error.value = e.message;
    } finally {
      loading.value = false;
    }
  }

  async function loadBySlug(slug: string) {
    try {
      const res = await apiFetch<Game>(`/games/${encodeURIComponent(slug)}`);
      currentGame.value = res;
      return res;
    } catch (e: any) {
      error.value = e.message;
      return null;
    }
  }

  async function getLeaderboard(slug: string, params = "") {
    return await apiFetch<GameScore[]>(
      `/games/${encodeURIComponent(slug)}/leaderboard${params ? params : ""}`,
    );
  }

  async function getRecent(slug: string, params = "") {
    return await apiFetch<GameScore[]>(
      `/games/${encodeURIComponent(slug)}/recent${params ? params : ""}`,
    );
  }

  async function submitScore(
    slug: string,
    payload: { alias: string; score: number; metadata?: any },
  ) {
    return await apiFetch(`/games/${encodeURIComponent(slug)}/score`, {
      method: "POST",
      body: payload,
    });
  }

  return {
    games,
    currentGame,
    loading,
    error,
    loadAll,
    loadBySlug,
    getLeaderboard,
    getRecent,
    submitScore,
  };
});
