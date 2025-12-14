<template>
  <nav class="sidebar" aria-label="Main navigation">
    <div class="brand">
      <NuxtLink to="/" class="brand-link">Lilith<span class="dot">.</span></NuxtLink>
    </div>

    <ul class="nav-list">
      <li><NuxtLink to="/" class="nav-item">Home</NuxtLink></li>
      <li><NuxtLink to="/blog" class="nav-item">Blog</NuxtLink></li>
      <li><NuxtLink to="/art/gallery" class="nav-item">Art</NuxtLink></li>
      <li><NuxtLink to="/profile" class="nav-item">Profile</NuxtLink></li>
      <li><NuxtLink to="/bluesky" class="nav-item">Bluesky</NuxtLink></li>
    </ul>

    <div class="sidebar-footer">
      <button class="btn btn-primary" @click="$emit('toggle-theme')">Toggle Theme</button>
      <small class="muted">privacy-first</small>
    </div>
  </nav>
</template>

<script setup lang="ts">
// No heavy logic here; this is a presentational navbar/sidebar.
</script>

<style scoped>
@import '~/assets/css/system.css';

/* sidebar base */
.sidebar {
  width: 220px;
  min-height: 100vh;
  background: linear-gradient(180deg, var(--theme-surface), color-mix(in srgb, var(--theme-surface) 90%, black));
  border-right: 1px solid var(--theme-border);
  padding: 1rem;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

/* brand */
.brand {
  font-weight: 700;
  font-size: 1.1rem;
  color: var(--theme-fg);
  display:flex;
  align-items:center;
  gap:.5rem;
}
.brand-link { text-decoration: none; color: inherit; }
.dot { color: var(--theme-accent); margin-left: .1rem; }

/* nav list */
.nav-list { list-style: none; padding: 0; margin: 0; display: flex; flex-direction: column; gap: .25rem; }
.nav-item {
  display:block;
  padding:.5rem .6rem;
  border-radius:6px;
  text-decoration:none;
  color:var(--theme-fg);
  border:1px solid transparent;
}
.nav-item:hover{
  background: color-mix(in srgb, var(--theme-primary) 20%, transparent);
  border-color: var(--theme-border);
}

/* footer */
.sidebar-footer { margin-top: auto; display:flex; flex-direction:column; gap:.5rem; align-items:flex-start; }
.muted { color:var(--theme-muted); font-size:.8rem; }

/* responsive */
@media (max-width: 900px) {
  .sidebar { position: sticky; width:100%; height:auto; min-height:auto; border-right:none; border-bottom:1px solid var(--theme-border); flex-direction:row; gap:.5rem; padding:.5rem; align-items:center; }
  .nav-list { flex-direction:row; gap:.5rem; margin-left: .5rem; }
  .brand { font-size:1rem; }
}
</style>
