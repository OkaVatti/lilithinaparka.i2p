<template>
  <div class="max-w-3xl mx-auto py-10 px-4">
    <div v-if="loading" class="text-gray-600">Loading…</div>

    <div v-else-if="profile">
      <ProfileCard :user="profile" />
      <section class="mt-6">
        <h2 class="text-lg font-semibold">Posts by {{ profile.displayName || profile.handle }}</h2>
        <ul class="mt-2">
          <li v-for="p in posts" :key="p.slug" class="py-2 border-b">
            <NuxtLink :to="`/blog/${p.slug}`" class="text-sm">{{ p.title }}</NuxtLink>
          </li>
        </ul>
      </section>
    </div>

    <div v-else class="text-gray-500">Profile not found.</div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import ProfileCard from '~/components/profile/ProfileCard.vue'
import { useApi } from '~/composables/useApi'

const route = useRoute()
const handle = String(route.params.handle || '')
const { apiFetch } = useApi()

const profile = ref<any | null>(null)
const posts = ref<any[]>([])
const loading = ref(true)

async function load() {
  loading.value = true
  try {
    // backend public user endpoint could be /users/:handle
    const res = await apiFetch(`/users/${encodeURIComponent(handle)}`)
    profile.value = res || null
    try {
      const postsRes = await apiFetch(`/blog/posts?author=${encodeURIComponent(handle)}`)
      posts.value = Array.isArray(postsRes) ? postsRes : (postsRes.items || [])
    } catch {
      posts.value = []
    }
  } catch (e) {
    console.error('failed to load profile', e)
    profile.value = null
  } finally {
    loading.value = false
  }
}

onMounted(load)
</script>
