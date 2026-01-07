<template>
  <div class="fixed inset-0 z-[-1] bg-[#051a05] overflow-hidden pointer-events-none">
    <div
      class="absolute inset-0 bg-[linear-gradient(to_right,#1f1f1f_1px,transparent_1px),linear-gradient(to_bottom,#1f1f1f_1px,transparent_1px)] bg-size-[4rem_4rem] mask-[radial-gradient(ellipse_60%_50%_at_50%_0%,#000_70%,transparent_100%)] opacity-20"
    ></div>

    <div class="absolute top-0 left-1/2 -translate-x-1/2 w-250 h-100 bg-skena/20 blur-[120px] rounded-full opacity-30 mix-blend-screen animate-pulse-slow"></div>

    <div class="absolute bottom-0 right-0 w-200 h-100 bg-purple-900/20 blur-[100px] rounded-full opacity-20"></div>
  </div>
</template>

<style scoped>
.animate-pulse-slow {
  animation: pulse 8s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

@keyframes pulse {
  0%,
  100% {
    opacity: 0.2;
  }
  50% {
    opacity: 0.4;
  }
}
</style>
