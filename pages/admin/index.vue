<script setup lang="ts">
definePageMeta({
  layout: "admin",
});
</script>

<template>
  <div>
    <h2 class="text-3xl font-black text-white mb-6">Dashboard Overview 📊</h2>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div class="p-6 rounded-2xl bg-zinc-900 border border-white/10">
        <p class="text-zinc-400 text-sm">Total Users</p>
        <p class="text-4xl font-black text-white mt-2">1,240</p>
      </div>
      <div class="p-6 rounded-2xl bg-zinc-900 border border-white/10">
        <p class="text-zinc-400 text-sm">Total Scans</p>
        <p class="text-4xl font-black text-skena mt-2">58,900</p>
      </div>
      <div class="p-6 rounded-2xl bg-zinc-900 border border-white/10">
        <p class="text-zinc-400 text-sm">Revenue</p>
        <p class="text-4xl font-black text-white mt-2">Rp 0,-</p>
      </div>
    </div>

    <div class="mt-8 p-10 text-center border-2 border-dashed border-zinc-800 rounded-3xl">
      <p class="text-zinc-500">Grafik Chart akan kita pasang nanti disini...</p>
    </div>
  </div>
</template>
