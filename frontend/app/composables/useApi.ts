import { useAuth } from "./useAuth";

export const useApi = () => {
  const config = useRuntimeConfig();
  const auth = useAuth();

  const apiFetch = async <T>(
    endpoint: string,
    options: RequestInit = {},
  ): Promise<T> => {
    const headers = new Headers(options.headers || {});
    headers.set("Content-Type", "application/json");

    if (auth.token.value) {
      headers.set("Authorization", `Bearer ${auth.token.value}`);
    }

    try {
      const response = await fetch(`${config.public.apiBase}${endpoint}`, {
        ...options,
        headers,
      });

      // Handle 401 Unauthorized
      if (response.status === 401) {
        auth.logout();
        throw new Error("Session expired. Please login again.");
      }

      if (!response.ok) {
        const errorText = await response.text();
        let errorData;
        try {
          errorData = JSON.parse(errorText);
        } catch {
          errorData = { error: errorText };
        }
        throw new Error(
          errorData.error ||
            `API Error: ${response.status} ${response.statusText}`,
        );
      }

      // Handle empty responses
      const contentType = response.headers.get("content-type");
      if (contentType?.includes("application/json")) {
        return await response.json();
      }

      return (await response.text()) as unknown as T;
    } catch (error) {
      console.error("API request failed:", error);
      throw error;
    }
  };

  const uploadFile = async (endpoint: string, file: File, data: any = {}) => {
    const formData = new FormData();
    formData.append("file", file);

    Object.keys(data).forEach((key) => {
      formData.append(key, data[key]);
    });

    const headers = new Headers();
    if (auth.token.value) {
      headers.set("Authorization", `Bearer ${auth.token.value}`);
    }

    const response = await fetch(`${config.public.apiBase}${endpoint}`, {
      method: "POST",
      headers,
      body: formData,
    });

    if (!response.ok) {
      throw new Error("Upload failed");
    }

    return response.json();
  };

  return {
    apiFetch,
    uploadFile,
    baseURL: config.public.apiBase,
  };
};
