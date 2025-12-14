<!-- app/pages/profile.vue -->
<template>
  <div class="profile-page">
    <div class="container">
      <div v-if="profileStore.loading" class="loading window">
        <div class="window-body">
          <LoadingIcon name="loader" size="24" class="spin" />
          <span>Loading profile...</span>
        </div>
      </div>

      <div v-else-if="profileStore.error" class="error window">
        <div class="window-body">
          <AlertCircleIcon name="alert-circle" size="24" />
          <h3>Error Loading Profile</h3>
          <p>{{ profileStore.error }}</p>
          <button @click="retry" class="btn btn-primary">
            <RetryIcon name="refresh-cw" size="16" />
            <span>Retry</span>
          </button>
        </div>
      </div>

      <div v-else-if="profile" class="profile-content">
        <!-- Profile Header -->
        <header class="profile-header window">
          <div class="title-bar">
            <div class="title-bar-text">
              <UserIcon name="user" size="26" />
              <span>@lilithinaparka.i2p</span>
            </div>
          </div>
          <div class="window-body">
            <div class="profile-header-content">
              <div class="profile-avatar">
                <img 
                  v-if="profile.pic" 
                  :src="profile.pic" 
                  :alt="profile.name" 
                  class="avatar-img"
                  @error="handleImageError"
                />
                <div v-else class="avatar-fallback">
                  <img class="avatar-img" src="../../../public/favicon.svg" alt={{ profile.name }}>
                </div>
              </div>
              
              <div class="profile-basic">
                <h1 class="profile-name">{{ profile.name }}</h1>
                <p v-if="profile.bio" class="profile-bio">{{ profile.bio }}</p>
                
                <div class="profile-meta">
                  <div v-if="profile.location" class="meta-item">
                    <FeatherIcon name="map-pin" size="14" />
                    <span>{{ profile.location }}</span>
                  </div>
                  <div v-if="profile.cake_day" class="meta-item">
                    <FeatherIcon name="calendar" size="14" />
                    <span>Born {{ profile.cake_day }}</span>
                  </div>
                  <div v-if="profile.timezone" class="meta-item">
                    <FeatherIcon name="clock" size="14" />
                    <span>{{ profile.timezone }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </header>

        <!-- Main Profile Sections -->
        <div class="profile-grid">
          <!-- Left Column: Personal Info -->
          <div class="profile-left">
            <!-- Interests -->
            <div v-if="interests.length > 0" class="profile-section window">
              <div class="title-bar">
                <div class="title-bar-text">
                  <ActivityIcon name="heart" size="14" />
                  <span>Interests</span>
                </div>
              </div>
              <div class="window-body">
                <div class="interests-list">
                  <span 
                    v-for="interest in interests" 
                    :key="interest" 
                    class="interest-tag"
                  >
                    {{ interest }}
                  </span>
                </div>
              </div>
            </div>

            <!-- Contact Info -->
            <div class="profile-section window">
              <div class="title-bar">
                <div class="title-bar-text">
                  <UsersIcon name="mail" size="14" />
                  <span>Contact</span>
                </div>
              </div>
              <div class="window-body">
                <div class="contact-list">
                  <div v-if="profile.email" class="contact-item">
                    <MailIcon name="mail" size="16" />
                    <a :href="`mailto:${profile.email}`" class="contact-link">
                      {{ profile.email }}
                    </a>
                  </div>
                  
                  <div v-if="profile.website" class="contact-item">
                    <GlobeIcon name="globe" size="16" />
                    <a 
                      :href="profile.website" 
                      target="_blank" 
                      rel="noopener noreferrer"
                      class="contact-link"
                    >
                      {{ profile.website.replace(/^https?:\/\//, '') }}
                    </a>
                  </div>
                  
                  <div v-if="profile.rss_feed" class="contact-item">
                    <RssIcon name="rss" size="16" />
                    <a 
                      :href="profile.rss_feed" 
                      target="_blank" 
                      rel="noopener noreferrer"
                      class="contact-link"
                    >
                      RSS Feed
                    </a>
                  </div>
                </div>
              </div>
            </div>

            <!-- Social Links -->
            <div v-if="hasSocialLinks" class="profile-section window">
              <div class="title-bar">
                <div class="title-bar-text">
                  <SmartphoneIcon name="share-2" size="14" />
                  <span>Social</span>
                </div>
              </div>
              <div class="window-body">
                <div class="social-links">
                  <a 
                    v-if="profile.github"
                    :href="profile.github"
                    target="_blank"
                    rel="noopener noreferrer"
                    class="social-link"
                    title="GitHub"
                  >
                    <GithubIcon name="github" size="20" />
                  </a>
                  
                  <a 
                    v-if="profile.bluesky"
                    :href="profile.bluesky"
                    target="_blank"
                    rel="noopener noreferrer"
                    class="social-link"
                    title="BlueSky"
                  >
                    <WindIcon name="message-circle" size="20" />
                  </a>
                </div>
              </div>
            </div>
          </div>

          <!-- Right Column: BlueSky Info & Donations -->
          <div class="profile-right">
            <!-- BlueSky Profile -->
            <div v-if="hasBskyProfile" class="profile-section window">
              <div class="title-bar">
                <div class="title-bar-text">
                  <FeatherIcon name="message-circle" size="14" />
                  <span>BlueSky</span>
                </div>
              </div>
              <div class="window-body">
                <div class="bsky-profile">
                  <div v-if="profile.bsky_display_name" class="bsky-header">
                    <div class="bsky-avatar" v-if="profile.bsky_avatar">
                      <img 
                        :src="profile.bsky_avatar" 
                        alt="BlueSky Avatar"
                        class="bsky-avatar-img"
                      />
                    </div>
                    <div class="bsky-info">
                      <h3>{{ profile.bsky_display_name }}</h3>
                      <div v-if="profile.bsky_description" class="bsky-bio">
                        {{ profile.bsky_description }}
                      </div>
                    </div>
                  </div>
                  
                  <div class="bsky-stats">
                    <div v-if="profile.bsky_posts_count !== undefined" class="bsky-stat">
                      <span class="stat-value">{{ profile.bsky_posts_count }}</span>
                      <span class="stat-label">Posts</span>
                    </div>
                    <div v-if="profile.bsky_followers_count !== undefined" class="bsky-stat">
                      <span class="stat-value">{{ profile.bsky_followers_count }}</span>
                      <span class="stat-label">Followers</span>
                    </div>
                    <div v-if="profile.bsky_follows_count !== undefined" class="bsky-stat">
                      <span class="stat-value">{{ profile.bsky_follows_count }}</span>
                      <span class="stat-label">Following</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Donation Addresses -->
            <div v-if="hasDonations" class="profile-section window">
              <div class="title-bar">
                <div class="title-bar-text">
                  <FeatherIcon name="gift" size="14" />
                  <span>Support</span>
                </div>
              </div>
              <div class="window-body">
                <div class="donations">
                  <p class="donations-note">
                    If you enjoy my content and would like to support my work:
                  </p>
                  
                  <div class="crypto-addresses">
                    <div 
                      v-if="profile.bitcoin_donation_addr" 
                      class="crypto-address"
                      @click="copyToClipboard(profile.bitcoin_donation_addr)"
                    >
                      <div class="crypto-header">
                        <FeatherIcon name="bitcoin" size="16" />
                        <span class="crypto-name">Bitcoin</span>
                        <button 
                          class="copy-btn"
                          title="Copy to clipboard"
                        >
                          <FeatherIcon name="copy" size="12" />
                        </button>
                      </div>
                      <code class="crypto-addr">
                        {{ profile.bitcoin_donation_addr }}
                      </code>
                    </div>
                    
                    <div 
                      v-if="profile.ethereum_donation_addr" 
                      class="crypto-address"
                      @click="copyToClipboard(profile.ethereum_donation_addr)"
                    >
                      <div class="crypto-header">
                        <FeatherIcon name="hexagon" size="16" />
                        <span class="crypto-name">Ethereum</span>
                        <button 
                          class="copy-btn"
                          title="Copy to clipboard"
                        >
                          <FeatherIcon name="copy" size="12" />
                        </button>
                      </div>
                      <code class="crypto-addr">
                        {{ profile.ethereum_donation_addr }}
                      </code>
                    </div>
                    
                    <div 
                      v-if="profile.solana_donation_addr" 
                      class="crypto-address"
                      @click="copyToClipboard(profile.solana_donation_addr)"
                    >
                      <div class="crypto-header">
                        <FeatherIcon name="circle" size="16" />
                        <span class="crypto-name">Solana</span>
                        <button 
                          class="copy-btn"
                          title="Copy to clipboard"
                        >
                          <FeatherIcon name="copy" size="12" />
                        </button>
                      </div>
                      <code class="crypto-addr">
                        {{ profile.solana_donation_addr }}
                      </code>
                    </div>
                    
                    <div 
                      v-if="profile.monero_donation_addr" 
                      class="crypto-address"
                      @click="copyToClipboard(profile.monero_donation_addr)"
                    >
                      <div class="crypto-header">
                        <FeatherIcon name="shield" size="16" />
                        <span class="crypto-name">Monero</span>
                        <button 
                          class="copy-btn"
                          title="Copy to clipboard"
                        >
                          <FeatherIcon name="copy" size="12" />
                        </button>
                      </div>
                      <code class="crypto-addr">
                        {{ profile.monero_donation_addr }}
                      </code>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useProfileStore } from '../../../stores/profile'
import type { Profile } from '../../../types'

const profileStore = useProfileStore()
const copyNotification = ref<string>('')

const profile = computed(() => profileStore.profile)

// Parse interests from JSON string
const interests = computed(() => {
  if (!profile.value?.interests) return []
  try {
    return JSON.parse(profile.value.interests)
  } catch {
    return []
  }
})

// Check for social links
const hasSocialLinks = computed(() => {
  return profile.value?.github || profile.value?.bluesky
})

// Check for BlueSky profile data
const hasBskyProfile = computed(() => {
  return profile.value?.bsky_display_name || 
         profile.value?.bsky_avatar || 
         profile.value?.bsky_description
})

// Check for donation addresses
const hasDonations = computed(() => {
  return profile.value?.bitcoin_donation_addr ||
         profile.value?.ethereum_donation_addr ||
         profile.value?.solana_donation_addr ||
         profile.value?.monero_donation_addr
})

// Get initials for avatar fallback
const getInitials = (name: string | undefined): string => {
  if (!name) return '?'
  return name
    .split(' ')
    .map(word => word[0])
    .join('')
    .toUpperCase()
    .substring(0, 2)
}

// Handle image loading errors
const handleImageError = (event: Event) => {
  const img = event.target as HTMLImageElement
  img.style.display = 'none'
  // The fallback will show via v-else
}

// Copy to clipboard function
const copyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text)
    copyNotification.value = 'Copied!'
    
    setTimeout(() => {
      copyNotification.value = ''
    }, 2000)
  } catch (err) {
    console.error('Failed to copy:', err)
    copyNotification.value = 'Failed to copy'
  }
}

