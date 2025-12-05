<template>
  <div class="max-w-4xl mx-auto">
    <div class="mb-8">
      <h1 class="text-3xl font-bold mb-4">
        <span class="text-accent">$</span> profile
      </h1>
    </div>
    
    <div v-if="loading" class="terminal-box">
      <p class="text-text-secondary">loading<span class="blink">_</span></p>
    </div>
    
    <div v-else-if="error" class="terminal-box">
      <p class="text-red-500">error: {{ error }}</p>
    </div>
    
    <div v-else-if="profile">
      <div class="terminal-box mb-6">
        <div class="flex flex-col md:flex-row gap-6">
          <div v-if="profile.pic" class="flex-shrink-0">
            <img 
              :src="profile.pic" 
              :alt="profile.name"
              class="w-32 h-32 border border-accent"
            />
          </div>
          
          <div class="flex-1">
            <h2 class="text-2xl font-bold mb-2">{{ profile.name }}</h2>
            <p class="text-text-secondary mb-4">{{ profile.bio }}</p>
            
            <div class="grid grid-cols-1 md:grid-cols-2 gap-2 text-sm">
              <div v-if="profile.cake_day">
                <span class="text-accent">birthday:</span> {{ profile.cake_day }}
              </div>
              <div v-if="profile.location">
                <span class="text-accent">location:</span> {{ profile.location }}
              </div>
              <div v-if="profile.timezone">
                <span class="text-accent">timezone:</span> {{ profile.timezone }}
              </div>
              <div v-if="profile.email">
                <span class="text-accent">email:</span> 
                <a :href="`mailto:${profile.email}`" class="ml-1">{{ profile.email }}</a>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <div v-if="interests.length > 0" class="terminal-box mb-6">
        <h3 class="text-xl font-bold mb-3">
          <span class="text-accent">$</span> interests
        </h3>
        <div class="flex flex-wrap gap-2">
          <span 
            v-for="interest in interests" 
            :key="interest"
            class="px-3 py-1 bg-bg-primary border border-accent text-accent text-sm"
          >
            {{ interest }}
          </span>
        </div>
      </div>
      
      <div class="terminal-box mb-6">
        <h3 class="text-xl font-bold mb-3">
          <span class="text-accent">$</span> links
        </h3>
        <div class="flex flex-col gap-2 text-sm">
          <a v-if="profile.website" :href="profile.website" target="_blank" rel="noopener" class="hover:text-accent">
            → website: {{ profile.website }}
          </a>
          <a v-if="profile.github" :href="profile.github" target="_blank" rel="noopener" class="hover:text-accent">
            → github: {{ profile.github }}
          </a>
          <a v-if="profile.bluesky" :href="profile.bluesky" target="_blank" rel="noopener" class="hover:text-accent">
            → bluesky: {{ profile.bluesky }}
          </a>
          <a v-if="profile.rss_feed" :href="profile.rss_feed" class="hover:text-accent">
            → rss feed: {{ profile.rss_feed }}
          </a>
        </div>
      </div>
      
      <div v-if="profile.bsky_display_name" class="terminal-box mb-6">
        <h3 class="text-xl font-bold mb-3">
          <span class="text-accent">$</span> bluesky_stats
        </h3>
        <div class="flex flex-col md:flex-row gap-6">
          <div v-if="profile.bsky_avatar" class="flex-shrink-0">
            <img 
              :src="profile.bsky_avatar" 
              :alt="profile.bsky_display_name"
              class="w-24 h-24 border border-link rounded-full"
            />
          </div>
          
          <div class="flex-1">
            <h4 class="text-lg font-bold mb-2">{{ profile.bsky_display_name }}</h4>
            <p v-if="profile.bsky_description" class="text-sm text-text-secondary mb-3">
              {{ profile.bsky_description }}
            </p>
            
            <div class="flex gap-6 text-sm">
              <div>
                <span class="text-accent">posts:</span> {{ profile.bsky_posts_count }}
              </div>
              <div>
                <span class="text-accent">followers:</span> {{ profile.bsky_followers_count }}
              </div>
              <div>
                <span class="text-accent">following:</span> {{ profile.bsky_follows_count }}
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <div class="terminal-box">
        <h3 class="text-xl font-bold mb-3">
          <span class="text-accent">$</span> support
        </h3>
        <p class="text-sm text-text-secondary mb-4">
          if you enjoy my content, consider supporting me with crypto
        </p>
        
        <div class="flex flex-col gap-3 text-xs font-mono">
          <div v-if="profile.bitcoin_donation_addr">
            <span class="text-accent">BTC:</span>
            <code class="ml-2 text-text-secondary break-all">{{ profile.bitcoin_donation_addr }}</code>
          </div>
          <div v-if="profile.ethereum_donation_addr">
            <span class="text-accent">ETH:</span>
            <code class="ml-2 text-text-secondary break-all">{{ profile.ethereum_donation_addr }}</code>
          </div>
          <div v-if="profile.solana_donation_addr">
            <span class="text-accent">SOL:</span>
            <code class="ml-2 text-text-secondary break-all">{{ profile.solana_donation_addr }}</code>
          </div>
          <div v-if="profile.monero_donation_addr">
            <span class="text-accent">XMR:</span>
            <code class="ml-2 text-text-secondary break-all">{{ profile.monero_donation_addr }}</code>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useProfileStore } from '../../stores/profile'

const profileStore = useProfileStore()

const loading = computed(() => profileStore.loading)
const error = computed(() => profileStore.error)
const profile = computed(() => profileStore.profile)
const interests = computed(() => profileStore.interests)

onMounted(async () => {
  await profileStore.fetchProfile()
})

useHead({
  title: 'Profile - Lilith Parker',
  meta: [
    { name: 'description', content: 'about me, my interests, and how to reach me' }
  ]
})
</script>