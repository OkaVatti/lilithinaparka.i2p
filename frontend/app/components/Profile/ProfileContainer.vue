<template>
  <div class="profile-pic-wrapper" :class="{ 'has-banner': profile.bsky_banner }">
    <!-- Banner (if available) -->
    <div v-if="profile.bsky_banner" class="profile-banner">
      <img 
        :src="profile.bsky_banner"
        :alt="profile.name + ' banner'"
        class="banner-image"
        loading="lazy"
      />
    </div>
    
    <!-- Profile Picture -->
    <div class="profile-pic-container" :class="{ 'with-banner': profile.bsky_banner }">
      <img 
        :src="profile.pic || profile.bsky_avatar"
        :alt="profile.name"
        class="profile-pic"
        loading="lazy"
        @error="handleImageError"
      />
      
      <!-- Online indicator (if you want to add status) -->
      <div v-if="isOnline" class="online-indicator"></div>
      
      <!-- Hover overlay -->
      <div class="profile-overlay">
        <span class="overlay-text">View Profile</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Profile } from "../../../types"
import { computed, ref } from 'vue'

const props = defineProps<{
  profile: Profile
}>()

// You can add logic for online status if needed
const isOnline = ref(false)

// Default avatar if image fails to load
const defaultAvatar = 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAwIiBoZWlnaHQ9IjIwMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cmVjdCB3aWR0aD0iMjAwIiBoZWlnaHQ9IjIwMCIgZmlsbD0iIzhiOGJlOSIvPjx0ZXh0IHg9IjUwJSIgeT0iNTAlIiBkb21pbmFudC1iYXNlbGluZT0ibWlkZGxlIiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmb250LWZhbWlseT0ibW9ub3NwYWNlIiBmb250LXNpemU9IjQ4IiBmaWxsPSJ3aGl0ZSI+TDwvdGV4dD48L3N2Zz4='

const handleImageError = (event: Event) => {
  const img = event.target as HTMLImageElement
  img.src = defaultAvatar
}

// If you want to compute initials for the fallback
const initials = computed(() => {
  if (!props.profile.name) return 'LP'
  return props.profile.name
    .split(' ')
    .map((word: string) => word[0])
    .join('')
    .toUpperCase()
    .slice(0, 2)
})
</script>

<style scoped>
.profile-pic-wrapper {
  @apply relative;
}

.profile-pic-wrapper.has-banner {
  @apply pb-16;
}

.profile-banner {
  @apply absolute top-0 left-0 right-0 h-32 overflow-hidden rounded-t-2xl;
}

.banner-image {
  @apply w-full h-full object-cover opacity-60;
}

.profile-pic-container {
  @apply relative w-48 h-48 overflow-hidden rounded-2xl border-4 
         border-background bg-background-secondary shadow-2xl;
  aspect-ratio: 1;
}

.profile-pic-container.with-banner {
  @apply absolute bottom-0 left-1/2 transform -translate-x-1/2 translate-y-1/2;
}

.profile-pic {
  @apply w-full h-full object-cover transition-all duration-500 
         hover:scale-110 hover:rotate-3;
}

.profile-overlay {
  @apply absolute inset-0 bg-accent/0 flex items-center justify-center 
         opacity-0 transition-all duration-300 hover:bg-accent/20 hover:opacity-100;
}

.overlay-text {
  @apply text-white font-mono text-sm font-bold bg-accent/80 px-3 py-1 
         rounded-full transform -translate-y-2 opacity-0 transition-all duration-300;
}

.profile-overlay:hover .overlay-text {
  @apply translate-y-0 opacity-100;
}

.online-indicator {
  @apply absolute bottom-2 right-2 w-4 h-4 rounded-full bg-success 
         border-2 border-background z-10;
}

/* Glitch effect on hover */
@keyframes glitch {
  0% {
    clip-path: inset(40% 0 61% 0);
    transform: translate(-2px, 2px);
  }
  5% {
    clip-path: inset(92% 0 1% 0);
    transform: translate(2px, -2px);
  }
  10% {
    clip-path: inset(43% 0 1% 0);
    transform: translate(-2px, -2px);
  }
  15% {
    clip-path: inset(25% 0 58% 0);
    transform: translate(2px, 2px);
  }
  20% {
    clip-path: inset(54% 0 7% 0);
    transform: translate(2px, -2px);
  }
  45% {
    clip-path: inset(58% 0 43% 0);
    transform: translate(-2px, 2px);
  }
  50% {
    clip-path: inset(98% 0 1% 0);
    transform: translate(2px, -2px);
  }
  55% {
    clip-path: inset(48% 0 18% 0);
    transform: translate(-2px, 2px);
  }
  60% {
    clip-path: inset(1% 0 60% 0);
    transform: translate(2px, 2px);
  }
  65% {
    clip-path: inset(75% 0 9% 0);
    transform: translate(2px, -2px);
  }
  70% {
    clip-path: inset(63% 0 14% 0);
    transform: translate(-2px, 2px);
  }
  75% {
    clip-path: inset(1% 0 83% 0);
    transform: translate(-2px, -2px);
  }
  80% {
    clip-path: inset(64% 0 5% 0);
    transform: translate(2px, 2px);
  }
  85% {
    clip-path: inset(26% 0 56% 0);
    transform: translate(-2px, -2px);
  }
  90% {
    clip-path: inset(78% 0 2% 0);
    transform: translate(2px, -2px);
  }
  95% {
    clip-path: inset(39% 0 47% 0);
    transform: translate(-2px, 2px);
  }
  100% {
    clip-path: inset(1% 0 60% 0);
    transform: translate(2px, -2px);
  }
}

.profile-pic:hover {
  animation: glitch 0.5s infinite;
}
</style>