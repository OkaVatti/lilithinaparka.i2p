<template>
  <div class="max-w-3xl mx-auto py-10 px-4">
    <div v-if="loading" class="text-gray-600">Loading…</div>

    <article v-else>
      <header class="mb-6">
        <h1 class="text-3xl font-bold">{{ post.title }}</h1>
        <div class="text-sm text-gray-500 mt-1">
          <span v-if="post.date">{{ formatDate(post.date) }}</span>
          <span v-if="authors.length"> — {{ authors.join(', ') }}</span>
        </div>
      </header>

      <div class="prose max-w-none">
        <!-- Render sanitized HTML on client; fall back to escaped text on server -->
        <div v-if="isClient" v-html="safeHtml"></div>
        <pre v-else class="whitespace-pre-wrap">{{ post.content }}</pre>
      </div>
    </article>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useApi } from '~/composables/useApi'
import { useRoute } from 'vue-router'
import { useNuxtApp } from '#app'

const route = useRoute()
const slug = String(route.params.slug || '')
const { apiFetch } = useApi()

const post = ref<any>({ title: '', content: '', date: '' })
const loading = ref(true)
const authors = ref<string[]>([])
const isClient = ref(false)

const nuxtApp = useNuxtApp()
const $sanitize = (nuxtApp as any).$sanitize as ((s: string) => string) | undefined

function sanitizeServerSide(html: string) {
  if (!html) return ''
  // naive server-side sanitizer: remove script tags and on* attributes
  let out = String(html).replace(/<script[\s\S]*?>[\s\S]*?<\/script>/gi, '')
  // remove attributes starting with on (onclick, onerror, etc.)
  out = out.replace(/\s(on\w+)=(".*?"|'.*?'|[^\s>]+)/gi, '')
  return out
}

const safeHtml = computed(() => {
  if (!post.value?.content) return ''
  if (isClient.value && $sanitize) {
    return $sanitize(post.value.content)
  }
  return sanitizeServerSide(post.value.content)
})

function formatDate(d?: string) {
  if (!d) return ''
  try {
    return new Date(d).toLocaleString()
  } catch {
    return d
  }
}

onMounted(async () => {
  isClient.value = true
})

async function load() {
  loading.value = true
  try {
    const res = await apiFetch(`/blog/posts/${encodeURIComponent(slug)}`)
    post.value = res || { title: 'Not found', content: '' }
    authors.value = res?.authors ? String(res.authors).split(',').map((s: string) => s.trim()).filter(Boolean) : []
  } catch (e) {
    console.error('Failed to load post', e)
    post.value = { title: 'Error', content: '' }
  } finally {
    loading.value = false
  }
}

load()
</script>
