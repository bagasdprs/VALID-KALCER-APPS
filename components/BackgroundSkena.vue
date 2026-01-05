<script setup lang="ts">
// Kita pake VueUse yang baru diinstall tadi
import { useWindowSize, useRafFn } from "@vueuse/core";

// 1. Ambil ukuran layar user (realtime)
const { width: windowWidth, height: windowHeight } = useWindowSize();

// 2. Settingan Posisi Awal
const x = ref(100);
const y = ref(100);
const dx = ref(2); // Kecepatan Horizontal (Makin gede makin ngebut)
const dy = ref(2); // Kecepatan Vertikal

// Ukuran elemen yang mau dipantulin (Logo/Stiker)
const elWidth = 150;
const elHeight = 150;

// 3. Logic Fisika Pantulan (Jantungnya)
// useRafFn = Request Animation Frame (Jalan 60fps smooth)
useRafFn(() => {
  // Cek nabrak tembok kanan/kiri?
  if (x.value + elWidth >= windowWidth.value || x.value <= 0) {
    dx.value = -dx.value; // Balik arah!
  }

  // Cek nabrak tembok atas/bawah?
  if (y.value + elHeight >= windowHeight.value || y.value <= 0) {
    dy.value = -dy.value; // Balik arah!
  }

  // Update posisi
  x.value += dx.value;
  y.value += dy.value;
});

const elementStyle = computed(() => ({
  transform: `translate(${x.value}px, ${y.value}px)`,
  width: `${elWidth}px`,
  height: `${elHeight}px`,
  // Kasih transisi dikit biar ga patah-patah banget
  willChange: "transform",
}));
</script>

<template>
  <div class="fixed inset-0 -z-10 overflow-hidden bg-skena-dark">
    <div class="absolute inset-0 opacity-20 pointer-events-none" style="background-image: url(&quot;https://grainy-gradients.vercel.app/noise.svg&quot;)"></div>

    <div :style="elementStyle" class="absolute top-0 left-0">
      <div class="w-full h-full flex items-center justify-center animate-glitch">
        <div class="text-skena drop-shadow-[0_0_15px_rgba(57,255,20,0.8)]">
          <UIcon name="i-heroicons-sparkles-solid" class="w-24 h-24" />
        </div>
      </div>
    </div>

    <div class="absolute top-20 right-20 text-skena animate-float opacity-50">
      <UIcon name="i-heroicons-star-solid" class="w-12 h-12" />
    </div>

    <div class="absolute bottom-20 left-20 text-purple-500 animate-float opacity-50" style="animation-delay: 1s">
      <UIcon name="i-heroicons-plus-solid" class="w-16 h-16" />
    </div>
  </div>
</template>
