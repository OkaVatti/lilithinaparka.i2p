<!-- app/pages/social.vue -->
<template>
  <div class="social-page">
    <!-- Social Header -->
    <div class="social-header mb-8">
      <h1 class="text-4xl font-bold text-accent mb-4 font-mono">$ social</h1>
      <p class="text-text-secondary font-mono">
        > my social media presence and microblogging
      </p>
    </div>
    
    <!-- Social Links -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12">
      <SocialLink
        v-for="link in socialLinks"
        :key="link.platform"
        :link="link"
      />
    </div>
    
    <!-- Bluesky Timeline -->
    <div class="bluesky-timeline">
      <div class="flex items-center justify-between mb-6">
        <h2 class="text-2xl font-bold font-mono text-accent">
          <span class="text-sky-500">𝕏</span> bluesky timeline
        </h2>
        <div class="flex items-center gap-4">
          <span class="text-text-secondary text-sm font-mono">
            @lilithinaparka.bsky.social
          </span>
          <button @click="refreshPosts" class="refresh-button">
            <svg class="w-4 h-4" :class="{ 'animate-spin': loading }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
          </button>
        </div>
      </div>
      
      <!-- Bluesky Posts -->
      <div class="space-y-4">
        <BskyPostCard
          v-for="post in bskyPosts"
          :key="post.id"
          :post="post"
          :detailed="true"
        />
      </div>
      
      <!-- Load More -->
      <div v-if="hasMorePosts" class="text-center mt-8">
        <button @click="loadMorePosts" class="load-more-button">
          load more posts
        </button>
      </div>
    </div>
    
    <!-- Social Statistics -->
    <div class="mt-12 p-6 rounded-xl bg-bg-secondary/50 border border-accent/10">
      <h3 class="text-xl font-bold mb-4 font-mono text-accent">$ social stats</h3>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
        <div class="social-stat">
          <div class="stat-number">{{ totalPosts.toLocaleString() }}</div>
          <div class="stat-label">Total Posts</div>
        </div>
        <div class="social-stat">
          <div class="stat-number">{{ followers.toLocaleString() }}</div>
          <div class="stat-label">Followers</div>
        </div>
        <div class="social-stat">
          <div class="stat-number">{{ following.toLocaleString() }}</div>
          <div class="stat-label">Following</div>
        </div>
        <div class="social-stat">
          <div class="stat-number">2022</div>
          <div class="stat-label">Joined</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import SocialLink from '~/components/Posts/SocialLink.vue'
import BskyPostCard from '~/components/Posts/BskyPostCard.vue'
import { ref, computed } from 'vue'

definePageMeta({
  layout: 'bsky'
})

const socialLinks = [
  { platform: 'Bluesky', username: '@lilithinaparka.bsky.social', url: 'https://bsky.app/profile/lilithinaparka.bsky.social', icon: '𝕏', color: 'text-sky-500', description: 'Primary microblogging account' },
  { platform: 'Bluesky', username: '@aVaOk.bsky.social', url: 'https://bsky.app/profile/aVaOk.bsky.social', icon: '𝕏', color: 'text-sky-500', description: 'Tech/programming focused' },
  { platform: 'GitHub', username: '@okavatti', url: 'https://github.com/okavatti', icon: '🐙', color: 'text-gray-300', description: 'Open source projects' },
  { platform: 'Gitten', username: 'lilithinaparka.gitten.i2p', url: 'http://lilithinaparka.gitten.i2p/profile', icon: '🐈', color: 'text-accent', description: 'I2P-based Git hosting' },
  { platform: 'Bottletail', username: '#lilithinaparka.bttl.dev', url: 'https://bttl.dev/#lilithinaparka', icon: '🍶', color: 'text-purple-500', description: 'Decentralized social media' },
  { platform: 'Email', username: 'lilithinaparka@mail.i2p', url: 'mailto:lilithinaparka@mail.i2p', icon: '✉️', color: 'text-red-400', description: 'Primary contact email' }
]

