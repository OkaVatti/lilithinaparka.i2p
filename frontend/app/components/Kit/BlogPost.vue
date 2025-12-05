<template>
  <article class="blog-post-full">
    <header class="mb-8">
      <h1 class="text-3xl md:text-4xl font-bold mb-4 text-text-primary">
        {{ post.title }}
      </h1>
      
      <div class="flex flex-wrap items-center gap-4 text-sm text-text-secondary mb-4">
        <time :datetime="post.date">{{ formatDate(post.date) }}</time>
        <span v-if="post.time">{{ post.time }}</span>
        <span v-if="authors.length > 0">by {{ authors.join(', ') }}</span>
      </div>
      
      <div class="flex flex-wrap gap-2 mb-4">
        <span 
          v-for="category in categories" 
          :key="category"
          class="blog-category-tag"
        >
          {{ category }}
        </span>
      </div>
      
      <div class="flex flex-wrap gap-2">
        <a 
          v-for="tag in tags" 
          :key="tag"
          :href="`/blog?tag=${encodeURIComponent(tag)}`"
          class="blog-tag hover:text-accent transition-colors"
        >
          #{{ tag }}
        </a>
      </div>
    </header>
    
    <div 
      v-if="post.featured_image" 
      class="mb-8 border border-border overflow-hidden"
    >
      <img 
        :src="post.featured_image" 
        :alt="post.title"
        class="w-full"
      />
    </div>
    
    <div 
      class="prose prose-invert max-w-none font-mono"
      v-html="content"
    ></div>
    
    <footer class="mt-12 pt-8 border-t border-border">
      <div class="flex justify-between items-center flex-wrap gap-4">
        <NuxtLink to="/blog" class="text-link hover:text-accent transition-colors">
          ← back to blog
        </NuxtLink>
        
        <div v-if="post.share" class="flex items-center gap-4 text-sm">
          <span class="text-text-secondary">share this post:</span>
          <button 
            @click="sharePost"
            class="text-link hover:text-accent transition-colors"
          >
            [copy link]
          </button>
        </div>
      </div>
    </footer>
  </article>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import type { BlogPost } from '../../../types'

const props = defineProps<{
  post: BlogPost
  content: string
}>()

const authors = computed(() => {
  try {
    return JSON.parse(props.post.authors)
  } catch {
    return []
  }
})

const categories = computed(() => {
  try {
    return JSON.parse(props.post.categories)
  } catch {
    return []
  }
})

const tags = computed(() => {
  try {
    return JSON.parse(props.post.tags)
  } catch {
    return []
  }
})

const formatDate = (dateStr: string) => {
  const date = new Date(dateStr)
  return date.toLocaleDateString('en-US', { 
    year: 'numeric', 
    month: 'long', 
    day: 'numeric' 
  })
}

const sharePost = async () => {
  const url = window.location.href
  try {
    await navigator.clipboard.writeText(url)
    alert('Link copied to clipboard!')
  } catch (err) {
    console.error('Failed to copy link:', err)
  }
}
</script>