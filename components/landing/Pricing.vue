<script setup lang="ts">
const plans = [
  {
    name: "NPC_TIER",
    command: "> run basic_scan",
    price: "0",
    period: "/forever",
    desc: "Standard latency. Iklan numpuk. Roastingan kasar tanpa solusi.",
    features: [
      { text: "3 Scans / Day", active: true },
      { text: "Brutal Roasting Only", active: true },
      { text: "Ads Everywhere", active: true },
      { text: "No Style Advice", active: false },
      { text: "No Leaderboard", active: false },
    ],
    highlight: false,
    btnCmd: "npm install basic",
  },
  {
    name: "MAIN_CHARACTER",
    command: "> run aesthetic_patch",
    price: "29k",
    period: "/month",
    desc: "Upgrade aura lo. Dapet saran perbaikan (Patch Notes) & akses Leaderboard.",
    features: [
      { text: "UNLIMITED Scans", active: true },
      { text: "Patch Notes (Saran)", active: true },
      { text: "No Ads (Clean)", active: true },
      { text: "Leaderboard Access", active: true },
      { text: "Priority Queue", active: true },
    ],
    highlight: true, // Ini yang bikin kartu tengah nyala
    btnCmd: "npm install pro --save",
  },
  {
    name: "FINAL_BOSS",
    command: "> run god_mode",
    price: "99k",
    period: "/month",
    desc: "Full admin access. Deep analytics, brand detection, & verified badge.",
    features: [
      { text: "Everything in Main Char", active: true },
      { text: "Brand Detection AI", active: true },
      { text: "Color Palette Analysis", active: true },
      { text: "'Verified Skena' Badge", active: true },
      { text: "24/7 Personal Stylist AI", active: true },
    ],
    highlight: false,
    btnCmd: "sudo install god_mode",
  },
];
</script>

<template>
  <section class="relative py-24 px-4 md:px-6 bg-zinc-950 font-mono overflow-hidden">
    <div
      class="absolute inset-0 bg-[linear-gradient(to_right,#18181b_1px,transparent_1px),linear-gradient(to_bottom,#18181b_1px,transparent_1px)] bg-[size:4rem_4rem] [mask-image:radial-gradient(ellipse_60%_50%_at_50%_0%,#000_70%,transparent_100%)] pointer-events-none"
    ></div>

    <div class="max-w-7xl mx-auto relative z-10">
      <div class="mb-16">
        <div class="flex items-center gap-2 mb-4">
          <span class="text-skena text-2xl md:text-4xl">>></span>
          <h2 class="text-3xl md:text-5xl font-black text-white uppercase tracking-tighter">INITIALIZING <span class="text-skena">PAYMENT PROTOCOLS</span>...</h2>
        </div>
        <p class="text-zinc-500 text-sm md:text-base max-w-2xl">
          // select_your_configuration. All plans include standard entropy reduction. <br />
          // encryption: enabled. refunds: disabled.
        </p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 items-start">
        <div v-for="(plan, index) in plans" :key="index" class="relative group transition-all duration-300 hover:-translate-y-2" :class="plan.highlight ? 'z-20' : 'z-10'">
          <div v-if="plan.highlight" class="absolute -inset-0.5 bg-gradient-to-b from-skena to-green-900 rounded-lg blur opacity-30 group-hover:opacity-75 transition duration-1000"></div>

          <div class="relative h-full bg-black border p-8 flex flex-col" :class="plan.highlight ? 'border-skena shadow-[0_0_30px_rgba(22,163,74,0.1)]' : 'border-zinc-800 hover:border-zinc-600'">
            <div v-if="plan.highlight" class="absolute -top-3 left-1/2 -translate-x-1/2 bg-skena text-black text-[10px] font-bold px-3 py-1 uppercase tracking-widest">Recommended</div>

            <div class="mb-6 border-b border-zinc-800 pb-6">
              <div class="text-skena text-sm mb-2 opacity-80">{{ plan.command }}</div>
              <div class="flex items-end gap-1">
                <span class="text-white text-4xl md:text-5xl font-bold tracking-tighter">{{ plan.price }}</span>
                <span class="text-zinc-500 text-sm mb-1">{{ plan.period }}</span>
              </div>
              <p class="text-zinc-400 text-xs mt-4 leading-relaxed min-h-[40px]">// {{ plan.desc }}</p>
            </div>

            <ul class="space-y-4 mb-8 flex-1">
              <li v-for="(feat, fIndex) in plan.features" :key="fIndex" class="flex items-start gap-3 text-sm" :class="feat.active ? 'text-zinc-300' : 'text-zinc-700 line-through'">
                <span v-if="feat.active" class="text-skena font-bold">[+]</span>
                <span v-else class="text-zinc-800 font-bold">[x]</span>
                {{ feat.text }}
              </li>
            </ul>

            <button
              class="w-full border bg-transparent py-4 px-4 text-left group-hover:bg-zinc-900 transition-colors relative overflow-hidden"
              :class="plan.highlight ? 'border-skena text-skena' : 'border-zinc-700 text-zinc-400 hover:text-white hover:border-white'"
            >
              <div class="flex justify-between items-center relative z-10">
                <span class="text-xs md:text-sm font-bold tracking-widest"> > {{ plan.btnCmd }}<span class="animate-pulse">_</span> </span>
                <span class="text-[10px] border border-current px-1 opacity-50">[ENTER]</span>
              </div>
            </button>
          </div>
        </div>
      </div>

      <div class="mt-12 text-center">
        <p class="text-zinc-600 text-[10px]">SERVER_USAGE: 64TB OK // CPU: 4% // UPTIME: 99.999%</p>
      </div>
    </div>
  </section>
</template>
