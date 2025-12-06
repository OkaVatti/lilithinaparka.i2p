<template>
  <div class="profile-section">
    <h3 class="text-xl font-bold mb-3 font-mono">
      <span class="text-accent">$</span> bio
    </h3>
    <p class="text-text-primary leading-relaxed">{{ profile.bio }}</p>
  </div>
</template>

<script setup lang="ts">
import type { Profile } from '../../../types';

defineProps<{
    profile: Profile
}>()
</script>