<!-- app/app.vue -->
<template>
  <div>
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
    
    <!-- Matrix Rain Background -->
    <div class="background-animation">
      <canvas ref="matrixCanvas" class="matrix-rain"></canvas>
    </div>
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
  
  for (let i = 0; i < columnCount; i++) {
    drops[i] = Math.random() * -100
  }
  
  const draw = () => {
    const canvas = matrixCanvas.value
    if (!ctx || !canvas) return
    
    ctx.fillStyle = 'rgba(10, 10, 10, 0.05)'
    ctx.fillRect(0, 0, canvas.width, canvas.height)
    
    ctx.fillStyle = '#8b8be9'
    ctx.font = `${fontSize}px monospace`
    
    for (let i = 0; i < drops.length; i++) {
      const char = getRandomChar()
      const x = i * fontSize
      const y = drops[i]! * fontSize
      
      ctx.fillText(char, x, y)
      
      if (y > canvas.height && Math.random() > 0.975) {
        drops[i] = 0
      }
      
      drops[i]!++
    }
    
    animationId = requestAnimationFrame(draw)
  }
  
  draw()
  
  const handleResize = () => {
    if (!matrixCanvas.value || !ctx) return
    
    const canvas = matrixCanvas.value
    canvas.width = window.innerWidth
    canvas.height = window.innerHeight
    
    const columnCount = Math.floor(canvas.width / fontSize)
    
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

<style>
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