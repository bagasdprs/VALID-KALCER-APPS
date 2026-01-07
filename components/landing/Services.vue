<script setup lang="ts">
const services = [
  {
    id: "01",
    title: "AI FIT CHECK",
    desc: "Jujurly, outfit lo valid gak? Upload OOTD, biar AI gue yang nge-judge dengan pedas dan akurat. No baper allowed.",
    icon: "scan",
  },
  {
    id: "02",
    title: "VIBE MATCHER",
    desc: "Cari klan lo. Apakah lo anak Skena Jaksel, Jamet Kuproy, atau Cyber-Goth? Kita cocokin aura lo.",
    icon: "target",
  },
  {
    id: "03",
    title: "THRIFT HUNTER",
    desc: "Gatekeep no more. Deteksi barang branded di tumpukan baju bekas. Hemat pangkal rich, Bestie.",
    icon: "tag",
  },
];
</script>

<template>
  <section class="relative py-24 px-6 z-10">
    <div class="max-w-7xl mx-auto">
      <div class="mb-16 border-l-4 border-skena pl-6">
        <div class="flex items-center gap-2 mb-2">
          <span class="w-2 h-2 bg-skena animate-pulse"></span>
          <p class="font-mono text-skena text-xs tracking-[0.2em]">SYSTEM_MODULES // V.2.0.4</p>
        </div>
        <h2 class="text-5xl md:text-7xl font-black text-white tracking-tighter uppercase">VALID <span class="text-transparent bg-clip-text bg-gradient-to-r from-skena to-zinc-500">SERVICES</span></h2>
        <p class="text-zinc-400 mt-4 max-w-2xl font-mono text-sm md:text-base">
          // Tools canggih buat validasi eksistensi lo. <br />
          > Pilih modul di bawah untuk memulai eksekusi.
        </p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div v-for="service in services" :key="service.id" class="group relative bg-zinc-900/20 backdrop-blur-sm p-8 min-h-[300px] flex flex-col justify-between hover:bg-skena/5 transition-colors duration-500 cursor-none sm:cursor-auto">
          <div class="absolute top-0 left-0 w-6 h-6 border-t-2 border-l-2 border-zinc-700 group-hover:border-skena transition-colors duration-300"></div>

          <div class="absolute top-0 right-0 w-6 h-6 border-t-2 border-r-2 border-zinc-700 group-hover:border-skena transition-colors duration-300"></div>

          <div class="absolute bottom-0 left-0 w-6 h-6 border-b-2 border-l-2 border-zinc-700 group-hover:border-skena transition-colors duration-300"></div>

          <div class="absolute bottom-0 right-0 w-6 h-6 border-b-2 border-r-2 border-zinc-700 group-hover:border-skena transition-colors duration-300"></div>

          <div>
            <div class="flex justify-between items-start mb-6">
              <div class="w-12 h-12 border border-zinc-700 bg-black/50 flex items-center justify-center group-hover:border-skena group-hover:bg-skena group-hover:text-black transition-all duration-300">
                <svg v-if="service.icon === 'scan'" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M3 7V5a2 2 0 0 1 2-2h2" />
                  <path d="M17 3h2a2 2 0 0 1 2 2v2" />
                  <path d="M21 17v2a2 2 0 0 1-2 2h-2" />
                  <path d="M7 21H5a2 2 0 0 1-2-2v-2" />
                  <rect width="10" height="10" x="7" y="7" rx="2" />
                </svg>

                <svg v-if="service.icon === 'target'" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="12" cy="12" r="10" />
                  <circle cx="12" cy="12" r="6" />
                  <circle cx="12" cy="12" r="2" />
                </svg>

                <svg v-if="service.icon === 'tag'" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M12 2H2v10l9.29 9.29c.94.94 2.48.94 3.42 0l6.58-6.58c.94-.94.94-2.48 0-3.42L12 2Z" />
                  <path d="M7 7h.01" />
                </svg>
              </div>

              <span class="font-mono text-zinc-600 text-xs group-hover:text-skena transition-colors"> [{{ service.id }}] </span>
            </div>

            <h3 class="text-2xl font-black text-white mb-3 uppercase tracking-tight group-hover:text-skena transition-colors duration-300">
              {{ service.title }}
            </h3>
            <p class="text-zinc-400 text-sm leading-relaxed border-l border-zinc-800 pl-3 group-hover:border-skena/50 transition-colors duration-300">
              {{ service.desc }}
            </p>
          </div>

          <div class="mt-8">
            <div class="w-full h-[1px] bg-zinc-800 relative overflow-hidden">
              <div class="absolute inset-0 bg-skena w-0 group-hover:w-full transition-all duration-700 ease-out"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
