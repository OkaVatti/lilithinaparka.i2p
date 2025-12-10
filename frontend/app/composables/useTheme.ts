import type { ThemeOption } from "~~/types";

export const useTheme = () => {
  const currentTheme = useState<string>("theme", () => "dracula");

  const themes: ThemeOption[] = [
    {
      name: "Dracula",
      value: "dracula",
      preview: { bg: "#282a36", fg: "#f8f8f2", primary: "#bd93f9" },
    },
    {
      name: "Nord",
      value: "nord",
      preview: { bg: "#2e3440", fg: "#eceff4", primary: "#81a1c1" },
    },
    {
      name: "Gruvbox",
      value: "gruvbox",
      preview: { bg: "#282828", fg: "#ebdbb2", primary: "#458588" },
    },
    {
      name: "Tokyo Night",
      value: "tokyo",
      preview: { bg: "#1a1b26", fg: "#c0caf5", primary: "#7aa2f7" },
    },
    {
      name: "Monokai",
      value: "monokai",
      preview: { bg: "#272822", fg: "#f8f8f2", primary: "#ae81ff" },
    },
  ];

  const setTheme = (theme: string) => {
    currentTheme.value = theme;
    if (process.client) {
      document.documentElement.setAttribute("data-theme", theme);
      localStorage.setItem("theme", theme);
    }
  };

  const initTheme = () => {
    if (process.client) {
      const saved = localStorage.getItem("theme");
      if (saved) {
        setTheme(saved);
      } else {
        setTheme("dracula");
      }
    }
  };

  return {
    currentTheme: readonly(currentTheme),
    themes,
    setTheme,
    initTheme,
  };
};
