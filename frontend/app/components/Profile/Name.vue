<template>
  <div class="name-container">
    <h1 class="name-display">{{ name }}</h1>
    <p v-if="username" class="username-display">{{ username }}</p>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { Profile } from '../../../types'

// Props with Profile object OR individual fields for flexibility
const props = withDefaults(defineProps<{
  profile?: Profile  // Optional Profile object
  name?: string      // Optional direct name prop
  username?: string  // Optional direct username prop
}>(), {
  profile: undefined,
  name: '',
  username: ''
})

// Computed properties to handle both ways of passing data
const displayName = computed(() => {
  return props.name || props.profile?.name || ''
})

const displayUsername = computed(() => {
  return props.username || props.profile?.username || ''
})
</script>

<style scoped>
.name-container {
  @apply space-y-2;
}

.name-display {
  @apply text-4xl font-bold text-accent font-mono 
         bg-gradient-to-r from-accent to-purple-500 bg-clip-text text-transparent
         transition-all duration-500 hover:tracking-wider;
  text-shadow: 0 0 10px rgba(139, 139, 233, 0.3);
}

.username-display {
  @apply text-lg text-text-secondary font-mono 
         px-4 py-2 rounded-lg bg-background-secondary/50 
         border border-accent/10 inline-block
         transition-all duration-300 hover:bg-background-secondary 
         hover:border-accent/20;
}

/* Glitch effect on hover */
@keyframes glitch {
  0% {
    transform: translate(0);
  }
  20% {
    transform: translate(-2px, 2px);
  }
  40% {
    transform: translate(-2px, -2px);
  }
  60% {
    transform: translate(2px, 2px);
  }
  80% {
    transform: translate(2px, -2px);
  }
  100% {
    transform: translate(0);
  }
}

.name-display:hover {
  animation: glitch 0.5s infinite;
}

/* Terminal cursor effect */
.username-display::after {
  content: '▋';
  @apply text-accent ml-1 opacity-0;
  animation: blink 1s infinite;
}

.username-display:hover::after {
  @apply opacity-100;
}

@keyframes blink {
  0%, 100% { opacity: 1; }
  50% { opacity: 0; }
}
</style>