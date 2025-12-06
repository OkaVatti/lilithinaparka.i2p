<template>
  <div class="profile-page">
    <!-- Terminal Header -->
    <div class="terminal-header">
      <div class="terminal-buttons">
        <div class="close"></div>
        <div class="minimize"></div>
        <div class="maximize"></div>
      </div>
      <div class="terminal-title font-mono text-sm">
        [user@lilithinaparka.i2p] >> [~/profile] $ cat ./about.md
      </div>
    </div>
    
    <!-- Loading State -->
    <div v-if="!profile" class="profile-content">
      <div class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-accent"></div>
        <p class="mt-4 text-text-secondary font-mono">Loading profile...</p>
      </div>
    </div>
    
    <!-- Profile Content when loaded -->
    <div v-else class="profile-content">
      <!-- Top Section: Profile Picture and Basic Info -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8 mb-8">
        <!-- Left Column: Profile Picture and Bluesky Info -->
        <div class="lg:col-span-1 space-y-6">
          <!-- Profile Picture with Banner -->
          <div class="profile-pic-container">
            <div v-if="profile.bsky_banner" class="profile-banner">
              <img 
                :src="profile.bsky_banner" 
                :alt="profile.name + ' banner'"
                class="banner-image"
                loading="lazy"
              />
            </div>
            <div class="profile-pic-wrapper" :class="{ 'with-banner': profile.bsky_banner }">
              <img 
                :src="profile.pic || profile.bsky_avatar"
                :alt="profile.name"
                class="profile-pic"
                loading="lazy"
                @error="handleImageError"
              />
              <!-- Online indicator -->
              <div class="online-indicator"></div>
            </div>
          </div>
          
          <!-- Bluesky Stats -->
          <div class="profile-section">
            <div class="flex items-center gap-2 mb-4">
              <svg class="w-5 h-5 text-sky-500" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 10.8c-1.087-2.114-4.046-6.053-6.798-7.995C2.566.944 1.561 1.266.902 1.565.139 1.908 0 3.08 0 3.768c0 .69.378 5.65.624 6.479.815 2.736 3.713 3.66 6.383 3.364.136-.02.275-.039.415-.056-.138.022-.278.04-.417.06-3.241.428-5.268 2.37-5.268 5.14 0 3.064 4.295 5.146 9.663 5.146 5.368 0 9.663-2.082 9.663-5.146 0-2.77-2.027-4.712-5.268-5.14-.139-.02-.279-.038-.417-.06.14.017.279.036.415.056 2.67.296 5.568-.628 6.383-3.364.246-.828.624-5.79.624-6.478 0-.69-.139-1.86-.902-2.206-.659-.298-1.664-.62-4.3 1.24C16.046 4.747 13.087 8.686 12 10.8z"/>
              </svg>
              <h3 class="text-lg font-bold font-mono text-accent">Bluesky Stats</h3>
            </div>
            <div class="grid grid-cols-3 gap-4 text-center">
              <div class="stat-box">
                <div class="stat-number">{{ profile.bsky_followers_count || 0 }}</div>
                <div class="stat-label font-mono text-xs">Followers</div>
              </div>
              <div class="stat-box">
                <div class="stat-number">{{ profile.bsky_follows_count || 0 }}</div>
                <div class="stat-label font-mono text-xs">Following</div>
              </div>
              <div class="stat-box">
                <div class="stat-number">{{ profile.bsky_posts_count || 0 }}</div>
                <div class="stat-label font-mono text-xs">Posts</div>
              </div>
            </div>
            <div v-if="profile.bsky_display_name" class="mt-4 pt-4 border-t border-accent/10">
              <div class="flex items-center gap-2 mb-2">
                <span class="text-accent font-mono text-sm">Display Name:</span>
                <span class="font-mono">{{ profile.bsky_display_name }}</span>
              </div>
              <p v-if="profile.bsky_description" class="text-text-secondary text-sm italic">
                "{{ profile.bsky_description }}"
              </p>
            </div>
          </div>
        </div>
        
        <!-- Right Column: Main Profile Info -->
        <div class="lg:col-span-2 space-y-6">
          <!-- Name and Username -->
          <div class="profile-section">
            <h1 class="text-3xl font-bold text-accent mb-2">{{ profile.name }}</h1>
            <div class="flex items-center gap-4 mb-4">
              <code class="username-badge font-mono">{{ profile.username || '@' + profile.name.toLowerCase().replace(' ', '') }}</code>
              <span class="text-text-secondary text-sm font-mono">
                ID: <span class="text-accent">{{ profile.id }}</span>
              </span>
            </div>
            
            <!-- Bio -->
            <div class="mb-6">
              <h3 class="text-lg font-bold mb-3 font-mono text-accent">$ bio</h3>
              <p class="text-text-primary leading-relaxed">{{ profile.bio }}</p>
            </div>
            
            <!-- Details Grid -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="detail-item" v-if="profile.cake_day">
                <span class="detail-label font-mono">Cake Day:</span>
                <span class="detail-value">{{ formatDate(profile.cake_day) }}</span>
              </div>
              <div class="detail-item" v-if="profile.location">
                <span class="detail-label font-mono">Location:</span>
                <span class="detail-value">{{ profile.location }}</span>
              </div>
              <div class="detail-item" v-if="profile.timezone">
                <span class="detail-label font-mono">Timezone:</span>
                <span class="detail-value">{{ profile.timezone }}</span>
              </div>
              <div class="detail-item" v-if="profile.created_at">
                <span class="detail-label font-mono">Created:</span>
                <span class="detail-value">{{ formatDate(profile.created_at) }}</span>
              </div>
              <div class="detail-item" v-if="profile.updated_at">
                <span class="detail-label font-mono">Updated:</span>
                <span class="detail-value">{{ formatDate(profile.updated_at) }}</span>
              </div>
            </div>
          </div>
          
          <!-- Interests -->
          <div v-if="profile.interests" class="profile-section">
            <h3 class="text-lg font-bold mb-4 font-mono text-accent">$ interests</h3>
            <div class="flex flex-wrap gap-2">
              <span 
                v-for="(interest, index) in parseInterests(profile.interests)" 
                :key="index"
                class="interest-tag"
              >
                {{ interest }}
              </span>
            </div>
          </div>
        </div>
      </div>
      
      <!-- Contact & Links Section -->
      <div class="profile-section mb-8">
        <h3 class="text-xl font-bold mb-6 font-mono text-accent">$ contact & links</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          <a v-if="profile.email" :href="`mailto:${profile.email}`" 
             class="contact-link group">
            <div class="flex items-center gap-3">
              <div class="contact-icon">✉</div>
              <div>
                <div class="contact-label">Email</div>
                <div class="contact-value">{{ profile.email }}</div>
              </div>
            </div>
            <div class="contact-arrow">→</div>
          </a>
          
          <a v-if="profile.website" :href="profile.website" target="_blank"
             class="contact-link group">
            <div class="flex items-center gap-3">
              <div class="contact-icon">🌐</div>
              <div>
                <div class="contact-label">Website</div>
                <div class="contact-value">{{ shortenUrl(profile.website) }}</div>
              </div>
            </div>
            <div class="contact-arrow">↗</div>
          </a>
          
          <a v-if="profile.github" :href="profile.github" target="_blank"
             class="contact-link group">
            <div class="flex items-center gap-3">
              <div class="contact-icon"><span class="text-lg">🐙</span></div>
              <div>
                <div class="contact-label">GitHub</div>
                <div class="contact-value">{{ extractGitHubUsername(profile.github) }}</div>
              </div>
            </div>
            <div class="contact-arrow">↗</div>
          </a>
          
          <a v-if="profile.bluesky" :href="`https://bsky.app/profile/${profile.bluesky.replace('@', '')}`" 
             target="_blank" class="contact-link group">
            <div class="flex items-center gap-3">
              <div class="contact-icon text-sky-500">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 10.8c-1.087-2.114-4.046-6.053-6.798-7.995C2.566.944 1.561 1.266.902 1.565.139 1.908 0 3.08 0 3.768c0 .69.378 5.65.624 6.479.815 2.736 3.713 3.66 6.383 3.364.136-.02.275-.039.415-.056-.138.022-.278.04-.417.06-3.241.428-5.268 2.37-5.268 5.14 0 3.064 4.295 5.146 9.663 5.146 5.368 0 9.663-2.082 9.663-5.146 0-2.77-2.027-4.712-5.268-5.14-.139-.02-.279-.038-.417-.06.14.017.279.036.415.056 2.67.296 5.568-.628 6.383-3.364.246-.828.624-5.79.624-6.478 0-.69-.139-1.86-.902-2.206-.659-.298-1.664-.62-4.3 1.24C16.046 4.747 13.087 8.686 12 10.8z"/>
                </svg>
              </div>
              <div>
                <div class="contact-label">Bluesky</div>
                <div class="contact-value">{{ profile.bluesky }}</div>
              </div>
            </div>
            <div class="contact-arrow">↗</div>
          </a>
          
          <a v-if="profile.rss_feed" :href="profile.rss_feed" target="_blank"
             class="contact-link group">
            <div class="flex items-center gap-3">
              <div class="contact-icon">📡</div>
              <div>
                <div class="contact-label">RSS Feed</div>
                <div class="contact-value">rss.xml</div>
              </div>
            </div>
            <div class="contact-arrow">↗</div>
          </a>
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
    
    <!-- Site Description Footer -->
    <div class="mt-12 p-6 rounded-xl bg-background-secondary/30 border border-accent/10">
      <h3 class="text-xl font-bold mb-4 font-mono text-accent">
        $ about this site
      </h3>
      <p class="text-text-primary mb-4 font-mono">
        > lilthinaparka.i2p - personal blog-site + art gallery + portfolio + whatever i need it to be
      </p>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm">
        <div>
          <h4 class="font-bold mb-2 text-accent font-mono">Frontend</h4>
          <ul class="space-y-1 text-text-secondary font-mono">
            <li>• Deno</li>
            <li>• Nuxt 4</li>
            <li>• Pinia</li>
            <li>• TailwindCSS v4</li>
            <li>• TypeScript</li>
          </ul>
        </div>
        <div>
          <h4 class="font-bold mb-2 text-accent font-mono">Backend</h4>
          <ul class="space-y-1 text-text-secondary font-mono">
            <li>• Go</li>
            <li>• GORM</li>
            <li>• Echo Framework v4</li>
            <li>• SQLite</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Donations from './Donations.vue'
