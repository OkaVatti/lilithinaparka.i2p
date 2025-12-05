<template>
  <header class="border-b border-border py-4 mb-8 bg-bg-primary">
    <div class="container mx-auto px-4">
      <nav class="flex flex-col md:flex-row items-center justify-between gap-4">
        <NuxtLink to="/" class="font-mono text-accent hover:text-text-primary transition-colors">
          <pre class="text-xs md:text-sm leading-tight">
            ::     :::::::: ::     :::::::: :::::::: ::    :: 
            ::        ::    ::        ::       ::    ::    ::
            ::        ::    ::        ::       ::    ::::::::
            :::::: :::::::: :::::: ::::::::    ::    ::    ::
          </pre>
        </NuxtLink>
        
        <div class="flex gap-4 md:gap-6 font-mono text-sm">
          <NuxtLink 
            to="/" 
            class="hover:text-accent transition-colors"
            :class="{ 'text-accent': $route.path === '/' }"
          >
            <span class="text-accent">[</span>home<span class="text-accent">]</span>
          </NuxtLink>
          <NuxtLink 
            to="/blog" 
            class="hover:text-accent transition-colors"
            :class="{ 'text-accent': $route.path.startsWith('/blog') }"
          >
            <span class="text-accent">[</span>blog<span class="text-accent">]</span>
          </NuxtLink>
          <NuxtLink 
            to="/social" 
            class="hover:text-accent transition-colors"
            :class="{ 'text-accent': $route.path === '/social' }"
          >
            <span class="text-accent">[</span>social<span class="text-accent">]</span>
          </NuxtLink>
          <NuxtLink 
            to="/profile" 
            class="hover:text-accent transition-colors"
            :class="{ 'text-accent': $route.path === '/profile' }"
          >
            <span class="text-accent">[</span>profile<span class="text-accent">]</span>
          </NuxtLink>
        </div>
      </nav>
    </div>
  </header>
</template>

<style scoped>
pre {
  font-family: 'Courier New', monospace;
  margin: 0;
}
</style>