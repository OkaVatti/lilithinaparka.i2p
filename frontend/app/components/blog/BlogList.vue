<template>
  <div>
    <div v-if="loading" class="text-gray-600">Loading…</div>
    <div v-else class="grid gap-4">
      <BlogCard v-for="post in postsToShow" :key="post.slug" :post="post" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watchEffect } from 'vue'
import BlogCard from './/BlogCard.vue'

const props = defineProps({
  initialPosts: { type: Array as () => any[], default: () => [] },
  loading: { type: Boolean, default: false }
})

const posts = ref<any[]>(props.initialPosts || [])

watchEffect(() => {
  posts.value = props.initialPosts || []
})

const postsToShow = computed(() => posts.value)
</script>
