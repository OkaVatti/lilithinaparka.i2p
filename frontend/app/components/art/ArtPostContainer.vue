<template>
  <article class="art-post">
    <div class="media">
      <img :src="image" :alt="title" />
    </div>

    <div class="content">
      <h1 class="title">{{ title }}</h1>
      <p class="meta">by <strong>{{ author }}</strong> — <time>{{ formattedDate }}</time></p>
      <div class="desc" v-html="description"></div>

      <div v-if="tags?.length" class="tags">
        <span class="tag" v-for="t in tags" :key="t">{{ t }}</span>
      </div>
    </div>
  </article>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  title: string
  image: string
  description?: string
  author?: string
  date?: string
  tags?: string[]
}>()

const formattedDate = computed(() => {
  try { return props.date ? new Date(props.date).toLocaleDateString() : '' } catch { return props.date || '' }
})
</script>

<style scoped>
@import '~/assets/css/system.css';

.art-post { display:flex; flex-direction:column; gap:1rem; border:1px solid var(--theme-border); padding:1rem; border-radius:10px; background:var(--theme-surface); }
.media { width:100%; height:520px; overflow:hidden; border-radius:8px; border:1px solid var(--theme-border); display:flex; align-items:center; justify-content:center; background:var(--theme-bg); }
.media img { width:100%; height:100%; object-fit:cover; display:block; }

.title { margin:0; color:var(--theme-fg); font-size:1.6rem; }
.meta { color:var(--theme-muted); margin:.3rem 0; }
.desc { color:var(--theme-muted); line-height:1.5; margin-top:.6rem; }

.tags { margin-top: .8rem; display:flex; gap:.4rem; flex-wrap:wrap; }
.tag { padding:.2rem .5rem; border-radius:6px; border:1px solid var(--theme-border); color:var(--theme-muted); }
</style>
