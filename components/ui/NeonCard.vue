<template>
  <div class="relative p-0.5 overflow-hidden rounded-xl group">
    <div class="absolute inset-0 bg-linear-to-r from-transparent via-green-500 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-500 animate-spin-slow"></div>

    <div class="absolute inset-0 border-2 border-green-500/20 rounded-xl"></div>

    <div class="relative bg-black rounded-xl p-6 h-full border border-white/5">
      <slot />
    </div>
  </div>
</template>

<style scoped>
/* Bikin animasi muter pelan buat gradientnya */
@keyframes spin-slow {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
.animate-spin-slow {
  animation: spin-slow 4s linear infinite;
  /* Bikin gradientnya lebih gede dari kotaknya biar muternya keliatan */
  width: 200%;
  height: 200%;
  top: -50%;
  left: -50%;
}
</style>
