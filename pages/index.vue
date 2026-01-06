<script setup lang="ts">
definePageMeta({
  layout: "default",
});
</script>

<template>
  <div class="flex flex-col items-center justify-center min-h-[80vh] text-center">
    <div class="space-y-6 max-w-2xl animate-fade-in-up">
      <div class="inline-block px-4 py-1 rounded-full border border-skena/30 bg-skena/10 text-skena text-xs font-bold tracking-widest uppercase mb-4">✨ AI Fashion Rater #1 di Jaksel</div>

      <h1 class="text-5xl md:text-7xl font-black text-white tracking-tighter leading-tight">
        CEK SEBERAPA <br />
        <span class="text-transparent bg-clip-text bg-gradient-to-r from-skena to-white">SKENA OUTFIT LO.</span>
      </h1>

      <p class="text-zinc-400 text-lg md:text-xl leading-relaxed">
        Jangan sampai salah kostum. Biar AI gue yang nge-judge outfit lo sebelum netizen yang julid.
        <span class="text-white font-bold">Gratis, pedes, dan akurat.</span>
      </p>

      <div class="flex flex-col md:flex-row gap-4 justify-center pt-8">
        <button class="group relative px-8 py-4 bg-skena text-black font-black rounded-full text-lg overflow-hidden transition-all hover:scale-105 shadow-[0_0_30px_rgba(57,255,20,0.4)]">
          <span class="relative z-10 flex items-center gap-2">
            <UIcon name="i-heroicons-camera-solid" class="w-6 h-6" />
            SCAN OUTFIT GUE
          </span>
          <div class="absolute inset-0 bg-white/20 translate-y-full group-hover:translate-y-0 transition-transform duration-300"></div>
        </button>

        <NuxtLink to="/settings" class="px-8 py-4 border border-white/20 text-white font-bold rounded-full hover:bg-white/5 transition-colors"> Tutorial Dulu </NuxtLink>
      </div>
    </div>

    <div class="mt-20 flex flex-col items-center gap-4 opacity-50 hover:opacity-100 transition-opacity">
      <p class="text-xs text-zinc-500 uppercase tracking-widest">Dipercaya oleh anak-anak skena di</p>
      <div class="flex gap-8 grayscale hover:grayscale-0 transition-all">
        <span class="font-bold text-zinc-400">📍 Blok M</span>
        <span class="font-bold text-zinc-400">📍 Tunjungan</span>
        <span class="font-bold text-zinc-400">📍 Braga</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.animate-fade-in-up {
  animation: fadeInUp 0.8s ease-out forwards;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