import type { Profile } from '../../../types'

const props = defineProps<{
  profile?: Profile
}>()

// Fallback for broken images
const handleImageError = (event: Event) => {
  const img = event.target as HTMLImageElement
  img.src = 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAwIiBoZWlnaHQ9IjIwMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cmVjdCB3aWR0aD0iMjAwIiBoZWlnaHQ9IjIwMCIgZmlsbD0iIzhiOGJlOSIvPjx0ZXh0IHg9IjUwJSIgeT0iNTAlIiBkb21pbmFudC1iYXNlbGluZT0ibWlkZGxlIiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmb250LWZhbWlseT0ibW9ub3NwYWNlIiBmb250LXNpemU9IjQ4IiBmaWxsPSJ3aGl0ZSI+TDwvdGV4dD48L3N2Zz4='
}

// Helper functions
const formatDate = (dateString: string): string => {
  if (!dateString) return 'Unknown'
  try {
    const date = new Date(dateString)
    return date.toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric'
    })
  } catch {
    return dateString
  }
}

const parseInterests = (interestsString: string): string[] => {
  if (!interestsString) return []
  // Split by comma or semicolon and trim each item
  return interestsString.split(/[,;]/).map(interest => interest.trim()).filter(Boolean)
}

const shortenUrl = (url: string): string => {
  if (!url) return ''
  try {
    const urlObj = new URL(url)
    return urlObj.hostname.replace('www.', '')
  } catch {
    return url.length > 30 ? url.slice(0, 30) + '...' : url
  }
}

