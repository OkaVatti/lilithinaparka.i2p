<template>
  <article class="search-result" @click="$emit('click', result)">
    <div class="result-type" :class="result.type">
      <FeatherIcon :name="typeIcon" size="16" />
      <span>{{ typeLabel }}</span>
    </div>
    
    <div class="result-content">
      <h4 class="result-title">{{ result.title }}</h4>
      
      <p class="result-description">{{ result.description }}</p>
      
      <div v-if="result.highlights && result.highlights.length > 0" class="result-highlights">
        <span class="highlight-label">Matches:</span>
        <div class="highlight-tags">
          <span v-for="highlight in result.highlights.slice(0, 3)" :key="highlight" class="highlight-tag">
            {{ highlight }}
          </span>
        </div>
      </div>
      
      <div class="result-meta">
        <span class="result-date">{{ formatDate(result.date) }}</span>
        <span class="result-score">Relevance: {{ Math.round(result.score * 100) }}%</span>
      </div>
    </div>
    
    <div class="result-action">
      <FeatherIcon name="chevron-right" size="20" />
    </div>
  </article>
</template>

<script setup lang="ts">
import type { SearchResult } from '~~/types'

defineProps<{
  result: SearchResult
}>()

defineEmits<{
  click: [result: SearchResult]
}>()

const typeIcon = computed(() => {
  const icons = {
    blog: 'file-text',
    game: 'play',
    media: 'image',
    page: 'file'
  }
  return icons[result.type as keyof typeof icons] || 'file'
})

const typeLabel = computed(() => {
  const labels = {
    blog: 'Blog Post',
    game: 'Game',
    media: 'Media',
    page: 'Page'
  }
  return labels[result.type as keyof typeof labels] || result.type
})

const formatDate = (dateStr: string | Date) => {
  const date = new Date(dateStr)
  return date.toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}
</script>

<style scoped>
.search-result {
  display: flex;
  align-items: flex-start;
  gap: 1rem;
  padding: 1.5rem;
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--theme-border);
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s;
}

.search-result:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
  transform: translateX(4px);
}

.result-type {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.25rem;
  padding: 0.5rem;
  min-width: 60px;
  border-radius: 4px;
  font-size: 0.8rem;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.result-type.blog {
  background: rgba(139, 147, 233, 0.2);
  color: #8b8be9;
  border: 1px solid #8b8be9;
}

.result-type.game {
  background: rgba(189, 147, 249, 0.2);
  color: #bd93f9;
  border: 1px solid #bd93f9;
}

.result-type.media {
  background: rgba(80, 250, 123, 0.2);
  color: #50fa7b;
  border: 1px solid #50fa7b;
}

.result-type.page {
  background: rgba(255, 184, 108, 0.2);
  color: #ffb86c;
  border: 1px solid #ffb86c;
}

.result-content {
  flex: 1;
}

.result-title {
  margin: 0 0 0.5rem 0;
  color: var(--theme-primary);
  font-size: 1.2rem;
}

.result-description {
  margin: 0 0 1rem 0;
  color: var(--theme-fg);
  opacity: 0.8;
  line-height: 1.5;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.result-highlights {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

.highlight-label {
  font-size: 0.9rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

.highlight-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 0.25rem;
}

.highlight-tag {
  padding: 0.25rem 0.5rem;
  background: rgba(189, 147, 249, 0.1);
  border: 1px solid var(--theme-primary);
  border-radius: 4px;
  font-size: 0.8rem;
  color: var(--theme-primary);
}

.result-meta {
  display: flex;
  gap: 1rem;
  font-size: 0.9rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

.result-action {
  display: flex;
  align-items: center;
  padding: 0.5rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

@media (max-width: 768px) {
  .search-result {
    flex-direction: column;
    gap: 1rem;
  }
  
  .result-type {
    flex-direction: row;
    align-self: flex-start;
    min-width: auto;
  }
  
  .result-meta {
    flex-direction: column;
    gap: 0.25rem;
  }
}
</style>