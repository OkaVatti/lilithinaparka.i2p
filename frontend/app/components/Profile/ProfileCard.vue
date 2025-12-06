<template>
  <div class="profile-card">
    <div class="flex flex-col md:flex-row gap-8 items-start">
      <!-- Left Column: Profile Image and Basic Info -->
      <div class="md:w-1/3 space-y-6">
        <div class="aspect-square w-full max-w-64 overflow-hidden rounded-2xl border-2 border-accent/20 bg-background-secondary">
          <img 
            :src="profile.pic" 
            :alt="profile.name"
            class="w-full h-full object-cover hover:scale-105 transition-transform duration-300"
          />
        </div>
        
        <div class="space-y-4">
          <div>
            <h1 class="text-3xl font-bold text-accent font-mono">{{ profile.name }}</h1>
            <p class="text-text-secondary font-mono">@{{ profile.username.replace('@', '') }}</p>
          </div>
          
          <div class="space-y-2 text-sm">
            <div v-if="profile.cake_day" class="flex items-center gap-2">
              <span class="text-accent font-mono">age:</span>
              <span class="font-mono">{{ profile.cake_day }}</span>
            </div>
            <div v-if="profile.location" class="flex items-center gap-2">
              <span class="text-accent font-mono">location:</span>
              <span class="font-mono">{{ profile.location }}</span>
            </div>
            <div v-if="profile.timezone" class="flex items-center gap-2">
              <span class="text-accent font-mono">timezone:</span>
              <span class="font-mono">{{ profile.timezone }}</span>
            </div>
            <div v-if="profile.cake_day" class="flex items-center gap-2">
              <span class="text-accent font-mono">cake day:</span>
              <span class="font-mono">{{ profile.cake_day }}</span>
            </div>
          </div>
        </div>
      </div>
      
      <!-- Right Column: Bio, Socials, and Details -->
      <div class="md:w-2/3 space-y-8">
        <!-- Bio Section -->
        <div class="profile-section">
          <h3 class="text-xl font-bold mb-3 font-mono">
            <span class="text-accent">$</span> bio
          </h3>
          <p class="text-text-primary leading-relaxed">{{ profile.bio }}</p>
        </div>
        
        <!-- Social Links -->
        <div class="profile-section">
          <h3 class="text-xl font-bold mb-4 font-mono">
            <span class="text-accent">$</span> links
          </h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
            <a v-if="profile.email" :href="`mailto:${profile.email}`" 
               class="social-link group">
              <span class="text-accent font-mono">email:</span>
              <span>{{ profile.email }}</span>
            </a>
            <a v-if="profile.website" :href="profile.website" target="_blank"
               class="social-link group">
              <span class="text-accent font-mono">website:</span>
              <span>{{ profile.website.replace('https://', '') }}</span>
            </a>
            <a v-if="profile.github" :href="profile.github" target="_blank"
               class="social-link group">
              <span class="text-accent font-mono">github:</span>
              <span>{{ profile.github.replace('https://github.com/', '') }}</span>
            </a>
            <a v-if="profile.gitten" :href="profile.gitten" target="_blank"
               class="social-link group">
              <span class="text-accent font-mono">gitten:</span>
              <span>{{ profile.gitten.split('/')[0] }}</span>
            </a>
            <div v-if="profile.bluesky" class="social-link">
              <span class="text-accent font-mono">bluesky:</span>
              <span>{{ profile.bluesky }}</span>
            </div>
            <div v-if="profile.bottletail" class="social-link">
              <span class="text-accent font-mono">bottletail:</span>
              <span>{{ profile.bottletail }}</span>
            </div>
            <a v-if="profile.rss" :href="profile.rss" target="_blank"
               class="social-link group">
              <span class="text-accent font-mono">rss:</span>
              <span>rss.xml</span>
            </a>
          </div>
        </div>
        
        <!-- Interests (if provided) -->
        <div v-if="profile.interests && profile.interests.length" class="profile-section">
          <h3 class="text-xl font-bold mb-4 font-mono">
            <span class="text-accent">$</span> interests
          </h3>
          <div class="flex flex-wrap gap-2">
            <span 
              v-for="(interest, index) in profile.interests" 
              :key="index"
              class="interest-tag"
            >
              {{ interest }}
            </span>
          </div>
        </div>
        
        <!-- Donations Section -->
        <Donations
          :bitcoin="profile.bitcoin"
          :ethereum="profile.ethereum"
          :solana="profile.solana"
          :monero="profile.monero"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { type Profile } from '../../../types'
import Donations from './Donations.vue'

defineProps<{
  profile: Profile
}>()
</script>

<style scoped>
.profile-card {
  @apply max-w-6xl mx-auto p-6;
}

.profile-section {
  @apply p-6 rounded-xl bg-background-secondary/50 border border-accent/10;
}

.social-link {
  @apply flex items-center justify-between p-3 rounded-lg bg-background/50 
         hover:bg-background-secondary border border-accent/5 transition-all duration-200;
}

.social-link.group:hover {
  @apply border-accent/30 transform -translate-y-0.5;
}

.interest-tag {
  @apply px-3 py-1.5 rounded-full bg-accent/10 text-accent text-sm font-mono 
         border border-accent/20 hover:bg-accent/20 transition-colors;
}
</style>