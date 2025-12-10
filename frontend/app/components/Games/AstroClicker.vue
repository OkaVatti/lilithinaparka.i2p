<!-- frontend/app/components/Games/AstroClicker.vue -->
<template>
  <div class="astro-clicker">
    <div class="game-header">
      <h2 class="text-2xl font-bold font-mono text-accent mb-2">AstroClicker</h2>
      <p class="text-text-secondary text-sm mb-4">Click to harvest cosmic energy!</p>
    </div>
    
    <div class="game-stats-grid">
      <div class="stat-card">
        <div class="stat-value">{{ formatNumber(energy) }}</div>
        <div class="stat-label">Energy</div>
      </div>
      <div class="stat-card">
        <div class="stat-value">{{ formatNumber(energyPerSecond) }}/s</div>
        <div class="stat-label">Per Second</div>
      </div>
      <div class="stat-card">
        <div class="stat-value">{{ totalClicks }}</div>
        <div class="stat-label">Total Clicks</div>
      </div>
      <div class="stat-card">
        <div class="stat-value">{{ prestige }}</div>
        <div class="stat-label">Prestige</div>
      </div>
    </div>
    
    <div class="click-area">
      <button 
        @click="handleClick" 
        class="planet-button"
        :class="{ 'clicked': isClicked }"
      >
        <div class="planet">
          <span class="planet-emoji">🪐</span>
          <div class="click-value">+{{ clickPower }}</div>
        </div>
      </button>
      
      <div 
        v-for="particle in particles" 
        :key="particle.id"
        class="particle"
        :style="{
          left: particle.x + 'px',
          top: particle.y + 'px',
          opacity: particle.opacity
        }"
      >
        +{{ particle.value }}
      </div>
    </div>
    
    <div class="upgrades-section">
      <h3 class="text-xl font-bold font-mono text-accent mb-4">Upgrades</h3>
      
      <div class="upgrades-grid">
        <div 
          v-for="upgrade in upgrades" 
          :key="upgrade.id"
          class="upgrade-card"
          :class="{ 'affordable': energy >= upgrade.cost, 'disabled': energy < upgrade.cost }"
          @click="buyUpgrade(upgrade)"
        >
          <div class="upgrade-icon">{{ upgrade.icon }}</div>
          <div class="upgrade-info">
            <div class="upgrade-name">{{ upgrade.name }}</div>
            <div class="upgrade-level">Level {{ upgrade.level }}</div>
            <div class="upgrade-effect">{{ upgrade.effect }}</div>
            <div class="upgrade-cost">Cost: {{ formatNumber(upgrade.cost) }}</div>
          </div>
        </div>
      </div>
    </div>
    
    <div class="prestige-section" v-if="energy >= prestigeThreshold">
      <button @click="doPrestige" class="prestige-button">
        ✨ Prestige (Reset for {{ prestigeBonus }}% bonus)
      </button>
    </div>
    
    <div class="game-actions">
      <button @click="saveGame" class="action-button">💾 Save</button>
      <button @click="loadGame" class="action-button">📂 Load</button>
      <button @click="resetGame" class="action-button">🔄 Reset</button>
    </div>
  </div>
</template>


<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'

interface Upgrade {
  id: string
  name: string
  icon: string
  effect: string
  baseCost: number
  cost: number
  level: number
  energyPerSecond: number
  multiplier: number
}

interface Particle {
  id: number
  x: number
  y: number
  value: number
  opacity: number
}

const energy = ref(0)
const totalClicks = ref(0)
const clickPower = ref(1)
const prestige = ref(0)
const isClicked = ref(false)
const particles = ref<Particle[]>([])

let particleId = 0
let gameInterval: number | null = null

