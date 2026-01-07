<script setup lang="ts">
const survivors = [
  {
    id: "#USER_8821",
    name: "@scenepop_princess",
    location: "LOC: JAKSEL",
    avatar: "https://images.unsplash.com/photo-1494790108377-be9c29b29330?q=80&w=150&auto=format&fit=crop", // Cewek kacamata
    score: 8.5,
    text: "Jujurly this AI read me to filth but valid bgt fits nya. The vibe check was BRUTAL but accurate. Akhirnya ada yang paham core aesthetics gue.",
    accuracy: 98,
    dripLevel: 85,
  },
  {
    id: "#USER_4200",
    name: "skaterboi_99",
    location: "LOC: BSD CITY",
    avatar: "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?q=80&w=150&auto=format&fit=crop", // Cowok cool
    score: 6.0,
    text: "Literally shaking, AI nya savage parah. It knew my thrift finds were 'mid'. Need emotional support rn but respect the honesty.",
    accuracy: 92,
    dripLevel: 45,
  },
  {
    id: "#USER_0666",
    name: "cyber_goth_gf",
    location: "LOC: BANDUNG",
    avatar: "https://images.unsplash.com/photo-1534528741775-53994a69daeb?q=80&w=150&auto=format&fit=crop", // Cewek dark
    score: 10.0,
    text: "Akhirnya ada yang ngerti aesthetic gue. 100% valid no cap. This AI passed the vibe check. My fit is legendary status now.",
    accuracy: 100,
    dripLevel: 99,
  },
];
</script>

<template>
  <section class="relative py-24 px-6 z-10 overflow-hidden">
    <div class="absolute inset-0 bg-[linear-gradient(to_right,#1f1f1f_1px,transparent_1px),linear-gradient(to_bottom,#1f1f1f_1px,transparent_1px)] bg-[size:4rem_4rem] opacity-20 pointer-events-none"></div>

    <div class="max-w-7xl mx-auto relative">
      <div class="mb-16">
        <div class="flex items-center gap-3 mb-4">
          <span class="relative flex h-3 w-3">
            <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-skena opacity-75"></span>
            <span class="relative inline-flex rounded-full h-3 w-3 bg-skena"></span>
          </span>
          <p class="font-mono text-skena text-sm tracking-widest border border-skena/30 px-2 py-1 bg-black/50">[ SYSTEM_LOG: USER_CONSENSUS ]</p>
        </div>

        <h2 class="text-5xl md:text-7xl font-black text-white leading-tight uppercase">
          THEY GOT <span class="bg-skena text-black px-2 transform -skew-x-6 inline-block">ROASTED.</span> <br />
          THEY LOVED IT.
        </h2>

        <div class="mt-6 border-l-2 border-zinc-800 pl-6">
          <p class="text-zinc-400 font-mono text-sm md:text-base max-w-2xl">
            // Who survived the vibe check? <br />
            > Wall of Fame (or Shame) para korban validasi AI.
          </p>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div v-for="survivor in survivors" :key="survivor.id" class="relative bg-zinc-900/40 backdrop-blur-md border border-zinc-800 p-6 rounded-none group hover:border-skena/50 transition-all duration-300 hover:-translate-y-1">
          <div class="flex justify-between items-start mb-6">
            <div class="flex gap-4">
              <div class="relative">
                <img :src="survivor.avatar" alt="User Avatar" class="w-12 h-12 rounded-full object-cover border-2 border-zinc-700 group-hover:border-skena grayscale group-hover:grayscale-0 transition-all" />
                <div class="absolute -bottom-1 -right-1 bg-skena text-black text-[10px] font-bold px-1 rounded flex items-center"><span class="mr-0.5">✓</span> VERIFIED</div>
              </div>

              <div>
                <h4 class="text-white font-bold text-sm tracking-wide">{{ survivor.name }}</h4>
                <p class="text-zinc-500 text-[10px] font-mono mt-0.5">{{ survivor.id }} // {{ survivor.location }}</p>
              </div>
            </div>

            <div class="flex flex-col items-end">
              <div class="bg-zinc-950 border border-skena/30 text-skena px-2 py-1 font-mono text-xs font-bold rounded group-hover:bg-skena group-hover:text-black transition-colors">SCORE: {{ survivor.score }}</div>
            </div>
          </div>

          <div class="mb-8">
            <p class="text-zinc-300 text-sm leading-relaxed italic relative">
              <span class="text-skena text-2xl absolute -top-3 -left-2 opacity-50">"</span>
              {{ survivor.text }}
            </p>
          </div>

          <div class="space-y-3 border-t border-zinc-800 pt-4 group-hover:border-skena/20 transition-colors">
            <div>
              <div class="flex justify-between text-[10px] font-mono text-zinc-500 mb-1">
                <span>ROAST_ACCURACY</span>
                <span class="text-skena">{{ survivor.accuracy }}%</span>
              </div>
              <div class="w-full h-1 bg-zinc-800 rounded-full overflow-hidden">
                <div class="h-full bg-skena" :style="{ width: survivor.accuracy + '%' }"></div>
              </div>
            </div>

            <div>
              <div class="flex justify-between text-[10px] font-mono text-zinc-500 mb-1">
                <span>DRIP_LEVEL</span>
                <span :class="survivor.dripLevel < 50 ? 'text-red-500' : 'text-skena'"> {{ survivor.dripLevel }}% </span>
              </div>
              <div class="w-full h-1 bg-zinc-800 rounded-full overflow-hidden">
                <div class="h-full transition-all duration-1000" :class="survivor.dripLevel < 50 ? 'bg-red-500' : 'bg-green-500'" :style="{ width: survivor.dripLevel + '%' }"></div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="mt-12 text-center">
        <button class="group relative inline-flex items-center justify-center px-8 py-3 overflow-hidden font-mono font-medium tracking-tighter text-white bg-transparent border border-zinc-700 hover:border-skena transition-colors">
          <span class="absolute w-0 h-0 transition-all duration-500 ease-out bg-skena rounded-full group-hover:w-56 group-hover:h-56 opacity-10"></span>
          <span class="relative flex items-center gap-2 group-hover:text-skena transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M21 12a9 9 0 0 0-9-9 9.75 9.75 0 0 0-6.74 2.74L3 8" />
              <path d="M3 3v5h5" />
              <path d="M3 12a9 9 0 0 0 9 9 9.75 9.75 0 0 0 6.74-2.74L21 16" />
              <path d="M16 21h5v-5" />
            </svg>
            LOAD_MORE_VICTIMS
          </span>
        </button>
      </div>
    </div>
  </section>
</template>
