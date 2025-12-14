export default defineNuxtConfig({
  compatibilityDate: "2025-12-04",

  modules: [
    "@pinia/nuxt",
    "nuxt-feather-icons",
  ],

  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || "http://localhost:8080/api",
      siteName: process.env.NUXT_PUBLIC_SITE_NAME || "lilithinaparka.i2p",
      bskyHandle: process.env.NUXT_PUBLIC_BSKY_HANDLE ||
        "@lilithinaparka.bsky.social",
      defaultTheme: process.env.NUXT_PUBLIC_DEFAULT_THEME || "dracula",
    },
  },

  css: [
    "./app/assets/css/system.css",
    "./app/assets/css/colors.css",
    "./app/assets/css/themes.css",
    "./app/assets/css/main.css",
  ],

  app: {
    head: {
      title: "Lilith in a Parka",
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
        {
          name: "description",
          content: "TUI-inspired blog and games portal on I2P",
        },
      ],
      link: [
        { rel: "icon", type: "image/svg+xml", href: "/favicon.svg" },
      ],
    },
  },

  ssr: true,

  devtools: { enabled: true },
});
