<template>
  <div class="max-w-3xl mx-auto py-10 px-4">
    <div v-if="loading" class="text-gray-600">Loading profile…</div>

    <div v-else>
      <div v-if="!user">
        <h1 class="text-2xl font-semibold mb-2">Not signed in</h1>
        <p class="text-gray-600 mb-4">Sign in to manage your profile, posts, and uploads.</p>
        <form @submit.prevent="doLogin" class="max-w-md">
          <label class="block mb-2">
            <span class="text-sm text-gray-700">Username</span>
            <input v-model="creds.username" class="w-full border rounded px-2 py-1 mt-1" />
          </label>
          <label class="block mb-4">
            <span class="text-sm text-gray-700">Password</span>
            <input v-model="creds.password" type="password" class="w-full border rounded px-2 py-1 mt-1" />
          </label>
          <button class="px-3 py-1 bg-black text-white rounded">Sign in</button>
        </form>
        <div v-if="loginError" class="text-red-500 mt-2">{{ loginError }}</div>
      </div>

      <div v-else>
        <h1 class="text-2xl font-bold mb-2">Welcome, {{ user.name || user.handle }}</h1>
        <ProfileCard :user="profileData" />
        <section class="mt-6">
          <h2 class="text-lg font-semibold">Your recent posts</h2>
          <ul class="mt-2">
            <li v-for="p in myPosts" :key="p.slug" class="py-2 border-b">
              <NuxtLink :to="`/blog/${p.slug}`" class="text-sm">{{ p.title }}</NuxtLink>
            </li>
          </ul>
        </section>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import ProfileCard from '~/components/profile/ProfileCard.vue'
import { useAuth } from '~/composables/useAuth'
import { useApi } from '~/composables/useApi'

const { user, login, isAuthenticated, refreshStatus } = useAuth()
const { apiFetch } = useApi()

const loading = ref(true)
const profileData = ref<any | null>(null)
const myPosts = ref<any[]>([])
const creds = ref({ username: '', password: '' })
const loginError = ref<string | null>(null)

async function loadProfile() {
  loading.value = true
  try {
    if (isAuthenticated.value && user.value) {
      // try backend profile endpoint
      try {
        const res = await apiFetch('/auth/profile')
        profileData.value = res.user || res
      } catch {
        profileData.value = user.value
      }
      // fetch posts by current user if backend supports /blog/posts?author=me or /blog/mine
      try {
        const postsRes = await apiFetch('/blog/mine')
        myPosts.value = Array.isArray(postsRes) ? postsRes : (postsRes.items || [])
      } catch {
        myPosts.value = []
      }
    } else {
      profileData.value = null
    }
  } catch (e) {
    console.error('profile load failed', e)
  } finally {
    loading.value = false
  }
}

async function doLogin() {
  loginError.value = null
  try {
    await login(creds.value.username, creds.value.password)
    await refreshStatus()
    await loadProfile()
  } catch (e: any) {
    loginError.value = e?.message || 'Login failed'
  }
}

onMounted(loadProfile)
</script>
