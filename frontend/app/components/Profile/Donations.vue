<!-- frontend/app/components/Profile/Donations.vue -->
<template>
  <div class="donations-section">
    <h3 class="text-xl font-bold mb-4 font-mono text-accent">
      $ crypto donations
    </h3>
    <p class="text-text-secondary text-sm mb-6 font-mono">
      > support my work through cryptocurrency donations
    </p>
    
    <div class="space-y-4">
      <!-- Bitcoin -->
      <div v-if="bitcoin" class="donation-card">
        <div class="flex items-center justify-between mb-3">
          <div class="flex items-center gap-3">
            <div class="crypto-icon">₿</div>
            <div>
              <div class="crypto-name">Bitcoin (BTC)</div>
              <div class="crypto-network">Main Network</div>
            </div>
          </div>
          <button 
            @click="copyAddress(bitcoin, 'Bitcoin')" 
            class="copy-button"
            :class="{ 'copied': copiedCrypto === 'Bitcoin' }"
          >
            {{ copiedCrypto === 'Bitcoin' ? '✓ Copied' : 'Copy' }}
          </button>
        </div>
        <div class="crypto-address-container">
          <code class="crypto-address">{{ bitcoin }}</code>
        </div>
      </div>

      <!-- Ethereum -->
      <div v-if="ethereum" class="donation-card">
        <div class="flex items-center justify-between mb-3">
          <div class="flex items-center gap-3">
            <div class="crypto-icon">Ξ</div>
            <div>
              <div class="crypto-name">Ethereum (ETH)</div>
              <div class="crypto-network">ERC-20</div>
            </div>
          </div>
          <button 
            @click="copyAddress(ethereum, 'Ethereum')" 
            class="copy-button"
            :class="{ 'copied': copiedCrypto === 'Ethereum' }"
          >
            {{ copiedCrypto === 'Ethereum' ? '✓ Copied' : 'Copy' }}
          </button>
        </div>
        <div class="crypto-address-container">
          <code class="crypto-address">{{ ethereum }}</code>
        </div>
      </div>

      <!-- Solana -->
      <div v-if="solana" class="donation-card">
        <div class="flex items-center justify-between mb-3">
          <div class="flex items-center gap-3">
            <div class="crypto-icon">◎</div>
            <div>
              <div class="crypto-name">Solana (SOL)</div>
              <div class="crypto-network">Solana Network</div>
            </div>
          </div>
          <button 
            @click="copyAddress(solana, 'Solana')" 
            class="copy-button"
            :class="{ 'copied': copiedCrypto === 'Solana' }"
          >
            {{ copiedCrypto === 'Solana' ? '✓ Copied' : 'Copy' }}
          </button>
        </div>
        <div class="crypto-address-container">
          <code class="crypto-address">{{ solana }}</code>
        </div>
      </div>

      <!-- Monero -->
      <div v-if="monero" class="donation-card">
        <div class="flex items-center justify-between mb-3">
          <div class="flex items-center gap-3">
            <div class="crypto-icon">ɱ</div>
            <div>
              <div class="crypto-name">Monero (XMR)</div>
              <div class="crypto-network">Private & Untraceable</div>
            </div>
          </div>
          <button 
            @click="copyAddress(monero, 'Monero')" 
            class="copy-button"
            :class="{ 'copied': copiedCrypto === 'Monero' }"
          >
            {{ copiedCrypto === 'Monero' ? '✓ Copied' : 'Copy' }}
          </button>
        </div>
        <div class="crypto-address-container">
          <code class="crypto-address">{{ monero }}</code>
        </div>
      </div>
    </div>

    <!-- Donation Info -->
    <div class="mt-6 p-4 rounded-lg bg-accent/5 border border-accent/10">
      <p class="text-text-secondary text-sm font-mono">
        <span class="text-accent">💡</span> All donations are appreciated and help support this site and my projects.
        These addresses are verified and secure.
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps<{
  bitcoin?: string
  ethereum?: string
  solana?: string
  monero?: string
}>()

const copiedCrypto = ref<string | null>(null)

const copyAddress = async (address: string, cryptoName: string) => {
  try {
    await navigator.clipboard.writeText(address)
    copiedCrypto.value = cryptoName
    
    setTimeout(() => {
      copiedCrypto.value = null
    }, 2000)
  } catch (err) {
    console.error('Failed to copy address:', err)
    alert('Failed to copy address. Please copy manually.')
  }
}
</script>

<style scoped>
.donations-section {
  @apply p-6 rounded-xl bg-background-secondary/50 border border-accent/10;
}

.donation-card {
  @apply p-4 rounded-lg bg-background/30 border border-accent/5 
         transition-all duration-300 hover:border-accent/20;
}

.crypto-icon {
  @apply w-12 h-12 rounded-full bg-accent/10 flex items-center justify-center
         text-2xl font-bold text-accent;
}

.crypto-name {
  @apply text-text-primary font-bold font-mono;
}

.crypto-network {
  @apply text-text-secondary text-xs font-mono;
}

.copy-button {
  @apply px-4 py-2 rounded-lg border border-accent/20 text-text-secondary
         hover:text-accent hover:border-accent/40 transition-all duration-200
         font-mono text-sm;
}

.copy-button.copied {
  @apply bg-accent/10 text-accent border-accent/40;
}

.crypto-address-container {
  @apply bg-background-secondary rounded-lg p-3 overflow-x-auto;
}

.crypto-address {
  @apply text-text-secondary text-xs font-mono break-all;
}

/* Scrollbar styling */
.crypto-address-container::-webkit-scrollbar {
  height: 6px;
}

.crypto-address-container::-webkit-scrollbar-track {
  @apply bg-background;
}

.crypto-address-container::-webkit-scrollbar-thumb {
  @apply bg-accent/30 rounded-full hover:bg-accent/50;
}
</style>