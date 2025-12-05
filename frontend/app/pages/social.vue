<template>
  <div class="max-w-4xl mx-auto">
    <div class="mb-8">
      <h1 class="text-3xl font-bold mb-4">
        <span class="text-accent">$</span> social
      </h1>
      <p class="text-text-secondary mb-4">
        my recent posts from bluesky
      </p>
      <a 
        href="https://bsky.app/profile/lilithinaparka.bsky.social" 
        target="_blank" 
        rel="noopener"
        class="inline-block px-4 py-2 border border-link text-link hover:bg-link hover:text-bg-primary transition-colors"
      >
        follow me on bluesky →
      </a>
    </div>
    
    <div v-if="loading" class="terminal-box">
      <p class="text-text-secondary">loading<span class="blink">_</span></p>
    </div>
    
    <div v-else-if="error" class="terminal-box">
      <p class="text-red-500">error: {{ error }}</p>
    </div>
    
    <div v-else-if="posts.length === 0" class="terminal-box">
      <p class="text-text-secondary">no posts found</p>
    </div>
    
    <div v-else>
      <BskyPostCard 
        v-for="post in posts" 
        :key="post.id" 
        :post="post" 
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useBskyStore } from '../../stores/bsky'

const bskyStore = useBskyStore()

const loading = computed(() => bskyStore.loading)
const error = computed(() => bskyStore.error)
const posts = computed(() => bskyStore.posts)

onMounted(async () => {
  await bskyStore.fetchPosts(50)
})

useHead({
  title: 'Social - Lilith Parker',
  meta: [
    { name: 'description', content: 'my recent posts from bluesky' }
  ]
})
</script>