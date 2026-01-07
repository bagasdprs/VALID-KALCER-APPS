<script setup lang="ts">
defineProps({
  glow: {
    type: Boolean,
    default: true,
  },
  hoverEffect: {
    type: Boolean,
    default: false,
  },
});
</script>

<template>
  <div
    class="relative bg-[#0a0a0a] border border-white/10 rounded-2xl p-6 overflow-hidden transition-all duration-300 group"
    :class="{
      'hover:border-skena/50 hover:shadow-[0_0_30px_rgba(57,255,20,0.1)] hover:-translate-y-1': hoverEffect,
    }"
  >
    <div v-if="glow" class="absolute top-0 left-0 w-full h-px bg-linear-to-r from-transparent via-skena to-transparent opacity-50"></div>

    <div class="relative z-10">
      <slot />
    </div>

    <div class="absolute inset-0 opacity-[0.03] bg-[url('https://grainy-gradients.vercel.app/noise.svg')] pointer-events-none"></div>
  </div>
</template>
