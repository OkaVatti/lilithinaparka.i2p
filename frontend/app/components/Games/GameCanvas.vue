<!-- frontend/app/components/Games/GameCanvas.vue -->
<template>
  <div class="game-canvas-container">
    <div class="game-header">
      <div class="flex items-center justify-between mb-4">
        <h2 class="text-2xl font-bold font-mono text-accent">{{ game.name }}</h2>
        <button @click="$emit('close')" class="close-button">✕</button>
      </div>
      <p class="text-text-secondary text-sm mb-4">{{ game.description }}</p>
      <div class="game-stats">
        <div class="stat">
          <span class="stat-label">Score:</span>
          <span class="stat-value">{{ score }}</span>
        </div>
        <div class="stat">
          <span class="stat-label">Level:</span>
          <span class="stat-value">{{ level }}</span>
        </div>
        <div class="stat">
          <span class="stat-label">High Score:</span>
          <span class="stat-value">{{ highScore }}</span>
        </div>
      </div>
    </div>
    
    <div class="game-viewport">
      <canvas 
        ref="canvas" 
        :width="canvasWidth" 
        :height="canvasHeight"
        class="game-canvas"
        @click="handleCanvasClick"
        @mousemove="handleMouseMove"
        @keydown="handleKeyDown"
        @keyup="handleKeyUp"
        tabindex="0"
      ></canvas>
      
      <div v-if="gameState === 'paused'" class="game-overlay">
        <div class="overlay-content">
          <h3 class="text-3xl font-bold mb-4">PAUSED</h3>
          <button @click="resume" class="game-button">Resume</button>
          <button @click="restart" class="game-button">Restart</button>
        </div>
      </div>
      
      <div v-if="gameState === 'gameover'" class="game-overlay">
        <div class="overlay-content">
          <h3 class="text-3xl font-bold mb-4">GAME OVER</h3>
          <p class="text-xl mb-6">Final Score: {{ score }}</p>
          <input 
            v-model="playerAlias" 
            type="text" 
            placeholder="Enter your name"
            class="alias-input mb-4"
            maxlength="20"
          />
          <button @click="submitScore" class="game-button">Submit Score</button>
          <button @click="restart" class="game-button">Play Again</button>
        </div>
      </div>
    </div>
    
    <div class="game-controls">
      <button @click="togglePause" class="control-button">
        {{ gameState === 'paused' ? '▶' : '⏸' }}
      </button>
      <button @click="restart" class="control-button">⟲</button>
      <button @click="toggleSound" class="control-button">
        {{ soundEnabled ? '🔊' : '🔇' }}
      </button>
      <button @click="saveGame" class="control-button">💾</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from 'vue'
import { useApi } from '~/composables/useApi';

const props = defineProps<{
  game: {
    id: number
    name: string
    slug: string
    description: string
    category: string
  }
}>()

const emit = defineEmits(['close', 'scoreSubmitted'])

const canvas = ref<HTMLCanvasElement | null>(null)
const canvasWidth = 800
const canvasHeight = 600

const score = ref(0)
const level = ref(1)
const highScore = ref(0)
const gameState = ref<'playing' | 'paused' | 'gameover'>('playing')
const playerAlias = ref('')
const soundEnabled = ref(true)

let ctx: CanvasRenderingContext2D | null = null
let animationId: number | null = null
let gameLoop: (() => void) | null = null

onMounted(() => {
  if (!canvas.value) return
  ctx = canvas.value.getContext('2d')
  if (!ctx) return
  
  canvas.value.focus()
  loadHighScore()
  initGame()
  startGameLoop()
})

onUnmounted(() => {
  stopGameLoop()
})

const initGame = () => {
  score.value = 0
  level.value = 1
  gameState.value = 'playing'
}

const startGameLoop = () => {
  const loop = () => {
    if (gameState.value === 'playing') {
      update()
      render()
    }
    animationId = requestAnimationFrame(loop)
  }
  loop()
}

const stopGameLoop = () => {
  if (animationId) {
    cancelAnimationFrame(animationId)
  }
}

const update = () => {
  // Game logic here
}

