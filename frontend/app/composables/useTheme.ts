import { computed, ref } from "vue";
import type { ThemeOption } from "~~/types";

export const useTheme = () => {
  const currentTheme = useState<string>("theme", () => "dracula");
  const currentColorScheme = useState<string>("colorscheme", () => "dark");

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
    {
      name: "Solarized Dark",
      value: "solarized",
      preview: { bg: "#002b36", fg: "#839496", primary: "#268bd2" },
    },
    {
      name: "Rose Pine",
      value: "rosepine",
      preview: { bg: "#191724", fg: "#e0def4", primary: "#c4a7e7" },
    },
    {
      name: "Catppuccin",
      value: "catppuccin",
      preview: { bg: "#1e1e2e", fg: "#cdd6f4", primary: "#cba6f7" },
    },
    {
      name: "Everforest",
      value: "everforest",
      preview: { bg: "#2b3339", fg: "#d3c6aa", primary: "#a7c080" },
    },
    {
      name: "Kanagawa",
      value: "kanagawa",
      preview: { bg: "#1f1f28", fg: "#dcd7ba", primary: "#7e9cd8" },
    },
    {
      name: "Terminal Green",
      value: "terminal",
      preview: { bg: "#0a0a0a", fg: "#00ff00", primary: "#00aa00" },
    },
    {
      name: "Autumn",
      value: "autumn",
      preview: { bg: "#1b1b17", fg: "#f7eee1", primary: "#d65a31" },
    },
  ];

  const seasonalThemes = [
    { name: "Winter", value: "winter" },
    { name: "Spring", value: "spring" },
    { name: "Summer", value: "summer" },
    { name: "Autumn", value: "autumn" },
  ];

  const setTheme = (theme: string) => {
    currentTheme.value = theme;
    if (process.client) {
      document.documentElement.setAttribute("data-theme", theme);
      localStorage.setItem("theme", theme);
      applyThemeVariables(theme);
    }
  };

  const setColorScheme = (scheme: string) => {
    currentColorScheme.value = scheme;
    if (process.client) {
      document.documentElement.setAttribute("data-colorscheme", scheme);
      localStorage.setItem("colorscheme", scheme);
    }
  };

  const applyThemeVariables = (themeName: string) => {
    const themeConfig = themes.find((t) => t.value === themeName);
    if (!themeConfig || !process.client) return;

    const root = document.documentElement;
    root.style.setProperty("--theme-bg", themeConfig.preview.bg);
    root.style.setProperty("--theme-fg", themeConfig.preview.fg);
    root.style.setProperty("--theme-primary", themeConfig.preview.primary);

    // Derive additional colors
    root.style.setProperty(
      "--theme-surface",
      adjustColor(themeConfig.preview.bg, 10),
    );
    root.style.setProperty(
      "--theme-border",
      adjustColor(themeConfig.preview.bg, 20),
    );
    root.style.setProperty("--theme-accent", themeConfig.preview.primary);
    root.style.setProperty(
      "--theme-muted",
      adjustColor(themeConfig.preview.fg, -30),
    );
  };

  const adjustColor = (hex: string, percent: number): string => {
    const num = parseInt(hex.replace("#", ""), 16);
    const amt = Math.round(2.55 * percent);
    const R = (num >> 16) + amt;
    const G = (num >> 8 & 0x00FF) + amt;
    const B = (num & 0x0000FF) + amt;
    return "#" + (0x1000000 + (R < 255 ? R < 1 ? 0 : R : 255) * 0x10000 +
      (G < 255 ? G < 1 ? 0 : G : 255) * 0x100 +
      (B < 255 ? B < 1 ? 0 : B : 255))
      .toString(16).slice(1);
  };

  const initTheme = () => {
    if (process.client) {
      const savedTheme = localStorage.getItem("theme");
      const savedScheme = localStorage.getItem("colorscheme");

      if (savedTheme && themes.some((t) => t.value === savedTheme)) {
        setTheme(savedTheme);
      } else {
        setTheme("dracula");
      }

      if (savedScheme) {
        setColorScheme(savedScheme);
      } else {
        setColorScheme("dark");
      }
    }
  };

  const renderVarsForTheme = () => {
    const theme = themes.find((t) => t.value === currentTheme.value) ||
      themes[0];
    return `:root {
      --theme-bg: ${theme.preview.bg};
      --theme-fg: ${theme.preview.fg};
      --theme-primary: ${theme.preview.primary};
      --theme-surface: ${adjustColor(theme.preview.bg, 10)};
      --theme-border: ${adjustColor(theme.preview.bg, 20)};
      --theme-accent: ${theme.preview.primary};
      --theme-muted: ${adjustColor(theme.preview.fg, -30)};
    }`;
  };

  return {
    currentTheme: readonly(currentTheme),
    currentColorScheme: readonly(currentColorScheme),
    themes,
    seasonalThemes,
    setTheme,
    setColorScheme,
    initTheme,
    renderVarsForTheme,
  };
};
