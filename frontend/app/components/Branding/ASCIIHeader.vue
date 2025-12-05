<template>
  <header class="border-b border-border py-4 mb-8">
    <div class="container mx-auto px-4">
      <nav class="flex items-center justify-between">
        <NuxtLink to="/" class="text-xl font-bold hover:text-accent">
          <span class="text-accent">[</span>
            <p>
                ::        :::::::   <br>
                ::        ::   :::  <br>
                ::        :::::::   <br>
                ::        :::       <br>
                ::::::::  :::       <br>
            </p>
          <span class="text-accent">]</span>
        </NuxtLink>
        
        <div class="flex gap-6">
          <NuxtLink to="/" class="hover:text-accent">
            /home/
          </NuxtLink>
          <NuxtLink to="/blog" class="hover:text-accent">
            /blog/
          </NuxtLink>
          <NuxtLink to="/social" class="hover:text-accent">
            /social/
          </NuxtLink>
          <NuxtLink to="/profile" class="hover:text-accent">
            /profile/
          </NuxtLink>
        </div>
      </nav>
    </div>
  </header>
</template>