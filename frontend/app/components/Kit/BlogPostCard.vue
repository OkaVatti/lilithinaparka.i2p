<template>
  <article class="terminal-box mb-6 hover:border-accent transition-colors">
    <div class="flex flex-col">
      <div class="flex justify-between items-start mb-2">
        <NuxtLink :to="`/blog/${post.slug}`" class="text-lg font-bold hover:text-accent">
          {{ post.title }}
        </NuxtLink>
        <span class="text-xs text-text-secondary whitespace-nowrap ml-4">
          {{ post.date }}
        </span>
      </div>
      
      <div class="flex flex-wrap gap-2 mb-3">
        <span 
          v-for="category in categories" 
          :key="category"
          class="text-xs px-2 py-1 bg-bg-primary text-accent border border-accent"
        >
          {{ category }}
        </span>
      </div>
      
      <p class="text-sm text-text-secondary mb-3">
        {{ post.summary }}
      </p>
      
      <div class="flex flex-wrap gap-2">
        <span 
          v-for="tag in tags" 
          :key="tag"
          class="text-xs text-link"
        >
          #{{ tag }}
        </span>
      </div>
      
      <div class="mt-4 text-xs text-text-secondary">
        <NuxtLink :to="`/blog/${post.slug}`" class="hover:text-accent">
          read more →
        </NuxtLink>
      </div>
    </div>
  </article>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { BlogPost } from '../../../types'

const props = defineProps<{
  post: BlogPost
}>()

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
</script>