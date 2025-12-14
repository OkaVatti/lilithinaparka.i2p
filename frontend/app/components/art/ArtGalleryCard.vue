<template>
  <NuxtLink :to="to" class="gallery-card">
    <img :src="image" :alt="title" />
    <div class="caption">
      <div class="title">{{ title }}</div>
      <div class="meta">{{ author }} • {{ shortDate }}</div>
    </div>
  </NuxtLink>
</template>

<script setup lang="ts">
import { computed } from 'vue'
const props = defineProps<{
  title: string
  image: string
  author?: string
  date?: string
  to?: string
}>()

const shortDate = computed(() => {
  if (!props.date) return ''
  try { return new Date(props.date).toLocaleDateString() } catch { return props.date || '' }
})
</script>

<style scoped>
@import '~/assets/css/system.css';

.gallery-card {
  display:block;
  margin-bottom:12px;
  border-radius:8px;
  overflow:hidden;
  text-decoration:none;
  color:inherit;
  border:1px solid var(--theme-border);
  background:var(--theme-surface);
}
.gallery-card img { width:100%; height:auto; display:block; object-fit:cover; }
.caption { padding:.5rem; }
.title { font-weight:600; color:var(--theme-fg); }
.meta { color:var(--theme-muted); font-size:.85rem; margin-top:.25rem; }
</style>
