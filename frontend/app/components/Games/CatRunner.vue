<!-- frontend/app/components/Games/CatRunner.vue -->
<template>
  <div class="cat-runner">
    <div class="game-header">
      <h2 class="text-2xl font-bold font-mono text-accent mb-2">CatRunner</h2>
      <p class="text-text-secondary text-sm mb-4">Press SPACE to jump! Avoid obstacles!</p>
    </div>
    
    <div class="game-stats">
      <div class="stat">Score: <span class="text-accent">{{ score }}</span></div>
      <div class="stat">High: <span class="text-accent">{{ highScore }}</span></div>
      <div class="stat">Speed: <span class="text-accent">{{ speed.toFixed(1) }}</span></div>
    </div>
    
    <div class="game-canvas-container" ref="gameContainer">
      <canvas 
        ref="canvas" 
        :width="800" 
        :height="400"
        class="game-canvas"
        tabindex="0"
        @keydown="handleKeyDown"
        @click="handleClick"
      ></canvas>
      
      <div v-if="gameState === 'waiting'" class="game-overlay">
        <div class="overlay-content">
          <pre class="ascii-cat">{{ waitingCat }}</pre>
          <p class="text-xl mb-4">Press SPACE or CLICK to start!</p>
        </div>
      </div>
      
      <div v-if="gameState === 'gameover'" class="game-overlay">
        <div class="overlay-content">
          <pre class="ascii-cat-dead">{{ deadCat }}</pre>
          <h3 class="text-3xl font-bold mb-2">GAME OVER</h3>
          <p class="text-xl mb-4">Score: {{ score }}</p>
          <input 
            v-if="score > 0"
            v-model="playerName" 
            type="text" 
            placeholder="Enter name"
            class="name-input mb-4"
            maxlength="20"
            @keydown.enter="submitScore"
          />
          <div class="flex gap-4 justify-center">
            <button @click="submitScore" v-if="score > 0" class="game-button">Submit Score</button>
            <button @click="restart" class="game-button">Play Again</button>
          </div>
        </div>
      </div>
    </div>
    
    <div class="game-controls">
      <p class="text-text-secondary text-sm font-mono">
        Controls: <span class="text-accent">SPACE</span> to jump | 
        <span class="text-accent">CLICK</span> to start/jump
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'

const canvas = ref<HTMLCanvasElement | null>(null)
const gameContainer = ref<HTMLDivElement | null>(null)
const score = ref(0)
const highScore = ref(0)
const speed = ref(5)
const gameState = ref<'waiting' | 'playing' | 'gameover'>('waiting')
const playerName = ref('')

let ctx: CanvasRenderingContext2D | null = null
let animationId: number | null = null

// ASCII art
const waitingCat = `
    /\\_/\\
   ( o.o )
    > ^ 
   /|   |\\
  (_|   |_)
`

const deadCat = `
    /\\_/\\
   ( x.x )
    > ~ 
   /|   |\\
  (_|   |_)
`

// Game objects
interface GameObject {
  x: number
  y: number
  width: number
  height: number
  velocity?: number
}

const cat = ref<GameObject>({
  x: 100,
  y: 280,
  width: 40,
  height: 40,
  velocity: 0
})

const obstacles = ref<GameObject[]>([])
const clouds = ref<GameObject[]>([])
const ground = 320

const gravity = 0.6
const jumpPower = -12
const obstacleSpeed = ref(5)

onMounted(() => {
  if (!canvas.value) return
  ctx = canvas.value.getContext('2d')
  if (!ctx) return
  
  loadHighScore()
  initClouds()
  canvas.value.focus()
})

onUnmounted(() => {
  stopGame()
})

const initClouds = () => {
  for (let i = 0; i < 5; i++) {
    clouds.value.push({
      x: Math.random() * 800,
      y: Math.random() * 150 + 50,
      width: 60,
      height: 30
    })
  }
}

const startGame = () => {
  gameState.value = 'playing'
  score.value = 0
  speed.value = 5
  obstacleSpeed.value = 5
  cat.value.y = 280
  cat.value.velocity = 0
  obstacles.value = []
  
  gameLoop()
}

