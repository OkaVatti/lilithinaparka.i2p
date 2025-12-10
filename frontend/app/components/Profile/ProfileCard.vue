<template>
  <div v-if="profile" class="profile-card card">
    <div class="profile-header">
      <img
        v-if="profile.pic"
        :src="profile.pic"
        :alt="profile.name"
        class="profile-pic"
      />
      <div class="profile-info">
        <h1>{{ profile.name }}</h1>
        <p class="bio">{{ profile.bio }}</p>
      </div>
    </div>
    
    <div class="profile-details">
      <div v-if="profile.cake_day" class="detail-item">
        <FeatherIcon name="cake" size="18" />
        <span>Birthday: {{ profile.cake_day }}</span>
      </div>
      
      <div v-if="profile.location" class="detail-item">
        <FeatherIcon name="map-pin" size="18" />
        <span>{{ profile.location }}</span>
      </div>
      
      <div v-if="profile.timezone" class="detail-item">
        <FeatherIcon name="clock" size="18" />
        <span>{{ profile.timezone }}</span>
      </div>
      
      <div v-if="profile.email" class="detail-item">
        <FeatherIcon name="mail" size="18" />
        <a :href="`mailto:${profile.email}`">{{ profile.email }}</a>
      </div>
    </div>
    
    <div v-if="interests.length" class="interests">
      <h3>Interests</h3>
      <div class="interest-tags">
        <span v-for="interest in interests" :key="interest" class="interest-tag">
          {{ interest }}
        </span>
      </div>
    </div>
    
    <div class="social-links">
      <a v-if="profile.github" :href="profile.github" target="_blank" rel="noopener" class="social-link">
        <FeatherIcon name="github" size="20" />
        <span>GitHub</span>
      </a>
      
      <a v-if="profile.bluesky" :href="profile.bluesky" target="_blank" rel="noopener" class="social-link">
        <FeatherIcon name="cloud" size="20" />
        <span>BlueSky</span>
      </a>
      
      <a v-if="profile.rss_feed" :href="profile.rss_feed" class="social-link">
        <FeatherIcon name="rss" size="20" />
        <span>RSS Feed</span>
      </a>
    </div>
    
    <div v-if="hasDonationAddresses" class="donations">
      <h3>Support</h3>
      <div class="donation-addresses">
        <div v-if="profile.bitcoin_donation_addr" class="donation-item">
          <strong>Bitcoin:</strong>
          <code>{{ profile.bitcoin_donation_addr }}</code>
        </div>
        <div v-if="profile.ethereum_donation_addr" class="donation-item">
          <strong>Ethereum:</strong>
          <code>{{ profile.ethereum_donation_addr }}</code>
        </div>
        <div v-if="profile.solana_donation_addr" class="donation-item">
          <strong>Solana:</strong>
          <code>{{ profile.solana_donation_addr }}</code>
        </div>
        <div v-if="profile.monero_donation_addr" class="donation-item">
          <strong>Monero:</strong>
          <code>{{ profile.monero_donation_addr }}</code>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Profile } from '~~/types'

const props = defineProps<{
  profile: Profile
}>()

const interests = computed(() => {
  try {
    return JSON.parse(props.profile.interests || '[]')
  } catch {
    return []
  }
})

const hasDonationAddresses = computed(() => {
  return props.profile.bitcoin_donation_addr ||
         props.profile.ethereum_donation_addr ||
         props.profile.solana_donation_addr ||
         props.profile.monero_donation_addr
})
</script>

<style scoped>
.profile-card {
  max-width: 800px;
  margin: 0 auto;
}

.profile-header {
  display: flex;
  gap: 2rem;
  align-items: flex-start;
  margin-bottom: 2rem;
}

.profile-pic {
  width: 150px;
  height: 150px;
  border-radius: 50%;
  border: 3px solid var(--theme-primary);
  object-fit: cover;
}

.profile-info {
  flex: 1;
}

.profile-info h1 {
  margin: 0 0 0.5rem 0;
  color: var(--theme-primary);
}

.bio {
  margin: 0;
  font-size: 1.1rem;
  line-height: 1.6;
  color: var(--theme-fg);
  opacity: 0.9;
}

.profile-details {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  margin-bottom: 2rem;
  padding: 1rem;
  background: rgba(0, 0, 0, 0.2);
  border-radius: 4px;
}

.detail-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  color: var(--theme-fg);
}

.detail-item a {
  color: var(--theme-accent);
}

.interests {
  margin-bottom: 2rem;
}

.interests h3 {
  margin-bottom: 1rem;
  color: var(--theme-primary);
}

.interest-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.interest-tag {
  padding: 0.5rem 1rem;
  background: rgba(189, 147, 249, 0.1);
  border: 1px solid var(--theme-primary);
  border-radius: 4px;
  color: var(--theme-primary);
}

.social-links {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
  margin-bottom: 2rem;
}

.social-link {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1.5rem;
  background: rgba(189, 147, 249, 0.1);
  border: 2px solid var(--theme-primary);
  border-radius: 4px;
  color: var(--theme-primary);
  transition: all 0.2s;
}

.social-link:hover {
  background: rgba(189, 147, 249, 0.2);
  transform: translateY(-2px);
}

.donations h3 {
  margin-bottom: 1rem;
  color: var(--theme-primary);
}

.donation-addresses {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.donation-item {
  padding: 0.75rem;
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
}

.donation-item strong {
  display: block;
  margin-bottom: 0.5rem;
  color: var(--theme-accent);
}

.donation-item code {
  display: block;
  word-break: break-all;
  font-size: 0.85rem;
}

@media (max-width: 768px) {
  .profile-header {
    flex-direction: column;
    align-items: center;
    text-align: center;
  }
}
</style>