const render = () => {
  if (!ctx) return
  
  // Clear canvas
  ctx.fillStyle = '#0a0a0a'
  ctx.fillRect(0, 0, canvasWidth, canvasHeight)
  
  // Render game objects
}

const handleCanvasClick = (event: MouseEvent) => {
  if (gameState.value !== 'playing') return
  // Handle click
}

const handleMouseMove = (event: MouseEvent) => {
  if (gameState.value !== 'playing') return
  // Handle mouse move
}

const handleKeyDown = (event: KeyboardEvent) => {
  if (event.code === 'Escape') {
    togglePause()
    return
  }
  if (gameState.value !== 'playing') return
  // Handle key down
}

const handleKeyUp = (event: KeyboardEvent) => {
  if (gameState.value !== 'playing') return
  // Handle key up
}

const togglePause = () => {
  if (gameState.value === 'playing') {
    gameState.value = 'paused'
  } else if (gameState.value === 'paused') {
    gameState.value = 'playing'
  }
}

const resume = () => {
  gameState.value = 'playing'
}

const restart = () => {
  initGame()
  startGameLoop()
}

const toggleSound = () => {
  soundEnabled.value = !soundEnabled.value
}

const saveGame = async () => {
  if (!playerAlias.value) {
    alert('Please enter your name first')
    return
  }
  
  const { game: gameApi } = useApi()
  try {
    const saveData = JSON.stringify({
      score: score.value,
      level: level.value,
      // Add more game state here
    })
    
    await gameApi.saveGame(props.game.slug, playerAlias.value, saveData)
    alert('Game saved!')
  } catch (error) {
    console.error('Failed to save game:', error)
    alert('Failed to save game')
  }
}

const submitScore = async () => {
  if (!playerAlias.value) {
    alert('Please enter your name')
    return
  }
  
  const { game: gameApi } = useApi()
  try {
    await gameApi.submitScore(props.game.slug, {
      alias: playerAlias.value,
      score: score.value,
      level: level.value,
    })
    
    if (score.value > highScore.value) {
      highScore.value = score.value
      localStorage.setItem(`${props.game.slug}_highscore`, score.value.toString())
    }
    
    emit('scoreSubmitted')
    alert('Score submitted!')
  } catch (error) {
    console.error('Failed to submit score:', error)
    alert('Failed to submit score')
  }
}

const loadHighScore = () => {
  const saved = localStorage.getItem(`${props.game.slug}_highscore`)
  if (saved) {
    highScore.value = parseInt(saved, 10)
  }
}
</script>

<style scoped>
.game-canvas-container {
  @apply bg-bg-secondary rounded-xl border border-accent/20 p-6;
}

.game-header {
  @apply mb-4;
}

.close-button {
  @apply w-8 h-8 rounded-full bg-bg-primary text-text-secondary 
         hover:text-accent hover:bg-accent/10 transition-all duration-200;
}

.game-stats {
  @apply flex gap-6 text-sm font-mono;
}

.stat {
  @apply flex items-center gap-2;
}

.stat-label {
  @apply text-text-secondary;
}

.stat-value {
  @apply text-accent font-bold;
}

.game-viewport {
  @apply relative bg-bg-primary rounded-lg overflow-hidden;
}

.game-canvas {
  @apply w-full h-auto border border-accent/10;
}

.game-overlay {
  @apply absolute inset-0 bg-bg-primary/95 backdrop-blur-sm 
         flex items-center justify-center;
}

.overlay-content {
  @apply text-center p-8 bg-bg-secondary rounded-xl border border-accent/20;
}

.game-button {
  @apply px-6 py-3 bg-accent text-white rounded-lg font-mono 
         hover:bg-accent/80 transition-all duration-200 mx-2;
}

.alias-input {
  @apply w-full px-4 py-2 bg-bg-primary border border-accent/20 
         rounded-lg text-text-primary font-mono focus:outline-none 
         focus:border-accent/40;
}

.game-controls {
  @apply flex justify-center gap-4 mt-4;
}

.control-button {
  @apply w-12 h-12 rounded-full bg-bg-primary text-text-secondary 
         hover:text-accent hover:bg-accent/10 transition-all duration-200 
         text-xl;
}
</style>