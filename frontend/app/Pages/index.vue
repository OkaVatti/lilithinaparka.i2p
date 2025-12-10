<template>
  <div class="home-page">
    <section class="hero">
      <h1 class="hero-title">Welcome to {{ profileStore.profile?.name || 'Lilith in a Parka' }}</h1>
      <p class="hero-subtitle">{{ profileStore.profile?.bio || 'A minimalist blog on I2P' }}</p>
      <div class="hero-actions">
        <NuxtLink to="/blog" class="btn">
          <FeatherIcon name="book-open" size="18" />
          <span>Read Blog</span>
        </NuxtLink>
        <NuxtLink to="/profile" class="btn btn-secondary">
          <FeatherIcon name="user" size="18" />
          <span>View Profile</span>
        </NuxtLink>
      </div>
    </section>
    
    <section class="recent-posts">
      <h2>Recent Posts</h2>
      
      <div v-if="blogStore.loading" class="loading">Loading posts</div>
      
      <div v-else-if="blogStore.error" class="error">
        <FeatherIcon name="alert-circle" size="20" />
        <span>{{ blogStore.error }}</span>
      </div>
      
      <div v-else-if="recentPosts.length" class="grid grid-2">
        <BlogPostCard
          v-for="post in recentPosts"
          :key="post.id"
          :post="post"
        />
      </div>
      
      <p v-else class="no-posts">No posts available yet.</p>
      
      <div v-if="recentPosts.length" class="view-all">
        <NuxtLink to="/blog" class="btn">
          View All Posts
          <FeatherIcon name="arrow-right" size="18" />
        </NuxtLink>
      </div>
    </section>
    
    <section v-if="bskyStore.posts.length" class="bsky-feed">
      <h2>Recent BlueSky Posts</h2>
      <div class="grid grid-2">
        <BskyPostCard
          v-for="post in recentBskyPosts"
          :key="post.id"
          :post="post"
        />
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { useBlogStore } from '~~/stores/blog'
import { useBskyStore } from '~~/stores/bsky'
import { useProfileStore } from '~~/stores/profile'

const blogStore = useBlogStore()
const bskyStore = useBskyStore()
const profileStore = useProfileStore()

const recentPosts = computed(() => blogStore.publishedPosts.slice(0, 6))
const recentBskyPosts = computed(() => bskyStore.posts.slice(0, 4))

onMounted(async () => {
  await Promise.all([
    blogStore.fetchPosts(),
    bskyStore.fetchPosts(10)
  ])
})
</script>

<style scoped>
.hero {
  text-align: center;
  padding: 4rem 0;
  border-bottom: 2px solid var(--theme-border);
  margin-bottom: 3rem;
}

.hero-title {
  font-size: 3rem;
  margin-bottom: 1rem;
  color: var(--theme-primary);
}

.hero-subtitle {
  font-size: 1.5rem;
  color: var(--theme-fg);
  opacity: 0.8;
  margin-bottom: 2rem;
}

.hero-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.hero-actions .btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 1.1rem;
  padding: 0.75rem 1.5rem;
}

.recent-posts,
.bsky-feed {
  margin-bottom: 3rem;
}

.recent-posts h2,
.bsky-feed h2 {
  margin-bottom: 2rem;
  font-size: 2rem;
  color: var(--theme-primary);
  text-align: center;
}

.no-posts {
  text-align: center;
  padding: 2rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

.view-all {
  text-align: center;
  margin-top: 2rem;
}

.view-all .btn {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
}

@media (max-width: 768px) {
  .hero-title {
    font-size: 2rem;
  }
  
  .hero-subtitle {
    font-size: 1.2rem;
  }
}
</style>