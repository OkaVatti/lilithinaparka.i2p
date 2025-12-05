<template>
  <div class="max-w-4xl mx-auto">
    <div class="mb-8">
      <h1 class="text-3xl font-bold mb-4 font-mono">
        <span class="text-accent">$</span> blog
      </h1>
      <p class="text-text-secondary">
        thoughts, tutorials, and random musings
      </p>
    </div>
    
    <!-- Category Filter -->
    <div class="mb-6">
      <div class="flex flex-wrap gap-2 mb-4">
        <button 
          @click="selectCategory(null)"
          :class="[
            'px-3 py-1 border text-sm transition-colors font-mono',
            filterCategory === null 
              ? 'border-accent bg-accent text-bg-primary' 
              : 'border-border hover:border-accent text-text-primary'
          ]"
        >
          all
        </button>
        <button 
          v-for="cat in categories" 
          :key="cat"
          @click="selectCategory(cat)"
          :class="[
            'px-3 py-1 border text-sm transition-colors font-mono',
            filterCategory === cat 
              ? 'border-accent bg-accent text-bg-primary' 
              : 'border-border hover:border-accent text-text-primary'
          ]"
        >
          {{ cat.toLowerCase() }}
        </button>
      </div>
      
      <!-- Tag Filter (if tag in URL) -->
      <div v-if="selectedTag" class="terminal-box text-sm">
        <span class="text-text-secondary">filtering by tag:</span>
        <span class="text-accent ml-2">#{{ selectedTag }}</span>
        <button 
          @click="clearTagFilter"
          class="ml-2 text-link hover:text-accent"
        >
          [clear]
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
    <div v-else-if="filteredPosts.length === 0" class="terminal-box">
      <p class="text-text-secondary">no posts found</p>
    </div>
    
    <!-- Posts List -->
    <div v-else>
      <div class="mb-4 text-sm text-text-secondary font-mono">
        showing {{ filteredPosts.length }} post{{ filteredPosts.length !== 1 ? 's' : '' }}
      </div>
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
const route = useRoute()
const router = useRouter()

const filterCategory = ref<string | null>(null)
const selectedTag = ref<string | null>(null)

const loading = computed(() => blogStore.loading)
const error = computed(() => blogStore.error)

const categories = ['Casual', 'Interlude', 'Serious']

const filteredPosts = computed(() => {
  let posts = blogStore.publishedPosts
  
  if (selectedTag.value) {
    posts = posts.filter(post => {
      try {
        const tags = JSON.parse(post.tags)
        return tags.includes(selectedTag.value)
      } catch {
        return false
      }
    })
  }
  
  if (filterCategory.value) {
    posts = posts.filter(post => {
      try {
        const categories = JSON.parse(post.categories)
        return categories.includes(filterCategory.value)
      } catch {
        return false
      }
    })
  }
  
  return posts
})

const selectCategory = (category: string | null) => {
  filterCategory.value = category
}

const clearTagFilter = () => {
  selectedTag.value = null
  router.push('/blog')
}

onMounted(async () => {
  await blogStore.fetchPosts()
  
  // Check for tag filter in URL
  if (route.query.tag) {
    selectedTag.value = route.query.tag as string
  }
})

watch(() => route.query.tag, (newTag) => {
  selectedTag.value = newTag ? newTag as string : null
})

useHead({
  title: 'Blog - Lilith Parker',
  meta: [
    { name: 'description', content: 'thoughts, tutorials, and random musings' }
  ]
})
</script>