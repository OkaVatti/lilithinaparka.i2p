<template>
  <div class="profile-card">
    <div class="head">
      <img v-if="avatar" :src="avatar" alt="avatar" class="avatar" />
      <div class="meta">
        <h3 class="name">{{ name }}</h3>
        <p class="bio">{{ bio }}</p>
      </div>
    </div>

    <div class="stats">
      <div v-for="(val, key) in stats" :key="key" class="stat">
        <div class="num">{{ val }}</div>
        <div class="label">{{ key }}</div>
      </div>
    </div>

    <div class="actions">
      <button class="btn">Follow</button>
      <button class="btn">Message</button>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  name: string
  avatar?: string
  bio?: string
  stats?: Record<string, number>
}>()

const stats = props.stats || {}
</script>

<style scoped>
@import '~/assets/css/system.css';

.profile-card {
  border:1px solid var(--theme-border);
  padding:1rem;
  border-radius:10px;
  background: linear-gradient(180deg, transparent, rgba(0,0,0,0.02));
}
.head { display:flex; gap:.8rem; align-items:center; }
.avatar { width:72px; height:72px; border-radius:12px; object-fit:cover; border:1px solid var(--theme-border); }
.name { margin:0; color:var(--theme-fg); }
.bio { color:var(--theme-muted); margin:0.2rem 0 0; }

.stats { display:flex; gap:.6rem; margin-top:.8rem; }
.stat { padding:.4rem .6rem; border-radius:8px; border:1px solid var(--theme-border); text-align:center; min-width:64px; }
.num { font-weight:700; color:var(--theme-fg); }
.label { font-size:.8rem; color:var(--theme-muted); }

.actions { display:flex; gap:.5rem; margin-top:.8rem; }
.btn { padding:.4rem .6rem; border-radius:8px; border:1px solid var(--theme-border); background:var(--theme-surface); color:var(--theme-fg); }
</style>
