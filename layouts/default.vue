<template>
  <div class="min-h-screen bg-black text-white flex flex-col font-sans">
    <header class="fixed top-0 w-full z-50 bg-black/80 backdrop-blur-md border-b border-white/10">
      <div class="max-w-7xl mx-auto px-6 h-16 flex items-center justify-between">
        <NuxtLink to="/" class="flex items-center gap-2">
          <div class="w-8 h-8 bg-skena rounded flex items-center justify-center">
            <UIcon name="i-heroicons-sparkles-solid" class="text-black w-5 h-5" />
          </div>
          <span class="font-black tracking-tighter text-xl">VALID <span class="text-skena">KALCER</span></span>
        </NuxtLink>

        <div class="flex items-center gap-4">
          <NuxtLink to="/login" class="text-sm font-bold hover:text-skena transition-colors">Login</NuxtLink>
          <NuxtLink to="/register" class="bg-white text-black px-5 py-2 rounded-full text-sm font-bold hover:bg-skena transition-colors"> Join Skena </NuxtLink>
          <NuxtLink to="/settings" class="flex items-center gap-3 hover:bg-white/5 px-3 py-2 rounded-full transition-colors">
            <div class="text-right hidden md:block">
              <p class="text-sm font-bold text-white leading-none">Bagas</p>
              <p class="text-[10px] text-skena uppercase tracking-wider">Member</p>
            </div>
            <div class="w-9 h-9 rounded-full bg-zinc-800 border border-white/20 overflow-hidden">
              <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Felix" />
            </div>
          </NuxtLink>
        </div>
      </div>
    </header>

    <main class="flex-1 pt-20 px-6">
      <div class="max-w-7xl mx-auto">
        <slot />
      </div>
    </main>

    <footer class="border-t border-white/10 py-8 mt-20 bg-zinc-950">
      <div class="max-w-7xl mx-auto px-6 text-center">
        <p class="text-zinc-500 text-sm">&copy; 2026 Valid Kalcer. Dibuat dengan <span class="text-red-500">☕</span> dan <span class="text-skena">Golang</span>.</p>
      </div>
    </footer>
  </div>
</template>