const gameLoop = () => {
  if (gameState.value !== 'playing') return
  
  update()
  render()
  
  animationId = requestAnimationFrame(gameLoop)
}

const stopGame = () => {
  if (animationId) {
    cancelAnimationFrame(animationId)
  }
}

const update = () => {
  // Update cat
  cat.value.velocity? += gravity
  cat.value.y += cat.value.velocity?
  
  // Keep cat on ground
  if (cat.value.y >= ground - cat.value.height) {
    cat.value.y = ground - cat.value.height
    cat.value.velocity = 0
  }
  
  // Update score
  score.value += 0.1
  
  // Increase speed over time
  speed.value = 5 + Math.floor(score.value / 100)
  obstacleSpeed.value = speed.value
  
  // Spawn obstacles
  if (obstacles.value.length === 0 || obstacles.value[obstacles.value.length - 1].x < 500) {
    spawnObstacle()
  }
  
  // Update obstacles
  obstacles.value = obstacles.value.filter(obstacle => {
    obstacle.x -= obstacleSpeed.value
    return obstacle.x > -obstacle.width
  })
  
  // Update clouds
  clouds.value.forEach(cloud => {
    cloud.x -= 1
    if (cloud.x < -cloud.width) {
      cloud.x = 800
      cloud.y = Math.random() * 150 + 50
    }
  })
  
  // Check collisions
  checkCollisions()
}

const spawnObstacle = () => {
  const types = ['cactus', 'rock', 'dog']
  const type = types[Math.floor(Math.random() * types.length)]
  
  let obstacle: GameObject
  
  switch (type) {
    case 'cactus':
      obstacle = { x: 800, y: ground - 50, width: 30, height: 50 }
      break
    case 'rock':
      obstacle = { x: 800, y: ground - 30, width: 40, height: 30 }
      break
    case 'dog':
      obstacle = { x: 800, y: ground - 40, width: 50, height: 40 }
      break
    default:
      obstacle = { x: 800, y: ground - 50, width: 30, height: 50 }
  }
  
  obstacles.value.push(obstacle)
}

const checkCollisions = () => {
  for (const obstacle of obstacles.value) {
    if (
      cat.value.x < obstacle.x + obstacle.width &&
      cat.value.x + cat.value.width > obstacle.x &&
      cat.value.y < obstacle.y + obstacle.height &&
      cat.value.y + cat.value.height > obstacle.y
    ) {
      endGame()
      return
    }
  }
}

const render = () => {
  if (!ctx) return
  
  // Clear canvas
  ctx.fillStyle = '#1a1a1a'
  ctx.fillRect(0, 0, 800, 400)
  
  // Draw sky gradient
  const gradient = ctx.createLinearGradient(0, 0, 0, 300)
  gradient.addColorStop(0, '#2a2a4a')
  gradient.addColorStop(1, '#1a1a3a')
  ctx.fillStyle = gradient
  ctx.fillRect(0, 0, 800, 300)
  
  // Draw clouds
  ctx.fillStyle = '#3a3a5a'
  clouds.value.forEach(cloud => {
    ctx!.beginPath()
    ctx!.arc(cloud.x, cloud.y, 20, 0, Math.PI * 2)
    ctx!.arc(cloud.x + 20, cloud.y, 25, 0, Math.PI * 2)
    ctx!.arc(cloud.x + 40, cloud.y, 20, 0, Math.PI * 2)
    ctx!.fill()
  })
  
  // Draw ground
  ctx.fillStyle = '#2a2a2a'
  ctx.fillRect(0, ground, 800, 80)
  
  // Draw ground line
  ctx.strokeStyle = '#8b8be9'
  ctx.lineWidth = 2
  ctx.beginPath()
  ctx.moveTo(0, ground)
  ctx.lineTo(800, ground)
  ctx.stroke()
  
  // Draw cat (ASCII art style)
  drawCat()
  
  // Draw obstacles
  ctx.fillStyle = '#ff6b6b'
  obstacles.value.forEach(obstacle => {
    drawObstacle(obstacle)
  })
}