const upgrades = ref<Upgrade[]>([
  {
    id: 'cursor',
    name: 'Cosmic Cursor',
    icon: '👆',
    effect: '+0.1 energy/s',
    baseCost: 10,
    cost: 10,
    level: 0,
    energyPerSecond: 0.1,
    multiplier: 1.15
  },
  {
    id: 'asteroid',
    name: 'Asteroid Miner',
    icon: '☄️',
    effect: '+1 energy/s',
    baseCost: 100,
    cost: 100,
    level: 0,
    energyPerSecond: 1,
    multiplier: 1.15
  },
  {
    id: 'satellite',
    name: 'Solar Satellite',
    icon: '🛰️',
    effect: '+5 energy/s',
    baseCost: 500,
    cost: 500,
    level: 0,
    energyPerSecond: 5,
    multiplier: 1.15
  },
  {
    id: 'station',
    name: 'Space Station',
    icon: '🏗️',
    effect: '+25 energy/s',
    baseCost: 2500,
    cost: 2500,
    level: 0,
    energyPerSecond: 25,
    multiplier: 1.15
  },
  {
    id: 'colony',
    name: 'Moon Colony',
    icon: '🌙',
    effect: '+100 energy/s',
    baseCost: 10000,
    cost: 10000,
    level: 0,
    energyPerSecond: 100,
    multiplier: 1.15
  },
  {
    id: 'dyson',
    name: 'Dyson Sphere',
    icon: '⭕',
    effect: '+500 energy/s',
    baseCost: 50000,
    cost: 50000,
    level: 0,
    energyPerSecond: 500,
    multiplier: 1.15
  }
])

const energyPerSecond = computed(() => {
  let total = 0
  upgrades.value.forEach(upgrade => {
    total += upgrade.energyPerSecond * upgrade.level
  })
  return total * (1 + prestige.value * 0.1)
})

const prestigeThreshold = 100000
const prestigeBonus = computed(() => Math.floor(prestige.value * 10 + 10))

onMounted(() => {
  loadGame()
  startGameLoop()
})

onUnmounted(() => {
  stopGameLoop()
})

const startGameLoop = () => {
  gameInterval = window.setInterval(() => {
    energy.value += energyPerSecond.value / 10
  }, 100)
}

const stopGameLoop = () => {
  if (gameInterval) {
    clearInterval(gameInterval)
  }
}

const handleClick = (event: MouseEvent) => {
  const clickValue = clickPower.value * (1 + prestige.value * 0.1)
  energy.value += clickValue
  totalClicks.value++
  
  isClicked.value = true
  setTimeout(() => {
    isClicked.value = false
  }, 100)
  
  createParticle(event.clientX, event.clientY, clickValue)
}

const createParticle = (x: number, y: number, value: number) => {
  const particle: Particle = {
    id: particleId++,
    x,
    y,
    value: Math.round(value),
    opacity: 1
  }
  
  particles.value.push(particle)
  
  const animation = setInterval(() => {
    const index = particles.value.findIndex(p => p.id === particle.id)
    if (index !== -1) {
      particles.value[index].y -= 2
      particles.value[index].opacity -= 0.02
      
      if (particles.value[index].opacity <= 0) {
        particles.value.splice(index, 1)
        clearInterval(animation)
      }
    }
  }, 16)
}

const buyUpgrade = (upgrade: Upgrade) => {
  if (energy.value >= upgrade.cost) {
    energy.value -= upgrade.cost
    upgrade.level++
    upgrade.cost = Math.floor(upgrade.baseCost * Math.pow(upgrade.multiplier, upgrade.level))
  }
}

const doPrestige = () => {
  if (confirm(`Prestige will reset all progress but give you a ${prestigeBonus.value}% permanent bonus. Continue?`)) {
    prestige.value++
    energy.value = 0
    totalClicks.value = 0
    upgrades.value.forEach(upgrade => {
      upgrade.level = 0
      upgrade.cost = upgrade.baseCost
    })
  }
}

const formatNumber = (num: number): string => {
  if (num < 1000) return Math.floor(num).toString()
  if (num < 1000000) return (num / 1000).toFixed(1) + 'K'
  if (num < 1000000000) return (num / 1000000).toFixed(1) + 'M'
  return (num / 1000000000).toFixed(1) + 'B'
}

