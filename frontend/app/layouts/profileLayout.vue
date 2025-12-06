<!-- app/layouts/profileLayout.vue -->
<template>
  <div class="min-h-screen bg-gradient-to-br from-bg-primary via-bg-secondary/20 to-bg-primary text-text-primary">
    <!-- Navigation -->
    <header class="sticky top-0 z-50 backdrop-blur-lg bg-bg-primary/80 border-b border-border">
      <div class="container mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <NuxtLink to="/" class="font-bold text-xl font-mono hover:text-accent transition-colors">
            Profile
          </NuxtLink>
          <nav class="flex items-center gap-6">
            <NuxtLink to="/" class="text-sm hover:text-accent transition-colors">Home</NuxtLink>
            <NuxtLink to="/artwork" class="text-sm hover:text-accent transition-colors">Artwork</NuxtLink>
            <NuxtLink to="/blog" class="text-sm hover:text-accent transition-colors">Blog</NuxtLink>
            <NuxtLink to="/profile" class="text-sm hover:text-accent transition-colors active">Profile</NuxtLink>
          </nav>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="container mx-auto px-4 py-12">
      <!-- Profile header -->
      <div class="max-w-4xl mx-auto">
        <div class="flex flex-col md:flex-row items-start md:items-center gap-8 mb-12">
          <!-- Avatar -->
          <div class="relative">
            <div class="w-32 h-32 rounded-full border-4 border-accent/20 overflow-hidden">
              <img src="/avatar.jpg" alt="Profile" class="w-full h-full object-cover">
            </div>
            <div class="absolute -bottom-2 -right-2 w-10 h-10 rounded-full bg-accent flex items-center justify-center">
              <span class="text-white text-lg">✨</span>
            </div>
          </div>
          
          <!-- Profile info -->
          <div class="flex-1">
            <h1 class="text-4xl font-bold font-mono mb-2">Your Name</h1>
            <p class="text-text-secondary text-lg mb-4">Digital Artist & Creative Developer</p>
            <div class="flex flex-wrap gap-3">
              <span class="px-3 py-1 rounded-full bg-accent/10 text-accent text-sm">Digital Art</span>
              <span class="px-3 py-1 rounded-full bg-accent/10 text-accent text-sm">Web Design</span>
              <span class="px-3 py-1 rounded-full bg-accent/10 text-accent text-sm">Open Source</span>
              <span class="px-3 py-1 rounded-full bg-accent/10 text-accent text-sm">Community</span>
            </div>
          </div>
        </div>

        <!-- Profile content -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
          <!-- Left column - About -->
          <div class="lg:col-span-2 space-y-8">
            <slot />
          </div>

          <!-- Right column - Sidebar -->
          <div class="space-y-6">
            <!-- Contact info -->
            <div class="p-6 rounded-xl border border-border bg-bg-secondary/30">
              <h3 class="text-lg font-bold font-mono mb-4">Contact & Links</h3>
              <div class="space-y-3">
                <a href="mailto:you@example.com" class="flex items-center gap-3 hover:text-accent transition-colors">
                  <span>📧</span>
                  <span>Email</span>
                </a>
                <a href="https://github.com/yourusername" target="_blank" class="flex items-center gap-3 hover:text-accent transition-colors">
                  <span>🐙</span>
                  <span>GitHub</span>
                </a>
                <a href="https://bsky.app/profile/yourhandle" target="_blank" class="flex items-center gap-3 hover:text-accent transition-colors">
                  <span>🦋</span>
                  <span>Bluesky</span>
                </a>
                <a href="https://yourportfolio.site" target="_blank" class="flex items-center gap-3 hover:text-accent transition-colors">
                  <span>🎨</span>
                  <span>Portfolio</span>
                </a>
              </div>
            </div>

            <!-- Skills -->
            <div class="p-6 rounded-xl border border-border bg-bg-secondary/30">
              <h3 class="text-lg font-bold font-mono mb-4">Skills</h3>
              <div class="space-y-4">
                <div v-for="skill in skills" :key="skill.name" class="space-y-1">
                  <div class="flex justify-between text-sm">
                    <span>{{ skill.name }}</span>
                    <span>{{ skill.level }}%</span>
                  </div>
                  <div class="h-2 rounded-full bg-border overflow-hidden">
                    <div class="h-full bg-accent rounded-full" :style="{ width: skill.level + '%' }"></div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Availability -->
            <div class="p-6 rounded-xl border border-accent/20 bg-accent/5">
              <h3 class="text-lg font-bold font-mono mb-2">Availability</h3>
              <p class="text-text-secondary text-sm mb-4">Open for select projects</p>
              <button class="w-full px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent/80 transition-colors">
                Contact for Commissions
              </button>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Footer -->
    <footer class="border-t border-border py-8 mt-12">
      <div class="container mx-auto px-4">
        <div class="max-w-4xl mx-auto">
          <div class="flex flex-col md:flex-row justify-between items-center">
            <div class="mb-4 md:mb-0">
              <p class="text-sm text-text-secondary">Made with creativity and code</p>
            </div>
            <div class="flex gap-6">
              <NuxtLink to="/profile/cv" class="text-sm hover:text-accent transition-colors">CV/Resume</NuxtLink>
              <NuxtLink to="/profile/projects" class="text-sm hover:text-accent transition-colors">Projects</NuxtLink>
              <NuxtLink to="/profile/testimonials" class="text-sm hover:text-accent transition-colors">Testimonials</NuxtLink>
            </div>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
const skills = [
  { name: 'Digital Art', level: 95 },
  { name: 'Web Design', level: 88 },
  { name: 'Vue/Nuxt', level: 85 },
  { name: 'UI/UX', level: 80 },
  { name: 'Illustration', level: 92 },
]
</script>