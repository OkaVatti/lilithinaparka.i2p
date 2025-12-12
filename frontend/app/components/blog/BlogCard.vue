<template>
  <article class="p-4 border rounded bg-white">
    <div class="flex items-start gap-4">
      <div class="flex-1">
        <h3 class="text-lg font-semibold">
          <NuxtLink :to="`/blog/${post.slug}`">{{ post.title }}</NuxtLink>
        </h3>
        <p class="text-sm text-gray-600 mt-1">{{ excerpt }}</p>
      </div>
      <div v-if="post.featured_image" class="w-24 h-16 overflow-hidden rounded">
        <img :src="post.featured_image" alt="" class="object-cover w-full h-full" />
      </div>
    </div>
  </article>
</template>

<script setup lang="ts">
import { computed } from 'vue'
const props = defineProps<{ post: any }>()

const excerpt = computed(() => {
  if (!props.post) return ''
  if (props.post.excerpt) return props.post.excerpt
  if (props.post.summary) return props.post.summary
  const txt = props.post.content ? String(props.post.content).replace(/<[^>]+>/g, '') : ''
  return txt.slice(0, 160) + (txt.length > 160 ? '…' : '')
})
</script>
