<template>
  <article class="blog-post-card card">
    <NuxtLink :to="`/blog/${post.slug}`" class="card-link">
      <div v-if="post.featured_image" class="card-image">
        <img :src="post.featured_image" :alt="post.title" loading="lazy" />
      </div>
      
      <div class="card-content">
        <div class="card-meta">
          <span class="date">
            <FeatherIcon name="calendar" size="14" />
            {{ formatDate(post.date) }}
          </span>
          <span v-if="categories.length" class="categories">
            <FeatherIcon name="folder" size="14" />
            {{ categories.join(', ') }}
          </span>
        </div>
        
        <h2 class="card-title">{{ post.title }}</h2>
        
        <p v-if="post.summary" class="card-summary">{{ post.summary }}</p>
        
        <div v-if="tags.length" class="card-tags">
          <span v-for="tag in tags" :key="tag" class="tag">
            <FeatherIcon name="tag" size="12" />
            {{ tag }}
          </span>
        </div>
        
        <div v-if="authors.length" class="card-authors">
          <FeatherIcon name="user" size="14" />
          <span>{{ authors.join(', ') }}</span>
        </div>
      </div>
    </NuxtLink>
  </article>
</template>

<script setup lang="ts">
import type { BlogPost } from '~~/types'

const props = defineProps<{
  post: BlogPost
}>()

const categories = computed(() => {
  try {
    return JSON.parse(props.post.categories || '[]')
  } catch {
    return []
  }
})

const tags = computed(() => {
  try {
    return JSON.parse(props.post.tags || '[]')
  } catch {
    return []
  }
})

const authors = computed(() => {
  try {
    return JSON.parse(props.post.authors || '[]')
  } catch {
    return []
  }
})

const formatDate = (dateStr: string) => {
  try {
    return new Date(dateStr).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    })
  } catch {
    return dateStr
  }
}
</script>

<style scoped>
.blog-post-card {
  transition: all 0.3s;
}

.card-link {
  text-decoration: none;
  color: inherit;
  display: block;
}

.card-image {
  width: 100%;
  height: 200px;
  overflow: hidden;
  border-radius: 4px;
  margin-bottom: 1rem;
  background: rgba(0, 0, 0, 0.2);
}

.card-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s;
}

.blog-post-card:hover .card-image img {
  transform: scale(1.05);
}

.card-content {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.card-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
  font-size: 0.9rem;
  color: var(--theme-fg);
  opacity: 0.7;
}

.card-meta span {
  display: flex;
  align-items: center;
  gap: 0.25rem;
}

.card-title {
  margin: 0;
  font-size: 1.5rem;
  color: var(--theme-primary);
  transition: color 0.2s;
}

.blog-post-card:hover .card-title {
  color: var(--theme-secondary);
}

.card-summary {
  margin: 0;
  color: var(--theme-fg);
  opacity: 0.9;
  line-height: 1.6;
}

.card-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.tag {
  display: flex;
  align-items: center;
  gap: 0.25rem;
  padding: 0.25rem 0.5rem;
  background: rgba(189, 147, 249, 0.1);
  border: 1px solid var(--theme-primary);
  border-radius: 4px;
  font-size: 0.85rem;
  color: var(--theme-primary);
}

.card-authors {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.9rem;
  color: var(--theme-accent);
  margin-top: 0.5rem;
  padding-top: 0.75rem;
  border-top: 1px solid var(--theme-border);
}
</style>