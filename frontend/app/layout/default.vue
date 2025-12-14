<template>
  <div class="app-root">
    <Header />
    <!-- SSR inline style with nonce to reduce FOUC; plugin/middleware inserts payload.cspNonce -->
    <style v-if="initialCss" v-bind="styleAttrs" v-html="initialCss"></style>

    <main class="container" role="main">
      <slot />
    </main>

    <footer class="container footer" role="contentinfo">
      <div>
        <small>© Lilith in a Parka — privacy-first portal</small>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
const { renderVarsForTheme } = useTheme()

// Build initial CSS for SSR to avoid FOUC. Note: renderVarsForTheme reads the SSR-initialized useState/useCookie.
const initialCss = renderVarsForTheme()

// read nonce injected into payload by server plugin/middleware
const nuxtApp = useNuxtApp()
const nonce = nuxtApp.payload?.cspNonce ?? null

// For the style tag we need to set the nonce attribute if available.
const styleAttrs = computed(() => {
  if (!nonce) return {}
  // Vue will render 'nonce' attribute
  return { nonce }
})
</script>

<style src="~/app/styles/themes.css"></style>

<style scoped>
.app-root {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}
main.container {
  flex: 1;
  margin: 2rem auto;
  max-width: 1100px;
}
footer.footer {
  margin-top: 2rem;
  padding: 1rem 0;
  border-top: 1px dashed var(--theme-border);
  text-align: center;
  color: var(--theme-muted);
}
</style>
