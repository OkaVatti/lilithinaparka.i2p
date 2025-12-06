<!-- app/components/Branding/ASCIIHeader.vue -->
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