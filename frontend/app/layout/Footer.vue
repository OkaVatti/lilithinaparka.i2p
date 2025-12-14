<template>
  <aside class="sidebar" role="navigation" aria-label="Main navigation">
    <div class="brand">
      <img :src="logo" alt="logo" class="brand-logo" v-if="logo"/>
      <h1 class="brand-title">Lilith</h1>
    </div>

    <nav class="nav">
      <NuxtLink :to="'/'" class="nav-link" active-class="active">Home</NuxtLink>
      <NuxtLink :to="'/blog'" class="nav-link" active-class="active">Blog</NuxtLink>
      <NuxtLink :to="'/art'" class="nav-link" active-class="active">Art</NuxtLink>
      <NuxtLink :to="'/profile'" class="nav-link" active-class="active">Profile</NuxtLink>
      <NuxtLink :to="'/about'" class="nav-link" active-class="active">About</NuxtLink>
    </nav>

    <div class="sidebar-footer">
      <slot name="footer">
        <button class="btn btn-primary" @click="toggleTheme">Toggle theme</button>
      </slot>
    </div>
  </aside>
</template>

<script setup lang="ts">
import { useTheme } from '~/composables/useTheme'
const props = defineProps({
  logo: { type: String, default: '/icon-192.png' }
})

const { currentTheme, themes, setTheme } = useTheme()

function toggleTheme() {
  const list = themes.value.map(t => t.value)
  const idx = list.indexOf(currentTheme.value)
  const next = list[(idx + 1) % list.length]
  setTheme(next)
}
</script>

<style scoped>
@import '~/system.css';

.sidebar {
  width: 220px;
  padding: 1rem;
  border-right: 1px solid var(--theme-border);
  background: linear-gradient(180deg, color-mix(in srgb, var(--theme-surface) 90%, transparent), transparent);
  min-height: 100vh;
  box-sizing: border-box;
}
.brand {
  display:flex;
  align-items:center;
  gap:.6rem;
  margin-bottom: 1.1rem;
}
.brand-logo {
  width:40px; height:40px; border-radius:6px; object-fit:cover;
  border:1px solid var(--theme-border);
}
.brand-title {
  font-size:1.05rem;
  color: var(--theme-fg);
  margin:0;
}
.nav {
  display:flex;
  flex-direction:column;
  gap:.45rem;
}
.nav-link {
  padding: .5rem .6rem;
  border-radius:6px;
  color:var(--theme-fg);
  text-decoration:none;
}
.nav-link.active, .nav-link:hover {
  background: var(--theme-primary);
  color: white;
}
.sidebar-footer {
  margin-top: auto;
  padding-top: 1rem;
  border-top: 1px dashed var(--theme-border);
}
.btn {
  width:100%;
}
</style>
