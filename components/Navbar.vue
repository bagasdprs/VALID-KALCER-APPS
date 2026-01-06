<script setup lang="ts">
const supabase = useSupabaseClient();
const user = useSupabaseUser();

// State Dropdown Profile
const isProfileOpen = ref(false);

// Logout function
const handleLogout = async () => {
  const { error } = await supabase.auth.signOut();
  if (!error) {
    navigateTo("/login");
  }
};

// Items buat Dropdown (Pake Nuxt UI Dropdown nanti kalau mau lebih canggih,
// tapi ini versi manual biar gampang di-style neon)
</script>

<template>
  <nav class="fixed top-0 left-0 right-0 z-50 px-6 py-4 flex items-center justify-between bg-black/50 backdrop-blur-md border-b border-white/5 transition-all duration-300">
    <NuxtLink to="/dashboard" class="flex items-center gap-3 group">
      <div class="w-10 h-10 rounded-lg bg-skena/10 border border-skena flex items-center justify-center group-hover:shadow-[0_0_15px_rgba(57,255,20,0.5)] transition-all">
        <UIcon name="i-heroicons-qr-code-solid" class="w-6 h-6 text-skena" />
      </div>
      <span class="text-xl font-black text-white tracking-tighter"> VIBE <span class="text-skena">CHECK</span> </span>
    </NuxtLink>

    <div class="flex items-center gap-4">
      <button class="w-10 h-10 rounded-full flex items-center justify-center text-zinc-400 hover:text-white hover:bg-white/10 transition-colors">
        <UIcon name="i-heroicons-cog-6-tooth-solid" class="w-6 h-6" />
      </button>

      <div class="relative">
        <button
          @click="isProfileOpen = !isProfileOpen"
          class="w-10 h-10 rounded-full bg-zinc-800 border border-zinc-700 flex items-center justify-center hover:border-skena hover:shadow-[0_0_10px_rgba(57,255,20,0.3)] transition-all overflow-hidden"
        >
          <img v-if="user?.user_metadata?.avatar_url" :src="user.user_metadata.avatar_url" class="w-full h-full object-cover" />
          <UIcon v-else name="i-heroicons-user-solid" class="w-5 h-5 text-white" />
        </button>

        <div v-if="isProfileOpen" class="absolute right-0 mt-2 w-48 bg-foreground border border-zinc-800 rounded-xl shadow-xl py-2 animate-in fade-in slide-in-from-top-2">
          <div class="px-4 py-2 border-b border-zinc-800 mb-1">
            <p class="text-xs text-zinc-500 font-bold uppercase">Signed in as</p>
            <p class="text-sm text-skena font-bold truncate">{{ user?.email }}</p>
          </div>

          <NuxtLink to="/profile" class="block px-4 py-2 text-sm text-zinc-300 hover:bg-zinc-900 hover:text-white"> My Profile </NuxtLink>

          <button @click="handleLogout" class="w-full text-left px-4 py-2 text-sm text-red-500 hover:bg-red-500/10 hover:text-red-400">Log Out</button>
        </div>

        <div v-if="isProfileOpen" @click="isProfileOpen = false" class="fixed inset-0 z-[-1] cursor-default"></div>
      </div>
    </div>
  </nav>
</template>
