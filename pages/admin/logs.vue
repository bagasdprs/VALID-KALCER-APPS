<script setup lang="ts">
definePageMeta({ layout: "admin" });
</script>

<template>
  <div>
    <h2 class="text-3xl font-black text-white mb-8">Scan Activity Logs 📜</h2>

    <div class="space-y-4">
      <div class="p-4 bg-zinc-900 border border-white/10 rounded-xl flex gap-4 items-start">
        <div class="w-16 h-16 bg-zinc-800 rounded-lg flex-shrink-0">
          <img src="https://placehold.co/100" class="w-full h-full object-cover rounded-lg opacity-50" />
        </div>
        <div class="flex-1">
          <div class="flex justify-between items-start">
            <h4 class="font-bold text-white">User #123 <span class="text-zinc-500 font-normal">scanned an outfit</span></h4>
            <span class="text-xs text-zinc-500">2 mins ago</span>
          </div>
          <p class="text-zinc-400 text-sm mt-1 italic">"Outfit lo kayak tukang parkir Blok M, tapi versi premium dikit..."</p>
          <div class="mt-2 flex gap-2">
            <span class="text-xs border border-skena/30 text-skena px-2 py-1 rounded bg-skena/10">Score: 78/100</span>
            <span class="text-xs border border-white/10 text-zinc-400 px-2 py-1 rounded">Style: Skena</span>
          </div>
        </div>
      </div>

      <div class="p-4 bg-zinc-900 border border-white/10 rounded-xl flex gap-4 items-start">
        <div class="w-16 h-16 bg-zinc-800 rounded-lg flex-shrink-0"></div>
        <div class="flex-1">
          <h4 class="font-bold text-white">User #99 <span class="text-zinc-500 font-normal">scanned an outfit</span></h4>
          <p class="text-zinc-400 text-sm mt-1 italic">"Warna nabrak parah, lu mau ke pasar malem apa fashion show?"</p>
        </div>
      </div>
    </div>
  </div>
</template>