// Retry loading profile
const retry = () => {
  profileStore.fetchProfile()
}

// Fetch profile on mount
onMounted(() => {
  if (!profile.value) {
    profileStore.fetchProfile()
  }
})

useHead({
  title: profile.value?.name ? `${profile.value.name} - Profile` : 'Profile',
  meta: [
    { name: 'description', content: profile.value?.bio || 'Site owner profile' }
  ]
})
</script>

<style scoped>
.profile-page {
  padding: 2rem 0;
  min-height: calc(100vh - 200px);
}

.profile-header {
  margin-bottom: 2rem;
}

.profile-header-content {
  display: flex;
  gap: 2rem;
  align-items: center;
  flex-wrap: wrap;
}

.profile-avatar {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  overflow: hidden;
  border: 2px solid var(--theme-border);
  background: rgba(0, 0, 0, 0.2);
  flex-shrink: 0;
}

.avatar-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar-fallback {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  font-weight: bold;
  color: var(--theme-primary);
  background: linear-gradient(135deg, var(--theme-primary) 0%, var(--theme-accent) 100%);
}

.profile-basic {
  flex: 1;
  min-width: 300px;
}

.profile-name {
  font-size: 2.5rem;
  margin: 0 0 0.5rem 0;
  color: var(--theme-primary);
}

