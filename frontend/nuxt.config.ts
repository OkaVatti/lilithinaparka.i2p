export default defineNuxtConfig({
  compatibilityDate: "2025-12-04",

  modules: [
    "@pinia/nuxt",
    "nuxt-feather-icons",
  ],

  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || "http://localhost:8080/api",
      i2pApiHost: process.env.NUXT_I2P_API_TUNNEL || "https://127.0.0.1:7667"
    },
  },

  css: ["./app/assets/css/main.css"],

  app: {
    head: {
      title: "Lilith Parker",
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
        {
          name: "description",
          content:
            "chronically online cat-girl who likes programming, music, and burritos",
        },
      ],
      link: [
        { rel: "icon", type: "image/x-icon", href: "/favicon.ico" },
      ],
    },
  },

  nitro: {
    preset: "node-server",
  },
});
