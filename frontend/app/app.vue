<template>
  <div>
    <Header />
    <main class="main-content">
      <div class="container">
        <NuxtPage />
      </div>
    </main>
  <Footer />
</div>
</template>
<script setup lang="ts">
import { useProfileStore } from '~~/stores/profile'
const profileStore = useProfileStore()

onMounted(() => {
  profileStore.fetchProfile()
})
</script>