.profile-bio {
  font-size: 1.1rem;
  line-height: 1.6;
  margin: 0 0 1rem 0;
  color: var(--theme-fg);
  opacity: 0.9;
}

.profile-meta {
  display: flex;
  gap: 1.5rem;
  flex-wrap: wrap;
  font-size: 0.9rem;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: var(--theme-fg);
  opacity: 0.7;
}

.profile-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
}

@media (max-width: 1100px) {
  .profile-grid {
    grid-template-columns: 1fr;
  }
}

.profile-section {
  margin-bottom: 1.5rem;
}

/* Interests */
.interests-list {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.interest-tag {
  padding: 0.25rem 0.75rem;
  background: rgba(139, 147, 233, 0.1);
  border: 1px solid var(--theme-primary);
  border-radius: 4px;
  font-size: 0.9rem;
  color: var(--theme-primary);
}

/* Contact Info */
.contact-list {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.contact-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.contact-link {
  color: var(--theme-accent);
  text-decoration: none;
  transition: color 0.2s;
  word-break: break-all;
}

.contact-link:hover {
  color: var(--theme-primary);
  text-decoration: underline;
}

/* Social Links */
.social-links {
  display: flex;
  gap: 1rem;
}

.social-link {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  border-radius: 4px;
  border: 1px solid var(--theme-border);
  color: var(--theme-fg);
  transition: all 0.2s;
}

.social-link:hover {
  border-color: var(--theme-primary);
  color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
}

/* BlueSky Profile */
.bsky-profile {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.bsky-header {
  display: flex;
  gap: 1rem;
  align-items: center;
}

.bsky-avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  overflow: hidden;
  border: 1px solid var(--theme-border);
}

.bsky-avatar-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.bsky-info h3 {
  margin: 0 0 0.5rem 0;
  font-size: 1.2rem;
  color: var(--theme-primary);
}

.bsky-bio {
  font-size: 0.9rem;
  color: var(--theme-fg);
  opacity: 0.8;
  line-height: 1.5;
}

.bsky-stats {
  display: flex;
  gap: 1.5rem;
  padding-top: 1rem;
  border-top: 1px solid var(--theme-border);
}

.bsky-stat {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.stat-value {
  font-size: 1.5rem;
  font-weight: bold;
  color: var(--theme-accent);
}

.stat-label {
  font-size: 0.8rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

/* Donations */
.donations-note {
  margin: 0 0 1.5rem 0;
  color: var(--theme-fg);
  opacity: 0.8;
  font-size: 0.95rem;
  line-height: 1.5;
}

.crypto-addresses {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.crypto-address {
  padding: 1rem;
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.2s;
}

.crypto-address:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.05);
}

.crypto-header {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.75rem;
  color: var(--theme-primary);
}

.crypto-name {
  font-weight: bold;
  flex: 1;
}

.copy-btn {
  padding: 0.25rem;
  background: transparent;
  border: 1px solid var(--theme-border);
  border-radius: 3px;
  color: var(--theme-fg);
  cursor: pointer;
  transition: all 0.2s;
}

.copy-btn:hover {
  border-color: var(--theme-primary);
  color: var(--theme-primary);
}

.crypto-addr {
  font-family: 'Courier New', monospace;
  font-size: 0.85rem;
  color: var(--theme-fg);
  opacity: 0.8;
  word-break: break-all;
  user-select: all;
}

/* Loading & Error States */
.loading, .error {
  text-align: center;
  padding: 3rem;
}

.loading svg.spin {
  animation: spin 1s linear infinite;
  margin-bottom: 1rem;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.error h3 {
  margin: 0 0 0.5rem 0;
  color: var(--theme-error);
}

/* Responsive */
@media (max-width: 768px) {
  .profile-header-content {
    flex-direction: column;
    text-align: center;
  }
  
  .profile-avatar {
    width: 100px;
    height: 100px;
  }
  
  .profile-name {
    font-size: 2rem;
  }
  
  .profile-meta {
    justify-content: center;
  }
  
  .profile-grid {
    gap: 1rem;
  }
}
</style>