<!-- app/components/Posts/BlogPostCard.vue -->
<template>
  <NuxtLink :to="`/blog/${post.slug}`" class="blog-post-card block">
    <div class="p-4 rounded-lg border border-border bg-bg-secondary/30 hover:border-accent/30 transition-all duration-300">
      <div class="flex items-start justify-between mb-2">
        <span class="blog-category-tag">{{ post.category }}</span>
        <span class="text-text-secondary text-xs font-mono">{{ formatDate(post.date) }}</span>
      </div>
      <h3 class="text-lg font-bold text-text-primary mb-2 font-mono">{{ post.title }}</h3>
      <p class="text-text-secondary text-sm mb-3">{{ post.summary }}</p>
      <div class="flex flex-wrap gap-1">
        <span v-for="tag in safeTags.slice(0, 3)" :key="tag" class="blog-tag">
          #{{ tag }}
        </span>
        <span v-if="safeTags.length > 3" class="blog-tag">+{{ safeTags.length - 3 }}</span>
      </div>
      <div v-if="post.words" class="mt-3 text-xs text-text-secondary font-mono">
        {{ post.words.toLocaleString() }} words
      </div>
    </div>
  </NuxtLink>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  post: {
    id: number
    title: string
    date: string
    category: string
    tags?: string[]  // Make tags optional
    summary: string
    slug: string
    words?: number
  }
}>()

// Add computed property with fallback
const safeTags = computed(() => props.post.tags || [])

const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
}
</script>