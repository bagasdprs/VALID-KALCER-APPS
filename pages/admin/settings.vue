<script setup lang="ts">
definePageMeta({ layout: "admin" });
</script>

<template>
  <div class="max-w-3xl">
    <h2 class="text-3xl font-black text-white mb-2">System Settings ⚙️</h2>
    <p class="text-zinc-400 mb-8">Atur otak AI dan limitasi aplikasi di sini.</p>

    <div class="space-y-6">
      <div class="p-6 bg-zinc-900 border border-white/10 rounded-2xl">
        <h3 class="font-bold text-xl text-white mb-4">🤖 AI Persona (Prompt Utama)</h3>
        <div class="space-y-2">
          <label class="text-sm text-zinc-400">System Prompt</label>
          <textarea rows="6" class="w-full bg-black border border-white/20 rounded-xl p-4 text-white focus:outline-none focus:border-skena transition-colors" placeholder="Ketik instruksi untuk Gemini di sini...">
Kamu adalah juri fashion yang galak, sarkas, tapi jujur. Gunakan bahasa gaul Jaksel...</textarea
          >
          <p class="text-xs text-zinc-500">Prompt ini akan dikirim ke Gemini setiap kali ada request scan.</p>
        </div>
      </div>

      <div class="p-6 bg-zinc-900 border border-white/10 rounded-2xl">
        <h3 class="font-bold text-xl text-white mb-4">⚡ Limit & Kuota</h3>
        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="text-sm text-zinc-400 block mb-2">Guest Limit (Daily)</label>
            <input type="number" value="1" class="w-full bg-black border border-white/20 rounded-xl p-3 text-white focus:border-skena outline-none" />
          </div>
          <div>
            <label class="text-sm text-zinc-400 block mb-2">Free User Limit (Daily)</label>
            <input type="number" value="3" class="w-full bg-black border border-white/20 rounded-xl p-3 text-white focus:border-skena outline-none" />
          </div>
        </div>
      </div>

      <div class="flex justify-end">
        <button class="bg-skena text-black font-black px-8 py-3 rounded-xl hover:bg-green-400 transition-colors shadow-[0_0_20px_rgba(57,255,20,0.3)]">SAVE CHANGES</button>
      </div>
    </div>
  </div>
</template>