const drawCat = () => {
  if (!ctx) return
  
  const x = cat.value.x
  const y = cat.value.y
  
  ctx.fillStyle = '#8b8be9'
  ctx.font = '40px monospace'
  ctx.fillText('🐱', x, y + 35)
}

const drawObstacle = (obstacle: GameObject) => {
  if (!ctx) return
  
  ctx.fillStyle = '#ff6b6b'
  
  // Determine obstacle type based on dimensions
  if (obstacle.height === 50) {
    // Cactus
    ctx.fillRect(obstacle.x, obstacle.y, obstacle.width, obstacle.height)
    ctx.fillStyle = '#ff8b8b'
    ctx.fillRect(obstacle.x + 5, obstacle.y + 10, 5, 15)
    ctx.fillRect(obstacle.x + 20, obstacle.y + 15, 5, 15)
  } else if (obstacle.height === 30) {
    // Rock
    ctx.beginPath()
    ctx.moveTo(obstacle.x + obstacle.width / 2, obstacle.y)
    ctx.lineTo(obstacle.x + obstacle.width, obstacle.y + obstacle.height)
    ctx.lineTo(obstacle.x, obstacle.y + obstacle.height)
    ctx.closePath()
    ctx.fill()
  } else {
    // Dog
    ctx.fillRect(obstacle.x, obstacle.y, obstacle.width, obstacle.height)
    ctx.fillStyle = '#ff8b8b'
    ctx.fillRect(obstacle.x + 5, obstacle.y - 5, 10, 10)
    ctx.fillRect(obstacle.x + 35, obstacle.y - 5, 10, 10)
  }
}

const jump = () => {
  if (gameState.value === 'playing' && cat.value.y >= ground - cat.value.height) {
    cat.value.velocity = jumpPower
  }
}

const handleKeyDown = (event: KeyboardEvent) => {
  if (event.code === 'Space') {
    event.preventDefault()
    if (gameState.value === 'waiting') {
      startGame()
    } else if (gameState.value === 'playing') {
      jump()
    } else if (gameState.value === 'gameover') {
      restart()
    }
  }
}

const handleClick = () => {
  if (gameState.value === 'waiting') {
    startGame()
  } else if (gameState.value === 'playing') {
    jump()
  }
}

const endGame = () => {
  gameState.value = 'gameover'
  stopGame()
  
  if (score.value > highScore.value) {
    highScore.value = Math.floor(score.value)
    localStorage.setItem('catrunner_highscore', highScore.value.toString())
  }
}

const restart = () => {
  startGame()
}

const submitScore = async () => {
  if (!playerName.value) {
    alert('Please enter your name')
    return
  }
  
  const { game } = useApi()
  try {
    await game.submitScore('catrunner', {
      alias: playerName.value,
      score: Math.floor(score.value),
      level: Math.floor(speed.value)
    })
    alert('Score submitted!')
    restart()
  } catch (error) {
    console.error('Failed to submit score:', error)
    alert('Failed to submit score')
  }
}

const loadHighScore = () => {
  const saved = localStorage.getItem('catrunner_highscore')
  if (saved) {
    highScore.value = parseInt(saved, 10)
  }
}
</script>

<style scoped>
.cat-runner {
  @apply p-6 bg-bg-secondary rounded-xl border border-accent/20;
}

.game-header {
  @apply mb-4 text-center;
}

.game-stats {
  @apply flex justify-center gap-6 mb-4 font-mono text-sm;
}

.stat {
  @apply text-text-secondary;
}

.game-canvas-container {
  @apply relative bg-bg-primary rounded-lg overflow-hidden mb-4;
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

.ascii-cat,
.ascii-cat-dead {
  @apply text-accent font-mono text-sm mb-4 whitespace-pre;
}

.name-input {
  @apply w-full max-w-xs px-4 py-2 bg-bg-primary border border-accent/20 
         rounded-lg text-text-primary font-mono focus:outline-none 
         focus:border-accent/40;
}

.game-button {
  @apply px-6 py-2 bg-accent text-white rounded-lg font-mono 
         hover:bg-accent/80 transition-all duration-200;
}

.game-controls {
  @apply text-center;
}
</style>