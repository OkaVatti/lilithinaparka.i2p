<!-- app/pages/projects/games/index.vue -->
<template>
  <div class="games-page">
    <!-- Games Header -->
    <div class="games-header mb-8">
      <h1 class="text-4xl font-bold text-accent mb-4 font-mono">$ games</h1>
      <p class="text-text-secondary font-mono">
        > interactive games and experiments
      </p>
      
      <!-- Games Stats -->
      <div class="flex flex-wrap gap-4 mt-6">
        <div class="game-stat">
          <div class="stat-number">{{ totalGames }}</div>
          <div class="stat-label">Total Games</div>
        </div>
        <div class="game-stat">
          <div class="stat-number">{{ categories.length }}</div>
          <div class="stat-label">Categories</div>
        </div>
        <div class="game-stat">
          <div class="stat-number">{{ multiplayerGames }}</div>
          <div class="stat-label">Multiplayer</div>
        </div>
      </div>
    </div>
    
    <!-- Category Filter -->
    <div class="flex flex-wrap gap-2 mb-8">
      <button
        v-for="category in categories"
        :key="category"
        @click="setActiveCategory(category)"
        class="game-filter"
        :class="{ 'active': activeCategory === category }"
      >
        {{ category }}
      </button>
    </div>
    
    <!-- Games Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12">
      <GameCard
        v-for="game in filteredGames"
        :key="game.id"
        :game="game"
        :detailed="true"
      />
    </div>
    
    <!-- Game Categories Explanation -->
    <div class="mt-12 p-6 rounded-xl bg-bg-secondary/50 border border-accent/10">
      <h3 class="text-xl font-bold mb-4 font-mono text-accent">$ game categories</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div class="game-category">
          <h4 class="text-lg font-bold mb-2 text-accent font-mono">Idle / Clicker</h4>
          <p class="text-text-secondary text-sm">
            Games focused on incremental progress and resource management.
          </p>
          <ul class="mt-2 space-y-1 text-xs text-text-secondary font-mono">
            <li>• AstroClicker</li>
          </ul>
        </div>
        <div class="game-category">
          <h4 class="text-lg font-bold mb-2 text-accent font-mono">Arcade</h4>
          <p class="text-text-secondary text-sm">
            Fast-paced action games with simple mechanics and high replay value.
          </p>
          <ul class="mt-2 space-y-1 text-xs text-text-secondary font-mono">
            <li>• CatRunner</li>
            <li>• Solitaire</li>
          </ul>
        </div>
        <div class="game-category">
          <h4 class="text-lg font-bold mb-2 text-accent font-mono">Strategy</h4>
          <p class="text-text-secondary text-sm">
            Games requiring planning, tactics, and strategic thinking.
          </p>
          <ul class="mt-2 space-y-1 text-xs text-text-secondary font-mono">
            <li>• Chess</li>
            <li>• Sudoku</li>
            <li>• Bones</li>
          </ul>
        </div>
        <div class="game-category">
          <h4 class="text-lg font-bold mb-2 text-accent font-mono">Word</h4>
          <p class="text-text-secondary text-sm">
            Games focused on language, vocabulary, and word puzzles.
          </p>
          <ul class="mt-2 space-y-1 text-xs text-text-secondary font-mono">
            <li>• Wordle</li>
          </ul>
        </div>
        <div class="game-category">
          <h4 class="text-lg font-bold mb-2 text-accent font-mono">Card</h4>
          <p class="text-text-secondary text-sm">
            Games based on card decks, with elements of chance and strategy.
          </p>
          <ul class="mt-2 space-y-1 text-xs text-text-secondary font-mono">
            <li>• Bones</li>
          </ul>
        </div>
      </div>
    </div>
    
    <!-- Multiplayer Section -->
    <div class="mt-12 p-6 rounded-xl bg-bg-secondary/30 border border-accent/10">
      <h3 class="text-xl font-bold mb-4 font-mono text-accent">$ multiplayer games</h3>
      <p class="text-text-secondary mb-6 font-mono">
        > games you can play with friends (or against AI)
      </p>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div v-for="game in multiplayerGamesList" :key="game.id" class="multiplayer-game">
          <div class="flex items-start justify-between mb-2">
            <h4 class="text-lg font-bold text-accent">{{ game.name }}</h4>
            <span class="px-3 py-1 rounded-full bg-accent/10 text-accent text-xs font-mono">
              {{ game.players }}
            </span>
          </div>
          <p class="text-text-secondary text-sm mb-4">{{ game.description }}</p>
          <div class="flex flex-wrap gap-2">
            <span v-for="mode in game.modes" :key="mode" class="px-2 py-1 rounded bg-bg-secondary text-xs text-text-secondary">
              {{ mode }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import GameCard from '../../../components/Posts/GameCard.vue'
import { ref, computed } from 'vue'

// Sample games data
const games = [
  { id: 1, name: 'AstroClicker', description: 'Space-themed Cookie Clicker clone with planets, stars, and galaxies to unlock', category: 'Idle', tags: ['space', 'clicker', 'idle'], players: 'Single', status: 'playable', features: ['Planetary upgrades', 'Star clusters', 'Galactic achievements'] },
  { id: 2, name: 'CatRunner', description: 'ASCII-art Dino Runner clone with cats, dogs, and obstacles to dodge', category: 'Arcade', tags: ['ascii', 'runner', 'cats'], players: 'Single', status: 'playable', features: ['Multiple difficulty modes', 'Procedural generation', 'High score system'] },
  { id: 3, name: 'Chess', description: 'Classic chess with AI opponents and multiplayer support', category: 'Strategy', tags: ['chess', 'ai', 'multiplayer'], players: '1-2', status: 'active', features: ['Multiple AI difficulties', 'Online multiplayer', 'Move analysis'] },
  { id: 4, name: 'Solitaire', description: 'Klondike and Spider solitaire with customizable rules', category: 'Arcade', tags: ['cards', 'solitaire', 'puzzle'], players: 'Single', status: 'playable', features: ['Klondike & Spider', 'Undo/redo', 'Statistics tracking'] },
  { id: 5, name: 'Sudoku', description: 'Number puzzle game with standard, killer, and nightmare variants', category: 'Strategy', tags: ['puzzle', 'numbers', 'logic'], players: 'Single', status: 'playable', features: ['Multiple difficulty levels', 'Hint system', 'Puzzle generator'] },
  { id: 6, name: 'Wordle', description: 'Word guessing game with 5, 6, and 7 letter variations', category: 'Word', tags: ['word', 'puzzle', 'guess'], players: 'Single', status: 'playable', features: ['Multiple word lengths', 'Daily challenges', 'Statistics'] },
  { id: 7, name: 'Bones', description: 'Inscryption-like table-top card game with dark themes', category: 'Card', tags: ['cards', 'strategy', 'tabletop'], players: '1-2', status: 'development', features: ['Deck building', 'Campaign mode', 'Online multiplayer'] }
]

const categories = ['All', 'Idle', 'Arcade', 'Strategy', 'Word', 'Card']
const activeCategory = ref('All')

const filteredGames = computed(() => {
  if (activeCategory.value === 'All') return games
  return games.filter(game => game.category === activeCategory.value)
})

const totalGames = computed(() => games.length)
const multiplayerGames = computed(() => games.filter(game => game.players.includes('2')).length)

const multiplayerGamesList = computed(() => games.filter(game => game.players.includes('2')).map(game => ({
  id: game.id,
  name: game.name,
  description: game.description,
  players: game.players,
  modes: game.players === '1-2' ? ['vs AI', 'vs Player'] : ['Multiplayer']
})))

const setActiveCategory = (category: string) => {
  activeCategory.value = category
}
</script>

<style scoped>
.games-page {
  @apply max-w-7xl mx-auto;
}

.games-header {
  @apply pb-8 border-b border-accent/20;
}

.game-stat {
  @apply px-6 py-3 rounded-lg bg-bg-secondary/30 border border-accent/10;
}

.game-stat .stat-number {
  @apply text-2xl font-bold text-accent font-mono;
}

.game-stat .stat-label {
  @apply text-text-secondary text-sm mt-1;
}

.game-filter {
  @apply px-4 py-2 rounded-full border border-accent/20 text-sm font-mono
         text-text-secondary hover:text-accent hover:border-accent/40
         transition-all duration-200;
}

.game-filter.active {
  @apply bg-accent/10 text-accent border-accent/40;
}

.game-category {
  @apply p-4 rounded-lg bg-bg-secondary/30 border border-accent/5;
}

.multiplayer-game {
  @apply p-4 rounded-lg bg-bg-secondary/30 border border-accent/5;
}
</style>