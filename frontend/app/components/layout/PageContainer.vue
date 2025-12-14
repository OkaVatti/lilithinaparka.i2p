<template>
  <div class="page-root">
    <slot />
  </div>
</template>

<script setup lang="ts">
// Simple page container: max width, padding, subtle surface using theme css vars.
</script>

<style scoped>
@import '~/assets/css/system.css';

.page-root {
  max-width: 1100px;
  margin: 0 auto;
  padding: 1.4rem;
  background: linear-gradient(180deg, transparent, rgba(0,0,0,0.02));
  border-radius: 10px;
  border: 1px solid var(--theme-border);
  box-shadow: 0 8px 30px rgba(0,0,0,0.35);
}
</style>
