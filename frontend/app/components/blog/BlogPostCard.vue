<template>
  <article class="card">
    <NuxtLink :to="to" class="img-wrap" v-if="image">
      <img :src="image" :alt="title" />
    </NuxtLink>

    <div class="meta">
      <div class="title-row">
        <h3 class="card-title"><NuxtLink :to="to">{{ title }}</NuxtLink></h3>
        <time class="date">{{ formattedDate }}</time>
      </div>
      <p class="excerpt">{{ excerpt }}</p>
      <div class="tags">
        <span class="tag" v-for="t in tags" :key="t">{{ t }}</span>
      </div>
    </div>
  </article>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  title: string
  excerpt?: string
  date?: string
  tags?: string[]
  to?: string
  image?: string
}>()

const formattedDate = computed(() => {
  try {
    if (!props.date) return ''
    const d = new Date(props.date)
    return d.toLocaleDateString()
  } catch {
    return props.date || ''
  }
})

const tags = props.tags || []
</script>

<style scoped>
@import '~/assets/css/system.css';

.card {
  background: var(--theme-bg);
  border: 1px solid var(--theme-border);
  border-radius: 8px;
  overflow: hidden;
  display:flex;
  flex-direction:column;
  transition: transform .15s ease, box-shadow .15s ease;
}
.card:hover { transform: translateY(-4px); box-shadow: 0 10px 30px rgba(0,0,0,0.35); }

.img-wrap { display:block; width:100%; height:180px; overflow:hidden; background:var(--theme-surface); }
.img-wrap img { width:100%; height:100%; object-fit:cover; display:block; }

.meta { padding: .8rem; display:flex; flex-direction:column; gap:.5rem; }
.title-row { display:flex; justify-content:space-between; align-items:center; gap:.5rem; }
.card-title { margin:0; font-size:1.05rem; color:var(--theme-fg); }
.date { color:var(--theme-muted); font-size:.85rem; }

.excerpt { color:var(--theme-muted); margin:0; font-size:.95rem; }
.tags { margin-top:.5rem; display:flex; gap:.4rem; flex-wrap:wrap; }
.tag { background: transparent; border:1px solid var(--theme-border); padding: .18rem .5rem; border-radius:6px; color:var(--theme-muted); font-size:.8rem; }
</style>
