<template>
  <div class="max-w-4xl mx-auto">
    <div class="mb-8">
      <h1 class="text-3xl font-bold mb-4">
        <span class="text-accent">$</span> blog
      </h1>
      <p class="text-text-secondary">
        thoughts, tutorials, and random musings
      </p>
    </div>
    
    <div class="mb-6 flex flex-wrap gap-4">
      <button 
        @click="filterCategory = null"
        :class="[
          'px-3 py-1 border text-sm transition-colors',
          filterCategory === null 
            ? 'border-accent bg-accent text-bg-primary' 
            : 'border-border hover:border-accent'
        ]"
      >
        all
      </button>
      <button 
        v-for="cat in categories" 
        :key="cat"
        @click="filterCategory = cat"
        :class="[
          'px-3 py-1 border text-sm transition-colors',
          filterCategory === cat 
            ? 'border-accent bg-accent text-bg-primary' 
            : 'border-border hover:border-accent'
        ]"
      >
        {{ cat.toLowerCase() }}
      </button>
    </div>
    
    <div v-if="loading" class="terminal-box">
      <p class="text-text-secondary">loading<span class="blink">_</span></p>
    </div>
    
    <div v-else-if="error" class="terminal-box">
      <p class="text-red-500">error: {{ error }}</p>
    </div>
    
    <div v-else-if="filteredPosts.length === 0" class="terminal-box">
      <p class="text-text-secondary">no posts found</p>
    </div>
    
    <div v-else>
      <BlogPostCard 
        v-for="post in filteredPosts" 
        :key="post.id" 
        :post="post" 
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useBlogStore } from '../../../stores/blog'

const blogStore = useBlogStore()
const filterCategory = ref<string | null>(null)

const loading = computed(() => blogStore.loading)
const error = computed(() => blogStore.error)

const categories = ['Casual', 'Interlude', 'Serious']

const filteredPosts = computed(() => {
  if (!filterCategory.value) {
    return blogStore.publishedPosts
  }
  return blogStore.postsByCategory(filterCategory.value)
})

onMounted(async () => {
  await blogStore.fetchPosts()
})

useHead({
  title: 'Blog - Lilith Parker',
  meta: [
    { name: 'description', content: 'thoughts, tutorials, and random musings' }
  ]
})
</script>