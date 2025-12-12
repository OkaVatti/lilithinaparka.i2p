<template>
  <div class="max-w-4xl mx-auto py-10 px-4">
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-3xl font-bold">Blog</h1>
      <NuxtLink to="/blog/rss.xml" class="text-sm text-gray-500">RSS</NuxtLink>
    </div>

    <div v-if="loading" class="text-gray-600">Loading posts…</div>
    <div v-else>
      <BlogList :initialPosts="posts" />
      <div v-if="!posts || !posts.length" class="mt-6 text-gray-500">No posts yet.</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useApi } from '~/composables/useApi'
import BlogList from '~/components/blog/BlogList.vue'

const { apiFetch } = useApi()
const posts = ref<any[]>([])
const loading = ref(true)

async function load() {
  loading.value = true
  try {
    const res = await apiFetch('/blog/posts')
    // backend may return either array or { items, pagination }
    posts.value = Array.isArray(res) ? res : (res.items || [])
  } catch (e) {
    console.error('Failed to load blog posts', e)
    posts.value = []
  } finally {
    loading.value = false
  }
}

onMounted(load)
</script>
