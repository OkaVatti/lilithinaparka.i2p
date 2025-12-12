<template>
  <div class="p-4 border rounded bg-white">
    <div class="flex items-center gap-4">
      <div class="w-16 h-16 rounded-full bg-gray-200 flex items-center justify-center text-xl font-semibold text-gray-700">
        {{ avatarInitial }}
      </div>
      <div>
        <div class="text-lg font-semibold">{{ displayName }}</div>
        <div class="text-sm text-gray-500">{{ handleDisplay }}</div>
      </div>
    </div>

    <p v-if="user.bio" class="mt-3 text-sm text-gray-700">{{ user.bio }}</p>

    <div class="mt-3 text-xs text-gray-500">
      <div v-if="user.location">📍 {{ user.location }}</div>
      <div v-if="user.website">🔗 <a :href="user.website" target="_blank" rel="noopener" class="underline">{{ user.website }}</a></div>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{ user: any }>()
const user = props.user || {}

const avatarInitial = user.displayName ? user.displayName[0].toUpperCase() : (user.handle ? user.handle[0].toUpperCase() : '?')
const displayName = user.displayName || user.name || user.handle || 'Unknown'
const handleDisplay = user.handle ? `@${user.handle}` : ''
</script>
