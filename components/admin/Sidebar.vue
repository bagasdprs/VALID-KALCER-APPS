<script setup lang="ts">
const route = useRoute();
const links = [
  { name: "Dashboard", icon: "i-heroicons-home", to: "/admin" },
  { name: "User Management", icon: "i-heroicons-users", to: "/admin/users" },
  { name: "Scan Logs", icon: "i-heroicons-document-text", to: "/admin/logs" },
  { name: "System Settings", icon: "i-heroicons-cog-6-tooth", to: "/admin/settings" },
];
</script>

<template>
  <aside class="w-64 h-screen bg-[#050505] border-r border-white/10 flex flex-col fixed left-0 top-0 z-50">
    <div class="p-6 border-b border-white/5 flex items-center gap-3">
      <div class="w-8 h-8 bg-skena rounded flex items-center justify-center">
        <UIcon name="i-heroicons-shield-check-solid" class="text-black w-5 h-5" />
      </div>
      <h1 class="font-black text-white tracking-tighter text-lg">ADMIN <span class="text-skena">PANEL</span></h1>
    </div>

    <nav class="flex-1 p-4 space-y-2">
      <NuxtLink
        v-for="link in links"
        :key="link.to"
        :to="link.to"
        class="flex items-center gap-3 px-4 py-3 rounded-xl text-sm font-bold transition-all"
        :class="route.path === link.to ? 'bg-skena text-black shadow-[0_0_15px_rgba(57,255,20,0.4)]' : 'text-zinc-500 hover:text-white hover:bg-white/5'"
      >
        <UIcon :name="link.icon" class="w-5 h-5" />
        {{ link.name }}
      </NuxtLink>
    </nav>

    <div class="p-4 border-t border-white/5">
      <button class="w-full flex items-center gap-3 px-4 py-3 text-red-500 hover:bg-red-500/10 rounded-xl transition-colors font-bold text-sm">
        <UIcon name="i-heroicons-arrow-left-on-rectangle" class="w-5 h-5" />
        Logout Admin
      </button>
    </div>
  </aside>
</template>
