export const useApi = () => {
  const config = useRuntimeConfig();
  const baseURL = config.public.apiBase;

  const apiFetch = async <T>(
    endpoint: string,
    options: RequestInit = {},
  ): Promise<T> => {
    try {
      const response = await fetch(`${baseURL}${endpoint}`, {
        ...options,
        headers: {
          "Content-Type": "application/json",
          ...options.headers,
        },
      });

      if (!response.ok) {
        throw new Error(`API Error: ${response.status} ${response.statusText}`);
      }

      return await response.json();
    } catch (error) {
      console.error("API request failed:", error);
      throw error;
    }
  };

  return {
    apiFetch,
    baseURL,
  };
};
