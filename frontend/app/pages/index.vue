<template>
  <div class="max-w-5xl mx-auto py-10 px-4">
    <section class="mb-12">
      <div class="flex items-center justify-between gap-4">
        <div>
          <h1 class="text-4xl font-extrabold">{{ siteName }}</h1>
          <p class="text-gray-600 mt-2 max-w-xl">
            Classic macOS inspired eepsite — privacy-first, retro UI, blog, games and art.
          </p>
          <div class="mt-4 flex gap-3">
            <NuxtLink to="/blog" class="btn">Read the blog</NuxtLink>
            <NuxtLink to="/profile" class="btn btn-outline">My profile</NuxtLink>
          </div>
        </div>

        <div class="hidden md:block w-56">
          <ProfileCard :user="currentUserPreview" v-if="currentUserPreview" />
        </div>
      </div>
    </section>

    <section>
      <h2 class="text-2xl font-semibold mb-4">Latest posts</h2>
      <BlogList :initialPosts="recentPosts" :loading="loading" />
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRuntimeConfig } from '#imports'
import { useApi } from '~/composables/useApi'
import BlogList from '../components/blog/BlogList.vue'
import ProfileCard from '../components/profile/ProfileCard.vue'
import { useAuth } from '~/composables/useAuth'

const config = useRuntimeConfig()
const siteName = config.public.siteName || 'lilithinaparka.i2p'

const { apiFetch } = useApi()
const recentPosts = ref<any[]>([])
const loading = ref(true)

const { user, isAuthenticated } = useAuth()
const currentUserPreview = ref(null)

async function loadRecent() {
  loading.value = true
  try {
    // prefer server-side limit param if backend supports it
    const res = await apiFetch('/blog/posts?limit=6')
    recentPosts.value = Array.isArray(res) ? res : (res.items || [])
  } catch (e) {
    console.error('Failed to fetch recent posts', e)
    recentPosts.value = []
  } finally {
    loading.value = false
  }
}

if (isAuthenticated.value && user.value) {
  currentUserPreview.value = {
    displayName: user.value.name || user.value.handle || 'You',
    handle: user.value.handle || '',
    bio: user.value.bio || 'Private profile'
  }
}

// Load on mount
loadRecent()
</script>

<style scoped>
.btn {
  @apply bg-black text-white px-3 py-1 rounded;
}
.btn-outline {
  @apply border border-black px-3 py-1 rounded;
}
</style>
