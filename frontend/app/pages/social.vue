<template>
  <div class="max-w-4xl mx-auto">
    <div class="mb-8">
      <h1 class="text-3xl font-bold mb-4 font-mono">
        <span class="text-accent">$</span> social
      </h1>
      <p class="text-text-secondary mb-4">
        my recent posts from bluesky
      </p>
      <div class="flex gap-4 flex-wrap">
        <a 
          href="https://bsky.app/profile/lilithinaparka.bsky.social" 
          target="_blank" 
          rel="noopener noreferrer"
          class="inline-block px-4 py-2 border border-link text-link hover:bg-link hover:text-bg-primary transition-colors font-mono"
        >
          [follow me on bluesky]
        </a>
        <button 
          @click="refreshPosts"
          :disabled="loading"
          class="inline-block px-4 py-2 border border-accent text-accent hover:bg-accent hover:text-bg-primary transition-colors font-mono disabled:opacity-50"
        >
          [refresh]
        </button>
      </div>
    </div>
    
    <!-- Loading State -->
    <div v-if="loading" class="terminal-box">
      <p class="text-text-secondary">loading<span class="blink">_</span></p>
    </div>
    
    <!-- Error State -->
    <div v-else-if="error" class="terminal-box">
      <p class="text-error">error: {{ error }}</p>
    </div>
    
    <!-- Empty State -->
    <div v-else-if="posts.length === 0" class="terminal-box">
      <p class="text-text-secondary">no posts found</p>
    </div>
    
    <!-- Posts List -->
    <div v-else>
      <div class="mb-4 text-sm text-text-secondary font-mono">
        showing {{ posts.length }} post{{ posts.length !== 1 ? 's' : '' }}
      </div>
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

const refreshPosts = async () => {
  await bskyStore.fetchPosts(50)
}

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