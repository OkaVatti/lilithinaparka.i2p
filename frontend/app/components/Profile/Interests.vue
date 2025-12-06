<template>
  <div class="profile-section">
    <h3 class="text-xl font-bold mb-4 font-mono">
      <span class="text-accent">$</span> interests
    </h3>
    <div class="flex flex-wrap gap-2">
      <span 
        v-for="(interest, index) in interests" 
        :key="index"
        class="interest-tag hover:scale-105 transition-all duration-200"
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

<style scoped>
.interest-tag {
  @apply px-3 py-1.5 rounded-full bg-accent/10 text-accent text-sm font-mono 
         border border-accent/20 hover:bg-accent/20 hover:border-accent/30 
         cursor-default;
}
</style>