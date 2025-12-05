<template>
  <div class="max-w-4xl mx-auto">
    <div class="terminal-box mb-8">
      <pre class="ascii-art text-accent">
 _     _ _ _ _   _       ____             _             
| |   (_) (_) |_| |__   |  _ \ __ _ _ __| | _____ _ __ 
| |   | | | | __| '_ \  | |_) / _` | '__| |/ / _ \ '__|
| |___| | | | |_| | | | |  __/ (_| | |  |   <  __/ |   
|_____|_|_|_|\__|_| |_| |_|   \__,_|_|  |_|\_\___|_|   
      </pre>
    </div>
    
    <section class="mb-12">
      <h1 class="text-2xl font-bold mb-4">
        <span class="text-accent">$</span> whoami
      </h1>
      <div class="terminal-box">
        <p class="mb-4">
          chronically online cat-girl who likes programming, music, and burritos
        </p>
        <p class="text-sm text-text-secondary">
          welcome to my corner of the I2P network. here you'll find my thoughts on code, 
          technology, and whatever else crosses my mind.
        </p>
      </div>
    </section>
    
    <section class="mb-12">
      <h2 class="text-xl font-bold mb-4">
        <span class="text-accent">$</span> recent_posts
      </h2>
      <div v-if="loading" class="terminal-box">
        <p class="text-text-secondary">loading<span class="blink">_</span></p>
      </div>
      <div v-else-if="error" class="terminal-box">
        <p class="text-red-500">error: {{ error }}</p>
      </div>
      <div v-else>
        <BlogPostCard 
          v-for="post in recentPosts" 
          :key="post.id" 
          :post="post" 
        />
        <NuxtLink 
          to="/blog" 
          class="inline-block mt-4 px-4 py-2 border border-accent text-accent hover:bg-accent hover:text-bg-primary transition-colors"
        >
          view all posts →
        </NuxtLink>
      </div>
    </section>
    
    <section>
      <h2 class="text-xl font-bold mb-4">
        <span class="text-accent">$</span> latest_thoughts
      </h2>
      <div v-if="bskyLoading" class="terminal-box">
        <p class="text-text-secondary">loading<span class="blink">_</span></p>
      </div>
      <div v-else-if="bskyError" class="terminal-box">
        <p class="text-red-500">error: {{ bskyError }}</p>
      </div>
      <div v-else>
        <BskyPostCard 
          v-for="post in latestPosts" 
          :key="post.id" 
          :post="post" 
        />
        <NuxtLink 
          to="/social" 
          class="inline-block mt-4 px-4 py-2 border border-accent text-accent hover:bg-accent hover:text-bg-primary transition-colors"
        >
          view all posts →
        </NuxtLink>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { useBlogStore } from '../../stores/blog'
import { useBskyStore } from '../../stores/bsky'

const blogStore = useBlogStore()
const bskyStore = useBskyStore()

const loading = computed(() => blogStore.loading)
const error = computed(() => blogStore.error)
const bskyLoading = computed(() => bskyStore.loading)
const bskyError = computed(() => bskyStore.error)

const recentPosts = computed(() => 
  blogStore.publishedPosts.slice(0, 3)
)

const latestPosts = computed(() => 
  bskyStore.posts.slice(0, 5)
)

onMounted(async () => {
  await blogStore.fetchPosts()
  await bskyStore.fetchPosts(10)
})

useHead({
  title: 'Lilith Parker - Home',
  meta: [
    { name: 'description', content: 'chronically online cat-girl who likes programming, music, and burritos' }
  ]
})
</script>