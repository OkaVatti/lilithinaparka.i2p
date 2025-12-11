export const useAuth = () => {
  const token = useState<string | null>("auth_token", () => null);
  const user = useState<any>("auth_user", () => null);

  const login = async (username: string, password: string) => {
    try {
      const config = useRuntimeConfig();
      const response = await fetch(`${config.public.apiBase}/auth/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
      });

      if (!response.ok) {
        throw new Error("Login failed");
      }

      const data = await response.json();
      token.value = data.token;
      user.value = data.user;

      // Store in localStorage for persistence
      if (process.client) {
        localStorage.setItem("auth_token", data.token);
        localStorage.setItem("auth_user", JSON.stringify(data.user));
      }

      return data;
    } catch (error) {
      console.error("Login error:", error);
      throw error;
    }
  };

  const logout = () => {
    token.value = null;
    user.value = null;

    if (process.client) {
      localStorage.removeItem("auth_token");
      localStorage.removeItem("auth_user");
    }
  };

  const checkAuth = async () => {
    if (!token.value && process.client) {
      const storedToken = localStorage.getItem("auth_token");
      const storedUser = localStorage.getItem("auth_user");

      if (storedToken && storedUser) {
        token.value = storedToken;
        user.value = JSON.parse(storedUser);
      }
    }

    if (token.value) {
      try {
        const config = useRuntimeConfig();
        const response = await fetch(`${config.public.apiBase}/auth/validate`, {
          headers: {
            "Authorization": `Bearer ${token.value}`,
          },
        });

        if (!response.ok) {
          logout();
          return false;
        }

        return true;
      } catch (error) {
        logout();
        return false;
      }
    }

    return false;
  };

  const isAuthenticated = computed(() => !!token.value);
  const isAdmin = computed(() => user.value?.role === "admin");

  return {
    token: readonly(token),
    user: readonly(user),
    login,
    logout,
    checkAuth,
    isAuthenticated,
    isAdmin,
  };
};