// Sample Bluesky posts
const bskyPosts = ref([
  { 
    id: 1, 
    text: 'Just finished implementing end-to-end encryption for my chat app. Privacy matters! #privacy #encryption', 
    author: '@lilithinaparka.bsky.social', 
    likes: 42, 
    reposts: 12, 
    replies: 8,
    timestamp: '2 hours ago',
    hasMedia: false
  },
  { 
    id: 2, 
    text: 'New art piece: "Digital Ghost in the Machine". Exploring the intersection of AI and consciousness. The idea that AI might develop some form of awareness is fascinating and terrifying at the same time.', 
    author: '@lilithinaparka.bsky.social', 
    likes: 89, 
    reposts: 23, 
    replies: 15,
    timestamp: '1 day ago',
    hasMedia: true,
    media: ['/images/art/digital-ghost.jpg']
  },
  { 
    id: 3, 
    text: 'Working on a new open-source shell written in Crystal. Going to call it CRSH. Features include: syntax highlighting, autocomplete, plugin system, and built-in package manager. Stay tuned! #programming #crystal', 
    author: '@aVaOk.bsky.social', 
    likes: 56, 
    reposts: 18, 
    replies: 12,
    timestamp: '2 days ago',
    hasMedia: false
  },
  { 
    id: 4, 
    text: 'Why do we still use centralized social media when decentralized alternatives exist? The answer: network effects and convenience. But at what cost to our privacy and freedom? #decentralization #privacy', 
    author: '@lilithinaparka.bsky.social', 
    likes: 67, 
    reposts: 31, 
    replies: 24,
    timestamp: '3 days ago',
    hasMedia: false
  },
  { 
    id: 5, 
    text: 'Just released chaos.cr v1.2.0 - now with support for hyperchaotic attractors in hashing algorithms. The security implications are interesting: deterministic yet unpredictable. #cryptography #chaos', 
    author: '@aVaOk.bsky.social', 
    likes: 34, 
    reposts: 9, 
    replies: 7,
    timestamp: '4 days ago',
    hasMedia: false
  }
])

const loading = ref(false)
const hasMorePosts = ref(true)

const totalPosts = computed(() => 1234)
const followers = computed(() => 1500)
const following = computed(() => 450)

const refreshPosts = async () => {
  loading.value = true
  // Simulate API call
  await new Promise(resolve => setTimeout(resolve, 1000))
  loading.value = false
}

const loadMorePosts = () => {
  // Simulate loading more posts
  const newPost = {
    id: bskyPosts.value.length + 1,
    text: 'This is an older post that was loaded. The pagination system works!',
    author: '@lilithinaparka.bsky.social',
    likes: 10,
    reposts: 2,
    replies: 1,
    timestamp: '1 week ago',
    hasMedia: false
  }
  bskyPosts.value.push(newPost)
}
</script>

<style scoped>
.social-page {
  @apply max-w-3xl mx-auto;
}

.social-header {
  @apply pb-8 border-b border-accent/20;
}

.bluesky-timeline {
  @apply p-6 rounded-xl bg-bg-secondary/30 border border-accent/10;
}

.refresh-button {
  @apply p-2 rounded-lg border border-accent/20 text-text-secondary 
         hover:text-accent hover:border-accent/40 transition-all duration-200;
}

.load-more-button {
  @apply px-6 py-2 rounded-full border border-accent/20 text-sm font-mono
         text-text-secondary hover:text-accent hover:border-accent/40
         transition-all duration-200;
}

.social-stat {
  @apply p-4 rounded-lg bg-bg-secondary/30 border border-accent/10 text-center;
}

.social-stat .stat-number {
  @apply text-2xl font-bold text-accent font-mono;
}

.social-stat .stat-label {
  @apply text-text-secondary text-sm mt-1;
}
</style>