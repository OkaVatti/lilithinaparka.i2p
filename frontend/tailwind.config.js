/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./app.vue",
  ],
  theme: {
    extend: {
      colors: {
        "bg-primary": "#0a0a0a",
        "bg-secondary": "#1a1a1a",
        "text-primary": "#e0e0e0",
        "text-secondary": "#a0a0a0",
        "accent": "#8b8be9",
        "link": "#7568ff",
        "border": "#333333",
      },
      fontFamily: {
        "mono": ["Courier New", "monospace"],
        "sans": ["Arial", "sans-serif"],
      },
      typography: {
        DEFAULT: {
          css: {
            color: "#e0e0e0",
            a: {
              color: "#b0afed",
              "&:hover": {
                color: "#8b8be9",
              },
            },
            h1: {
              color: "#e0e0e0",
            },
            h2: {
              color: "#e0e0e0",
            },
            h3: {
              color: "#e0e0e0",
            },
            code: {
              color: "#8b8be9",
              backgroundColor: "#1a1a1a",
              padding: "0.2em 0.4em",
              borderRadius: "3px",
            },
            "code::before": {
              content: '""',
            },
            "code::after": {
              content: '""',
            },
            pre: {
              backgroundColor: "#1a1a1a",
              color: "#e0e0e0",
            },
            blockquote: {
              color: "#a0a0a0",
              borderLeftColor: "#333333",
            },
          },
        },
      },
    },
  },
  plugins: [
    require("@tailwindcss/typography"),
  ],
};
