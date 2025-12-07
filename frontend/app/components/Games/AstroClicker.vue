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