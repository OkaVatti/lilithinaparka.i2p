<!-- app/components/Kit/BackgroundAnimation.vue -->
<template>
  <div class="background-animation">
    <canvas ref="matrixCanvas" class="matrix-rain"></canvas>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'

const matrixCanvas = ref<HTMLCanvasElement | null>(null)
let animationId: number | null = null
let ctx: CanvasRenderingContext2D | null = null

const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789@#$%^&*()_+-=[]{}|;:,.<>?/~`'
const fontSize = 14
const drops: number[] = []

const getRandomChar = (): string => {
  const index = Math.floor(Math.random() * characters.length)
  return characters[index]!
}

onMounted(() => {
  if (!matrixCanvas.value) return
  
  const canvas = matrixCanvas.value
  const context = canvas.getContext('2d')
  
  if (!context) return
  ctx = context
  
  canvas.width = window.innerWidth
  canvas.height = window.innerHeight
  
  const columnCount = Math.floor(canvas.width / fontSize)
  
  // Initialize drops array
  for (let i = 0; i < columnCount; i++) {
    drops[i] = Math.random() * -100
  }
  
  const draw = () => {
    const canvas = matrixCanvas.value
    if (!ctx || !canvas) return
    
    // Clear with semi-transparent black for fade effect
    ctx.fillStyle = 'rgba(10, 10, 10, 0.05)'
    ctx.fillRect(0, 0, canvas.width, canvas.height)
    
    ctx.fillStyle = '#8b8be9'
    ctx.font = `${fontSize}px monospace`
    
    for (let i = 0; i < drops.length; i++) {
      const char = getRandomChar() // Use helper function to ensure non-undefined
      const x = i * fontSize
      const y = drops[i]! * fontSize // Use non-null assertion since we initialized all indices
      
      ctx.fillText(char, x, y)
      
      // Reset drop when it goes past bottom with some randomness
      if (y > canvas.height && Math.random() > 0.975) {
        drops[i] = 0
      }
      
      drops[i]!++ // Use non-null assertion
    }
    
    animationId = requestAnimationFrame(draw)
  }
  
  draw()
  
  const handleResize = () => {
    if (!matrixCanvas.value || !ctx) return
    
    const canvas = matrixCanvas.value
    canvas.width = window.innerWidth
    canvas.height = window.innerHeight
    
    // Recalculate drops array for new width
    const columnCount = Math.floor(canvas.width / fontSize)
    
    // Reset drops array for new column count
    drops.length = 0
    for (let i = 0; i < columnCount; i++) {
      drops[i] = Math.random() * -100
    }
  }
  
  window.addEventListener('resize', handleResize)
  
  onUnmounted(() => {
    if (animationId !== null) {
      cancelAnimationFrame(animationId)
    }
    window.removeEventListener('resize', handleResize)
  })
})
</script>

<style scoped>
.background-animation {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
  z-index: -1;
}

.matrix-rain {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
  opacity: 0.05;
}
</style>