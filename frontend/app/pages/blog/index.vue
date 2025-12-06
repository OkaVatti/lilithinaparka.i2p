<!-- app/pages/blog/index.vue -->
<template>
  <div class="blog-page">
    <!-- Blog Header -->
    <div class="blog-header mb-8">
      <h1 class="text-4xl font-bold text-accent mb-4 font-mono">$ blog</h1>
      <p class="text-text-secondary font-mono">
        > thoughts, tutorials, stories, and everything in between
      </p>
      
      <!-- Category Filter -->
      <div class="flex flex-wrap gap-2 mt-6">
        <button
          v-for="category in categories"
          :key="category"
          @click="setActiveCategory(category)"
          class="category-filter"
          :class="{ 'active': activeCategory === category }"
        >
          {{ category }}
        </button>
      </div>
    </div>
    
    <!-- Blog Posts Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12">
      <BlogPostCard
        v-for="post in filteredPosts"
        :key="post.id"
        :post="post"
        :detailed="true"
      />
    </div>
    
    <!-- Blog Categories Explanation -->
    <div class="blog-categories-explanation mt-12 p-6 rounded-xl bg-bg-secondary/50 border border-accent/10">
      <h3 class="text-xl font-bold mb-4 font-mono text-accent">$ blog categories</h3>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="category-card">
          <h4 class="text-lg font-bold mb-2 text-accent font-mono">Casual</h4>
          <p class="text-text-secondary text-sm">
            Lighthearted posts, memes, jokes, and short stories. Perfect for casual reading.
          </p>
          <ul class="mt-2 space-y-1 text-xs text-text-secondary font-mono">
            <li>• memes / jokes</li>
            <li>• short stories</li>
            <li>• daily life</li>
          </ul>
        </div>
        <div class="category-card">
          <h4 class="text-lg font-bold mb-2 text-accent font-mono">Interlude</h4>
          <p class="text-text-secondary text-sm">
            Intermediate posts about technology, interesting finds, tutorials, and vlogs.
          </p>
          <ul class="mt-2 space-y-1 text-xs text-text-secondary font-mono">
            <li>• technology</li>
            <li>• programming</li>
            <li>• tutorials</li>
            <li>• vlog content</li>
          </ul>
        </div>
        <div class="category-card">
          <h4 class="text-lg font-bold mb-2 text-accent font-mono">Serious</h4>
          <p class="text-text-secondary text-sm">
            Serious discussions about philosophy, sociology, politics, and deep topics.
          </p>
          <ul class="mt-2 space-y-1 text-xs text-text-secondary font-mono">
            <li>• philosophy</li>
            <li>• sociology</li>
            <li>• politics</li>
            <li>• ethics</li>
          </ul>
        </div>
      </div>
    </div>
    
    <!-- Blog Stats -->
    <div class="mt-8 grid grid-cols-2 md:grid-cols-4 gap-4">
      <div class="blog-stat">
        <div class="stat-number">{{ totalPosts }}</div>
        <div class="stat-label">Total Posts</div>
      </div>
      <div class="blog-stat">
        <div class="stat-number">{{ categories.length }}</div>
        <div class="stat-label">Categories</div>
      </div>
      <div class="blog-stat">
        <div class="stat-number">{{ totalWords.toLocaleString() }}+</div>
        <div class="stat-label">Words Written</div>
      </div>
      <div class="blog-stat">
        <div class="stat-number">2019</div>
        <div class="stat-label">Since</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import BlogPostCard from '../../components/Posts/BlogPostCard.vue'
import { ref, computed } from 'vue'

definePageMeta({
  layout: 'blog'
})

// Sample blog data
const blogPosts = [
  { id: 1, title: 'Building a decentralized social network', date: '2024-01-15', category: 'Interlude', tags: ['programming', 'decentralization', 'i2p'], summary: 'How I built my own social network using Crystal and I2P', slug: 'decentralized-social-network', words: 2450 },
  { id: 2, title: 'The ethics of AI art generation', date: '2024-01-10', category: 'Serious', tags: ['ai', 'ethics', 'art'], summary: 'Exploring the moral implications of AI-generated artwork', slug: 'ai-art-ethics', words: 3200 },
  { id: 3, title: 'My favorite burrito recipe', date: '2024-01-05', category: 'Casual', tags: ['food', 'recipe', 'casual'], summary: 'A step-by-step guide to making the perfect breakfast burrito', slug: 'burrito-recipe', words: 1200 },
  { id: 4, title: 'Crystal vs Ruby: Performance Comparison', date: '2024-01-02', category: 'Interlude', tags: ['programming', 'crystal', 'ruby', 'benchmark'], summary: 'Comparing the performance of Crystal and Ruby for web applications', slug: 'crystal-vs-ruby', words: 1800 },
  { id: 5, title: 'The philosophy of free software', date: '2023-12-28', category: 'Serious', tags: ['philosophy', 'open-source', 'ethics'], summary: 'Examining the ethical foundations of the free software movement', slug: 'free-software-philosophy', words: 2800 },
  { id: 6, title: 'ASCII art animation tutorial', date: '2023-12-20', category: 'Interlude', tags: ['tutorial', 'ascii', 'animation'], summary: 'Learn how to create animated ASCII art in the terminal', slug: 'ascii-art-tutorial', words: 1500 },
  { id: 7, title: 'Why I use I2P instead of Tor', date: '2023-12-15', category: 'Interlude', tags: ['privacy', 'i2p', 'tor', 'security'], summary: 'My reasons for preferring I2P over Tor for daily use', slug: 'i2p-vs-tor', words: 2200 },
  { id: 8, title: 'Funny programming bugs collection', date: '2023-12-10', category: 'Casual', tags: ['programming', 'humor', 'bugs'], summary: 'A collection of hilarious and weird bugs I\'ve encountered', slug: 'funny-programming-bugs', words: 900 },
  { id: 9, title: 'The social impact of decentralization', date: '2023-12-05', category: 'Serious', tags: ['sociology', 'decentralization', 'society'], summary: 'How decentralized technologies are changing social structures', slug: 'decentralization-social-impact', words: 3500 }
]

const categories = ['All', 'Casual', 'Interlude', 'Serious']
const activeCategory = ref('All')

const filteredPosts = computed(() => {
  if (activeCategory.value === 'All') return blogPosts
  return blogPosts.filter(post => post.category === activeCategory.value)
})

const totalPosts = computed(() => blogPosts.length)
const totalWords = computed(() => blogPosts.reduce((sum, post) => sum + post.words, 0))

const setActiveCategory = (category: string) => {
  activeCategory.value = category
}
</script>

<style scoped>
.blog-page {
  @apply max-w-7xl mx-auto;
}

.blog-header {
  @apply pb-8 border-b border-accent/20;
}

.category-filter {
  @apply px-4 py-2 rounded-full border border-accent/20 text-sm font-mono
         text-text-secondary hover:text-accent hover:border-accent/40
         transition-all duration-200;
}

.category-filter.active {
  @apply bg-accent/10 text-accent border-accent/40;
}

.category-card {
  @apply p-4 rounded-lg bg-bg-secondary/30 border border-accent/5;
}

.blog-stat {
  @apply p-4 rounded-lg bg-bg-secondary/30 border border-accent/10 text-center;
}

.blog-stat .stat-number {
  @apply text-2xl font-bold text-accent font-mono;
}

.blog-stat .stat-label {
  @apply text-text-secondary text-sm mt-1;
}
</style>