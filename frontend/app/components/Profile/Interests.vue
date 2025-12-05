<template>
  <div class="profile-section">
    <h3 class="text-xl font-bold mb-3 font-mono">
      <span class="text-accent">$</span> interests
    </h3>
    <div class="flex flex-wrap gap-2">
      <span 
        v-for="interest in interests" 
        :key="interest"
        class="profile-interest-tag"
      >
        {{ interest }}
      </span>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  interests: string[]
}>()
</script>