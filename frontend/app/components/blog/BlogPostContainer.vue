<template>
  <article class="post-container">
    <header class="post-header">
      <h2 class="title"><slot name="title" /></h2>
      <div class="meta">
        <time :datetime="date">{{ formattedDate }}</time>
        <span v-if="readingTime">• {{ readingTime }}</span>
      </div>
    </header>

    <section class="post-body">
      <slot />
    </section>

    <footer class="post-footer">
      <slot name="footer" />
    </footer>
  </article>
</template>

<script setup lang="ts">
import { computed } from 'vue'
const props = defineProps({
  date: { type: String, default: () => new Date().toISOString() },
  readingTime: { type: String, default: '' }
})

const formattedDate = computed(() => {
  try {
    return new Date(props.date).toLocaleDateString()
  } catch {
    return props.date
  }
})
</script>

<style scoped>
@import '~/system.css';

.post-container {
  background: var(--theme-surface);
  color: var(--theme-fg);
  border: 1px solid var(--theme-border);
  border-radius: 10px;
  padding: 1rem;
  box-shadow: 0 6px 18px rgba(0,0,0,0.35);
}
.post-header .title {
  margin: 0 0 .4rem 0;
}
.post-header .meta {
  font-size: 0.85rem;
  color: var(--theme-muted);
  margin-bottom: .75rem;
}
.post-body {
  line-height: 1.65;
}
.post-footer {
  margin-top: 1rem;
  border-top: 1px dashed var(--theme-border);
  padding-top: .6rem;
  color: var(--theme-muted);
  display:flex;
  justify-content:space-between;
  gap:.5rem;
}
</style>