const extractGitHubUsername = (url: string): string => {
  if (!url) return ''
  const match = url.match(/github\.com\/([^\/]+)/)
  return match ? match[0] : url.replace('https://github.com/', '')
}
</script>

<style scoped>
.profile-page {
  @apply max-w-6xl mx-auto p-4 md:p-8;
}

.terminal-header {
  @apply bg-background-secondary rounded-t-xl p-4 border-b border-accent/20 
         flex items-center justify-between mb-6 backdrop-blur-sm;
}

.terminal-buttons {
  @apply flex gap-2;
}

.terminal-buttons > div {
  @apply w-3 h-3 rounded-full transition-all duration-200;
}

.terminal-buttons .close { 
  @apply bg-red-500 hover:bg-red-600; 
}
.terminal-buttons .minimize { 
  @apply bg-yellow-500 hover:bg-yellow-600; 
}
.terminal-buttons .maximize { 
  @apply bg-green-500 hover:bg-green-600; 
}

.terminal-title {
  @apply text-text-secondary tracking-wider;
}

.profile-content {
  @apply p-6 rounded-b-xl bg-background/50 border border-accent/10 
         backdrop-blur-sm shadow-2xl;
}

/* Profile Picture Section */
.profile-pic-container {
  @apply relative;
}

.profile-banner {
  @apply h-32 overflow-hidden rounded-t-2xl bg-gradient-to-r from-accent/20 to-purple-500/20;
}

