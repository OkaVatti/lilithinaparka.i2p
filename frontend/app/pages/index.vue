<template>
  <div class="max-width-container">
    <!-- ASCII Art Banner -->
    <section class="terminal-box mb-8">
      <pre class="ascii-art text-accent text-xs sm:text-sm md:text-base">
 ___       ___  ___       ___  _________  ___  ___     
|\  \     |\  \|\  \     |\  \|\___   ___\\  \|\  \    
\ \  \    \ \  \ \  \    \ \  \|___ \  \_\ \  \\\  \   
 \ \  \    \ \  \ \  \    \ \  \   \ \  \ \ \   __  \  
  \ \  \____\ \  \ \  \____\ \  \   \ \  \ \ \  \ \  \ 
   \ \_______\ \__\ \_______\ \__\   \ \__\ \ \__\ \__\
    \|_______|\|__|\|_______|\|__|    \|__|  \|__|\|__|
      </pre>
    </section>

    <!-- Welcome Section -->
    <section class="mb-12">
      <h1 class="text-2xl font-bold mb-4 font-mono">
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
        <div class="mt-4 text-xs text-text-secondary">
          <span class="text-accent">&gt;&gt;</span> last updated: {{ currentDate }}
        </div>
      </div>
    </section>

    <!-- Recent Posts -->
    <section class="mb-12">
      <h2 class="text-xl font-bold mb-4 font-mono">
        <span class="text-accent">$</span> recent_posts
      </h2>
      <div v-if="loading" class="terminal-box">
        <p class="text-text-secondary">
          loading<span class="blink">_</span>
        </p>
      </div>
      <div v-else-if="error" class="terminal-box">
        <p class="text-error">error: {{ error }}</p>
      </div>
      <div v-else>
        <BlogPostCard 
          v-for="post in recentPosts" 
          :key="post.id" 
          :post="post" 
        />
        <NuxtLink 
          to="/blog" 
          class="inline-block mt-4 px-4 py-2 border border-accent text-accent hover:bg-accent hover:text-bg-primary transition-colors font-mono"
        >
          <span class="text-accent">[</span> view all posts <span class="text-accent">]</span>
        </NuxtLink>
      </div>
    </section>

    <!-- Latest Thoughts (BlueSky) -->
    <section>
      <h2 class="text-xl font-bold mb-4 font-mono">
        <span class="text-accent">$</span> latest_thoughts
      </h2>
      <div v-if="bskyLoading" class="terminal-box">
        <p class="text-text-secondary">
          loading<span class="blink">_</span>
        </p>
      </div>
      <div v-else-if="bskyError" class="terminal-box">
        <p class="text-error">error: {{ bskyError }}</p>
      </div>
      <div v-else>
        <BskyPostCard 
          v-for="post in latestPosts" 
          :key="post.id" 
          :post="post" 
        />
        <NuxtLink 
          to="/social" 
          class="inline-block mt-4 px-4 py-2 border border-accent text-accent hover:bg-accent hover:text-bg-primary transition-colors font-mono"
        >
          <span class="text-accent">[</span> view all posts <span class="text-accent">]</span>
        </NuxtLink>
      </div>
    </section>

    <!-- Quick Stats -->
    <section class="mt-12 terminal-box">
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4 text-center text-sm">
        <div>
          <div class="text-2xl text-accent font-bold">{{ publishedPostsCount }}</div>
          <div class="text-text-secondary">blog posts</div>
        </div>
        <div>
          <div class="text-2xl text-accent font-bold">{{ bskyPostsCount }}</div>
          <div class="text-text-secondary">social posts</div>
        </div>
        <div>
          <div class="text-2xl text-accent font-bold">{{ categoriesCount }}</div>
          <div class="text-text-secondary">categories</div>
        </div>
        <div>
          <div class="text-2xl text-accent font-bold">∞</div>
          <div class="text-text-secondary">caffeine consumed</div>
        </div>
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

const publishedPostsCount = computed(() => blogStore.publishedPosts.length)
const bskyPostsCount = computed(() => bskyStore.posts.length)
const categoriesCount = computed(() => 3)

const currentDate = computed(() => {
  const date = new Date()
  return date.toLocaleDateString('en-US', { 
    year: 'numeric', 
    month: 'long', 
    day: 'numeric' 
  })
})

onMounted(async () => {
  await blogStore.fetchPosts()
  await bskyStore.fetchPosts(10)
})

useHead({
  title: 'Lilith Parker - Home',
  meta: [
    { 
      name: 'description', 
      content: 'chronically online cat-girl who likes programming, music, and burritos' 
    }
  ]
})
</script>

<style scoped>
.ascii-art {
  font-family: 'Courier New', monospace;
  line-height: 1.2;
  margin: 0;
  overflow-x: auto;
}

.max-width-container {
  max-width: 900px;
  margin: 0 auto;
}
</style>