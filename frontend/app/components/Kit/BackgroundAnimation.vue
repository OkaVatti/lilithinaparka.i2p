<template>
  <div class="background-animation">
    <div class="matrix-rain" ref="matrixCanvas"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'

const matrixCanvas = ref<HTMLDivElement | null>(null)
let animationId: number | null = null

const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789@#$%^&*()_+-=[]{}|;:,.<>?/~`'
const fontSize = 14
const columns: number[] = []
const drops: number[] = []

onMounted(() => {
  if (!matrixCanvas.value) return
  
  const canvas = document.createElement('canvas')
  const ctx = canvas.getContext('2d')
  
  if (!ctx) return
  
  canvas.width = window.innerWidth
  canvas.height = window.innerHeight
  canvas.style.position = 'fixed'
  canvas.style.top = '0'
  canvas.style.left = '0'
  canvas.style.zIndex = '-1'
  canvas.style.opacity = '0.05'
  
  matrixCanvas.value.appendChild(canvas)
  
  const columnCount = Math.floor(canvas.width / fontSize)
  
  for (let i = 0; i < columnCount; i++) {
    drops[i] = Math.random() * -100
  }
  
  const draw = () => {
    ctx.fillStyle = 'rgba(10, 10, 10, 0.05)'
    ctx.fillRect(0, 0, canvas.width, canvas.height)
    
    ctx.fillStyle = '#8b8be9'
    ctx.font = `${fontSize}px monospace`
    
    for (let i = 0; i < drops.length; i++) {
      const char = characters[Math.floor(Math.random() * characters.length)]
      const x = i * fontSize
      const y = drops[i] * fontSize
      
      ctx.fillText(char, x, y)
      
      if (y > canvas.height && Math.random() > 0.975) {
        drops[i] = 0
      }
      
      drops[i]++
    }
    
    animationId = requestAnimationFrame(draw)
  }
  
  draw()
  
  const handleResize = () => {
    canvas.width = window.innerWidth
    canvas.height = window.innerHeight
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
  width: 100%;
  height: 100%;
}
</style>