.banner-image {
  @apply w-full h-full object-cover opacity-40;
}

.profile-pic-wrapper {
  @apply relative w-48 h-48 mx-auto overflow-hidden rounded-2xl 
         border-4 border-background bg-background-secondary shadow-2xl 
         -mt-24 z-10;
}

.profile-pic-wrapper.with-banner {
  @apply -mt-12;
}

.profile-pic {
  @apply w-full h-full object-cover transition-all duration-500 
         hover:scale-110 hover:rotate-3;
}

.online-indicator {
  @apply absolute bottom-4 right-4 w-4 h-4 rounded-full bg-green-500 
         border-2 border-background shadow-lg;
}

/* Profile Section */
.profile-section {
  @apply p-6 rounded-xl bg-background-secondary/50 border border-accent/10 
         transition-all duration-300 hover:border-accent/20;
}

/* Username Badge */
.username-badge {
  @apply px-3 py-1 rounded-lg bg-accent/10 text-accent border 
         border-accent/20 text-sm font-bold;
}

/* Detail Items */
.detail-item {
  @apply flex items-center justify-between p-3 rounded-lg bg-background/30 
         hover:bg-background/50 transition-colors;
}

.detail-label {
  @apply text-accent font-bold;
}

.detail-value {
  @apply text-text-primary font-mono;
}

/* Interest Tags */
.interest-tag {
  @apply px-3 py-1.5 rounded-full bg-accent/10 text-accent text-sm 
         font-mono border border-accent/20 hover:bg-accent/20 
         hover:border-accent/30 transition-all duration-200 cursor-default;
}

/* Contact Links */
.contact-link {
  @apply flex items-center justify-between p-4 rounded-xl bg-background/30 
         border border-accent/5 hover:border-accent/30 hover:bg-background/50 
         transition-all duration-300 hover:-translate-y-1 hover:shadow-lg;
}

.contact-icon {
  @apply text-2xl opacity-80 group-hover:opacity-100 group-hover:scale-110 
         transition-all duration-300;
}

.contact-label {
  @apply text-xs text-text-secondary font-mono;
}

.contact-value {
  @apply text-text-primary font-mono text-sm truncate;
}

.contact-arrow {
  @apply text-accent opacity-0 group-hover:opacity-100 transform 
         group-hover:translate-x-1 transition-all duration-300;
}

/* Stat Boxes */
.stat-box {
  @apply p-3 rounded-lg bg-background/30 border border-accent/5 
         hover:border-accent/20 transition-all duration-300;
}

.stat-number {
  @apply text-2xl font-bold text-accent font-mono;
}

.stat-label {
  @apply text-text-secondary mt-1;
}

/* Custom scrollbar */
:deep() {
  ::-webkit-scrollbar {
    width: 10px;
  }
  
  ::-webkit-scrollbar-track {
    @apply bg-background-secondary;
  }
  
  ::-webkit-scrollbar-thumb {
    @apply bg-accent/30 rounded-full hover:bg-accent/50 transition-colors;
  }
}
</style>