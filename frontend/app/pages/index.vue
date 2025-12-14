<template>
  <div class="home-page">
    <div class="container">
      <section class="hero window">
        <div class="title-bar">
          <div class="title-bar-text">Welcome</div>
        </div>
        <div class="window-body">
          <div class="hero-content">
            <h1>{{ siteName }}</h1>
            <p class="tagline">
              Privacy-first, retro-inspired eepsite on I2P
            </p>
            <p class="description">
              A minimalist blog, game suite, and art gallery with a classic terminal aesthetic.
              Built for the dark web with privacy and simplicity in mind.
            </p>
            
            <div class="quick-links">
              <NuxtLink to="/blog">
                <button>📖 Read Blog</button>
              </NuxtLink>
              <NuxtLink to="/games">
                <button>🎮 Play Games</button>
              </NuxtLink>
              <NuxtLink to="/profile">
                <button>👤 View Profile</button>
              </NuxtLink>
            </div>
          </div>
        </div>
      </section>

      <section class="recent-posts">
        <div class="window">
          <div class="title-bar">
            <div class="title-bar-text">Latest Posts</div>
          </div>
          <div class="window-body">
            <div v-if="loading" class="loading">Loading posts</div>
            <div v-else-if="recentPosts.length > 0">
              <BlogList :initialPosts="recentPosts" />
            </div>
            <div v-else class="no-content">
              <p>No posts yet. Check back soon!</p>
            </div>
          </div>
        </div>
      </section>

      <section v-if="profileStore.profile" class="profile-preview">
        <div class="window">
          <div class="title-bar">
            <div class="title-bar-text">Profile</div>
          </div>
          <div class="window-body">
            <ProfileCard :user="profileStore.profile" />
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from '@vue/runtime-core'
import { useRuntimeConfig } from '#imports'
import { useApi } from '~/composables/useApi'
import { useProfileStore } from '~~/stores/profile'
import BlogList from '../components/blog/BlogList.vue'
import ProfileCard from '../components/profile/ProfileCard.vue'

const config = useRuntimeConfig()
const siteName = config.public.siteName || 'lilithinaparka.i2p'

const { apiFetch } = useApi()
const profileStore = useProfileStore()

const recentPosts = ref<any[]>([])
const loading = ref(true)

async function loadRecent() {
  loading.value = true
  try {
    const res = await apiFetch('/blog/posts?limit=6')
    recentPosts.value = Array.isArray(res) ? res : (res.items || [])
  } catch (e) {
    console.error('Failed to fetch recent posts', e)
    recentPosts.value = []
  } finally {
    loading.value = false
  }
}

onMounted(async () => {
  await Promise.all([
    loadRecent(),
    profileStore.fetchProfile()
  ])
})

useHead({
  title: siteName,
  meta: [
    { name: 'description', content: 'Privacy-first retro blog and game portal on I2P' }
  ]
})
</script>

<style scoped>
.home-page {
  padding: 2rem 0;
}

.hero {
  margin-bottom: 2rem;
}

.hero-content {
  text-align: center;
  padding: 2rem 1rem;
}

.hero-content h1 {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
}

.tagline {
  font-size: 1.2rem;
  margin-bottom: 1rem;
  color: var(--theme-accent);
}

.description {
  max-width: 600px;
  margin: 0 auto 2rem;
  line-height: 1.8;
}

.quick-links {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.quick-links button {
  min-width: 150px;
}

.recent-posts,
.profile-preview {
  margin-bottom: 2rem;
}

.no-content {
  text-align: center;
  padding: 2rem;
  color: var(--theme-fg);
  opacity: 0.7;
}

@media (max-width: 768px) {
  .hero-content h1 {
    font-size: 1.8rem;
  }
  
  .quick-links {
    flex-direction: column;
    align-items: stretch;
  }
}
</style>