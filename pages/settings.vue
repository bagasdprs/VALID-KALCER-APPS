<script setup lang="ts">
// Pakai Layout Default (yang ada Header & Footer)
definePageMeta({
  layout: "default",
});

// Data Dummy buat tampilan dulu (Nanti ini diambil dari API Golang)
const form = ref({
  name: "User Skena 01",
  email: "user@skena.com",
  bio: "Hobi nongkrong di Coffee Shop, outfit minimalis tapi dompet menangis.",
  instagram: "@bagas_skena",
  aesthetic: "Skena Dark",
});

const isSaving = ref(false);

const handleSave = () => {
  isSaving.value = true;
  // Simulasi loading save
  setTimeout(() => {
    isSaving.value = false;
    alert("Profil berhasil di-update! ✨");
  }, 1000);
};
</script>

<template>
  <div class="max-w-2xl mx-auto py-10">
    <div class="mb-8 border-b border-white/10 pb-6">
      <h1 class="text-3xl font-black text-white">Edit Profile ✏️</h1>
      <p class="text-zinc-400 mt-2">Personal branding itu penting, biar dikira valid.</p>
    </div>

    <form @submit.prevent="handleSave" class="space-y-8">
      <div class="flex items-center gap-6">
        <div class="relative group">
          <div class="w-24 h-24 rounded-full bg-zinc-800 border-2 border-white/10 overflow-hidden">
            <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Felix" alt="Avatar" class="w-full h-full object-cover" />
          </div>
          <div class="absolute inset-0 bg-black/50 rounded-full flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity cursor-pointer">
            <UIcon name="i-heroicons-camera" class="w-6 h-6 text-white" />
          </div>
        </div>
        <div>
          <button type="button" class="bg-white/10 text-white text-sm font-bold px-4 py-2 rounded-lg hover:bg-white/20 transition-colors">Change Photo</button>
          <p class="text-xs text-zinc-500 mt-2">JPG/PNG, Max 2MB.</p>
        </div>
      </div>

      <div class="space-y-5">
        <div>
          <label class="block text-sm font-bold text-zinc-400 mb-2">Display Name</label>
          <input v-model="form.name" type="text" class="w-full bg-zinc-900 border border-white/10 rounded-xl px-4 py-3 text-white focus:outline-none focus:border-skena focus:ring-1 focus:ring-skena transition-all" />
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
          <div>
            <label class="block text-sm font-bold text-zinc-400 mb-2">Instagram Handle</label>
            <div class="relative">
              <span class="absolute left-4 top-3.5 text-zinc-500">@</span>
              <input v-model="form.instagram" type="text" class="w-full bg-zinc-900 border border-white/10 rounded-xl pl-8 pr-4 py-3 text-white focus:outline-none focus:border-skena" />
            </div>
          </div>
          <div>
            <label class="block text-sm font-bold text-zinc-400 mb-2">Vibe / Style</label>
            <select v-model="form.aesthetic" class="w-full bg-zinc-900 border border-white/10 rounded-xl px-4 py-3 text-white focus:outline-none focus:border-skena appearance-none">
              <option>Skena Dark</option>
              <option>Y2K Chrome</option>
              <option>Old Money</option>
              <option>Gorpcore</option>
            </select>
          </div>
        </div>

        <div>
          <label class="block text-sm font-bold text-zinc-400 mb-2">Bio (Buat pamer)</label>
          <textarea v-model="form.bio" rows="3" class="w-full bg-zinc-900 border border-white/10 rounded-xl px-4 py-3 text-white focus:outline-none focus:border-skena"></textarea>
          <p class="text-xs text-right text-zinc-600 mt-1">Max 150 characters</p>
        </div>
      </div>

      <div class="flex justify-end gap-4 pt-6 border-t border-white/10">
        <NuxtLink to="/" class="px-6 py-3 rounded-xl text-sm font-bold text-zinc-400 hover:text-white transition-colors"> Cancel </NuxtLink>
        <button type="submit" class="bg-skena text-black px-8 py-3 rounded-xl text-sm font-black hover:bg-green-400 transition-all shadow-[0_0_15px_rgba(57,255,20,0.3)]" :disabled="isSaving">
          {{ isSaving ? "SAVING..." : "SAVE CHANGES" }}
        </button>
      </div>
    </form>
  </div>
</template>