const saveGame = () => {
  const saveData = {
    energy: energy.value,
    totalClicks: totalClicks.value,
    clickPower: clickPower.value,
    prestige: prestige.value,
    upgrades: upgrades.value.map(u => ({
      id: u.id,
      level: u.level,
      cost: u.cost
    }))
  }
  
  localStorage.setItem('astroclicker_save', JSON.stringify(saveData))
  alert('Game saved!')
}

const loadGame = () => {
  const saved = localStorage.getItem('astroclicker_save')
  if (saved) {
    try {
      const data = JSON.parse(saved)
      energy.value = data.energy || 0
      totalClicks.value = data.totalClicks || 0
      clickPower.value = data.clickPower || 1
      prestige.value = data.prestige || 0
      
      if (data.upgrades) {
        data.upgrades.forEach((savedUpgrade: any) => {
          const upgrade = upgrades.value.find(u => u.id === savedUpgrade.id)
          if (upgrade) {
            upgrade.level = savedUpgrade.level
            upgrade.cost = savedUpgrade.cost
          }
        })
      }
    } catch (e) {
      console.error('Failed to load save:', e)
    }
  }
}

const resetGame = () => {
  if (confirm('Are you sure you want to reset all progress?')) {
    localStorage.removeItem('astroclicker_save')
    energy.value = 0
    totalClicks.value = 0
    clickPower.value = 1
    prestige.value = 0
    upgrades.value.forEach(upgrade => {
      upgrade.level = 0
      upgrade.cost = upgrade.baseCost
    })
  }
}
</script>
<style scoped>
.astro-clicker {
  @apply p-6 bg-bg-secondary rounded-xl border border-accent/20;
}

.game-header {
  @apply mb-6 text-center;
}

.game-stats-grid {
  @apply grid grid-cols-2 md:grid-cols-4 gap-4 mb-6;
}

.stat-card {
  @apply p-4 bg-bg-primary rounded-lg border border-accent/10 text-center;
}

.stat-value {
  @apply text-2xl font-bold text-accent font-mono;
}

.stat-label {
  @apply text-text-secondary text-sm mt-1;
}

.click-area {
  @apply relative flex justify-center items-center h-64 mb-6 bg-bg-primary rounded-lg overflow-hidden;
}

.planet-button {
  @apply relative transform transition-transform duration-100;
}

.planet-button.clicked {
  @apply scale-95;
}

.planet {
  @apply relative;
}

.planet-emoji {
  @apply text-8xl cursor-pointer hover:scale-110 transition-transform duration-200;
}

.click-value {
  @apply absolute inset-x-0 bottom-0 text-center text-accent font-bold font-mono;
}

.particle {
  @apply absolute text-accent font-bold font-mono pointer-events-none;
  animation: float-up 2s ease-out;
}

@keyframes float-up {
  from {
    transform: translateY(0);
  }
  to {
    transform: translateY(-50px);
  }
}

.upgrades-section {
  @apply mb-6;
}

.upgrades-grid {
  @apply grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4;
}

.upgrade-card {
  @apply p-4 bg-bg-primary rounded-lg border border-accent/10 cursor-pointer
         transition-all duration-200 hover:border-accent/30;
}

.upgrade-card.affordable {
  @apply hover:bg-accent/5;
}

.upgrade-card.disabled {
  @apply opacity-50 cursor-not-allowed;
}

.upgrade-icon {
  @apply text-4xl mb-2;
}

.upgrade-name {
  @apply font-bold text-text-primary;
}

.upgrade-level {
  @apply text-accent text-sm font-mono;
}

.upgrade-effect {
  @apply text-text-secondary text-sm;
}

.upgrade-cost {
  @apply text-accent font-mono text-sm mt-2;
}

.prestige-section {
  @apply mb-6 text-center;
}

.prestige-button {
  @apply px-6 py-3 bg-gradient-to-r from-purple-500 to-accent text-white 
         rounded-lg font-mono font-bold hover:scale-105 transition-transform duration-200;
}

.game-actions {
  @apply flex justify-center gap-4;
}

.action-button {
  @apply px-4 py-2 bg-bg-primary text-text-secondary rounded-lg font-mono
         hover:text-accent hover:bg-accent/10 transition-all duration-200;
}
</style>
