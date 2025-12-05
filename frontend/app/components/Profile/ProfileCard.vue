<template>
  <div class="profile-section">
    <div class="flex flex-col md:flex-row gap-6">
      <div v-if="profile.pic" class="flex-shrink-0">
        <img 
          :src="profile.pic" 
          :alt="profile.name"
          class="profile-pic object-cover"
        />
      </div>
      
      <div class="flex-1">
        <h2 class="text-2xl font-bold mb-2 text-accent">{{ profile.name }}</h2>
        <p class="text-text-secondary mb-4">{{ profile.bio }}</p>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-2 text-sm font-mono">
          <div v-if="profile.cake_day" class="flex">
            <span class="text-accent mr-2">birthday:</span>
            <span class="text-text-primary">{{ profile.cake_day }}</span>
          </div>
          <div v-if="profile.location" class="flex">
            <span class="text-accent mr-2">location:</span>
            <span class="text-text-primary">{{ profile.location }}</span>
          </div>
          <div v-if="profile.timezone" class="flex">
            <span class="text-accent mr-2">timezone:</span>
            <span class="text-text-primary">{{ profile.timezone }}</span>
          </div>
          <div v-if="profile.email" class="flex">
            <span class="text-accent mr-2">email:</span>
            <a :href="`mailto:${profile.email}`" class="profile-link">
              {{ profile.email }}
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Profile } from '../../../types'

defineProps<{
  profile: Profile
}>()
</script>