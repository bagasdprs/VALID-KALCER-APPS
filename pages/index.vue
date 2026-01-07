<script setup lang="ts">
import ScannerHero from "~/components/skena/ScannerHero.vue";
import Services from "~/components/landing/Services.vue";
import Testimonials from "~/components/landing/Testimonials.vue";

definePageMeta({
  layout: "default",
});
</script>

<template>
  <div>
    <ScannerHero />
    <Services />
    <Testimonials />

    <!-- <div class="min-h-[50vh] bg-black flex flex-col items-center justify-center border-t border-white/10">
      <p class="text-zinc-500 font-mono text-sm mb-8">// DIPERCAYA OLEH ANAK SKENA DI:</p>

      <div class="flex flex-wrap justify-center gap-8 opacity-50 hover:opacity-100 transition-opacity">
        <span class="font-bold text-zinc-400 text-xl tracking-tighter">📍 BLOK M</span>
        <span class="font-bold text-zinc-400 text-xl tracking-tighter">📍 TUNJUNGAN</span>
        <span class="font-bold text-zinc-400 text-xl tracking-tighter">📍 BRAGA</span>
      </div>
    </div> -->
  </div>
</template>
