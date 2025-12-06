<template>
  <div class="profile-pic-container">
    <img 
      :src="profile.pic"
      :alt="profile.name"
      class="profile-pic"
      loading="lazy"
    />
  </div>
</template>

<script setup lang="ts">
import type { Profile } from "../../../types"

// Note: Changed from 'Profile' to 'profile' (lowercase) for Vue convention
defineProps<{
  profile: Profile
}>()
</script>

<style scoped>
.profile-pic-container {
  @apply relative overflow-hidden rounded-2xl border-2 border-accent/20 
         bg-background-secondary shadow-lg;
  aspect-ratio: 1;
}

.profile-pic {
  @apply w-full h-full object-cover transition-all duration-500 
         hover:scale-105 hover:rotate-2;
}

.profile-pic:hover {
  filter: saturate(1.2) contrast(1.1);
}

/* Fallback for broken images */
.profile-pic:before {
  content: attr(alt);
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, rgba(139, 139, 233, 0.1), rgba(139, 139, 233, 0.05));
  color: var(--color-accent);
  font-family: monospace;
  font-weight: bold;
  font-size: 1.5rem;
}

.profile-pic:not([src]):before,
.profile-pic[src=""]:before {
  display: flex;
}
</style>