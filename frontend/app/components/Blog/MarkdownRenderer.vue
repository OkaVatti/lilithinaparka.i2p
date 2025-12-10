<template>
  <div class="markdown-content" v-html="renderedContent"></div>
</template>

<script setup lang="ts">
import { marked } from 'marked'
import DOMPurify from 'isomorphic-dompurify'

const props = defineProps<{
  content: string
}>()

const renderedContent = computed(() => {
  const html = marked.parse(props.content)
  return DOMPurify.sanitize(html.toString())
})
</script>

<style scoped>
.markdown-content {
  line-height: 1.8;
  color: var(--theme-fg);
}

.markdown-content :deep(h1),
.markdown-content :deep(h2),
.markdown-content :deep(h3),
.markdown-content :deep(h4),
.markdown-content :deep(h5),
.markdown-content :deep(h6) {
  color: var(--theme-primary);
  margin-top: 2rem;
  margin-bottom: 1rem;
  font-weight: bold;
}

.markdown-content :deep(h1) { font-size: 2.5rem; }
.markdown-content :deep(h2) { font-size: 2rem; }
.markdown-content :deep(h3) { font-size: 1.75rem; }
.markdown-content :deep(h4) { font-size: 1.5rem; }
.markdown-content :deep(h5) { font-size: 1.25rem; }
.markdown-content :deep(h6) { font-size: 1.1rem; }

.markdown-content :deep(p) {
  margin-bottom: 1rem;
}

.markdown-content :deep(a) {
  color: var(--theme-accent);
  text-decoration: none;
  border-bottom: 1px solid var(--theme-accent);
  transition: all 0.2s;
}

.markdown-content :deep(a:hover) {
  color: var(--theme-secondary);
  border-color: var(--theme-secondary);
}

.markdown-content :deep(code) {
  background: rgba(0, 0, 0, 0.3);
  padding: 0.2rem 0.4rem;
  border-radius: 3px;
  font-family: 'Courier New', monospace;
  font-size: 0.9em;
}

.markdown-content :deep(pre) {
  background: rgba(0, 0, 0, 0.4);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  padding: 1rem;
  overflow-x: auto;
  margin: 1.5rem 0;
}

.markdown-content :deep(pre code) {
  background: none;
  padding: 0;
}

.markdown-content :deep(blockquote) {
  border-left: 4px solid var(--theme-primary);
  padding-left: 1rem;
  margin: 1.5rem 0;
  color: var(--theme-fg);
  opacity: 0.9;
  font-style: italic;
}

.markdown-content :deep(ul),
.markdown-content :deep(ol) {
  margin: 1rem 0;
  padding-left: 2rem;
}

.markdown-content :deep(li) {
  margin: 0.5rem 0;
}

.markdown-content :deep(img) {
  max-width: 100%;
  height: auto;
  border-radius: 4px;
  margin: 1.5rem 0;
  border: 1px solid var(--theme-border);
}

.markdown-content :deep(table) {
  width: 100%;
  border-collapse: collapse;
  margin: 1.5rem 0;
  overflow-x: auto;
  display: block;
}

.markdown-content :deep(table th),
.markdown-content :deep(table td) {
  border: 1px solid var(--theme-border);
  padding: 0.75rem;
  text-align: left;
}

.markdown-content :deep(table th) {
  background: rgba(189, 147, 249, 0.2);
  color: var(--theme-primary);
  font-weight: bold;
}

.markdown-content :deep(table tr:nth-child(even)) {
  background: rgba(0, 0, 0, 0.2);
}

.markdown-content :deep(hr) {
  border: none;
  border-top: 2px solid var(--theme-border);
  margin: 2rem 0;
}
</style>