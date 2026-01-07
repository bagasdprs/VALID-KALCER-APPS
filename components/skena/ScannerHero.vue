<script setup lang="ts">
// 1. Variabel Text yang akan muncul di layar
const terminalText = ref("");

// 2. Daftar kalimat yang bakal gantian muncul
const phrases = ["Analyzing aesthetic...", "Detecting basic outfit...", "Calculating cringe levels...", "Scanning skena vibes...", "Loading toxic comments..."];

// 3. State untuk logika animasi
let phraseIndex = 0;
let charIndex = 0;
let isDeleting = false;
let typingSpeed = 100;

onMounted(() => {
  typeWriter();
});

// 4. Fungsi Utama Animasi
const typeWriter = () => {
  const currentPhrase = phrases[phraseIndex] || "";

  if (isDeleting) {
    terminalText.value = currentPhrase.substring(0, charIndex - 1);
    charIndex--;
    typingSpeed = 50;
  } else {
    terminalText.value = currentPhrase.substring(0, charIndex + 1);
    charIndex++;
    typingSpeed = 100;
  }

  if (!isDeleting && charIndex === currentPhrase.length) {
    isDeleting = true;
    typingSpeed = 2000;
  } else if (isDeleting && charIndex === 0) {
    isDeleting = false;
    phraseIndex++;
    typingSpeed = 500; // Istirahat bentar sebelum ngetik lagi

    if (phraseIndex === phrases.length) {
      phraseIndex = 0;
    }
  }

  // Jalanin fungsi ini lagi sesuai kecepatan yang ditentukan
  setTimeout(typeWriter, typingSpeed);
};
</script>

<template>
  <div class="relative min-h-screen flex flex-col justify-center pt-20 overflow-hidden">
    <div class="absolute inset-0 bg-[url('https://grainy-gradients.vercel.app/noise.svg')] opacity-20 pointer-events-none"></div>

    <div class="max-w-7xl mx-auto px-6 grid grid-cols-1 lg:grid-cols-2 gap-12 items-center relative z-10">
      <div class="space-y-8">
        <div class="inline-flex items-center gap-2 px-3 py-1 rounded border border-skena/30 bg-skena/5 text-skena text-xs font-mono mb-4">
          <span class="relative flex h-2 w-2">
            <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-skena opacity-75"></span>
            <span class="relative inline-flex rounded-full h-2 w-2 bg-skena"></span>
          </span>
          SYSTEM ONLINE v.2.0
        </div>

        <h1 class="text-6xl md:text-8xl font-black tracking-tighter leading-[0.9] text-white">
          CEK <br />
          SEBERAPA <br />
          <span class="text-transparent bg-clip-text bg-gradient-to-r from-skena to-white">SKENA</span> <br />
          OUTFIT LO.
        </h1>

        <div class="border-l-2 border-skena/50 pl-6 py-2 bg-gradient-to-r from-skena/5 to-transparent">
          <p class="text-zinc-400 font-mono text-sm md:text-base">
            // AI Judge yang mulutnya pedes. Gratis, akurat, no baper. <br />
            > Upload foto lo. Terima nasib lo.
          </p>
        </div>

        <div class="font-mono text-xs text-skena h-6">> {{ terminalText }}<span class="animate-pulse">_</span></div>

        <div class="flex flex-wrap gap-4 pt-4">
          <button class="bg-skena text-black px-8 py-4 rounded-none skew-x-[-10deg] font-black text-lg hover:bg-white transition-colors shadow-[5px_5px_0px_rgba(255,255,255,0.2)]">
            <div class="skew-x-[10deg] flex items-center gap-2">SCAN OUTFIT GUE 📸</div>
          </button>

          <button class="border border-white/20 text-white px-8 py-4 rounded-none skew-x-[-10deg] font-bold text-sm hover:bg-white/5 transition-colors">
            <div class="skew-x-[10deg]">VIEW COMMANDS</div>
          </button>
        </div>

        <div class="w-full max-w-md border-y border-skena/20 bg-black/30 backdrop-blur-sm py-2 overflow-hidden mt-4">
          <div class="flex animate-marquee whitespace-nowrap font-mono text-xs text-skena tracking-widest">
            <span class="mx-4">/// JANGAN BAPER</span>
            <span class="mx-4">/// VALID KALCER</span>
            <span class="mx-4">/// MENTAL BAJA ONLY</span>
            <span class="mx-4">/// ⚠️ AI TOXIC MODE</span>
            <span class="mx-4">/// JANGAN BAPER</span>
            <span class="mx-4">/// VALID KALCER</span>
            <span class="mx-4">/// MENTAL BAJA ONLY</span>
            <span class="mx-4">/// ⚠️ AI TOXIC MODE</span>
          </div>
        </div>
      </div>

      <div class="relative hidden lg:block group">
        <div class="absolute inset-0 bg-skena/30 blur-[80px] rounded-full scale-75 opacity-60 group-hover:opacity-100 transition-opacity duration-700"></div>

        <div class="relative z-10">
          <div class="absolute inset-0 z-20 bg-[linear-gradient(to_right,#39ff14_1px,transparent_1px),linear-gradient(to_bottom,#39ff14_1px,transparent_1px)] bg-[size:2rem_2rem] opacity-10 mix-blend-overlay pointer-events-none"></div>

          <img
            src="https://images.unsplash.com/photo-1552374196-1ab2a1c593e8?q=80&w=1000&auto=format&fit=crop"
            class="w-full h-[600px] object-cover rounded-2xl grayscale hover:grayscale-0 transition-all duration-500 mask-image-gradient"
            alt="Skena Outfit"
          />

          <div class="absolute top-0 left-0 w-full h-1 bg-skena shadow-[0_0_30px_#39FF14] animate-scan z-30"></div>

          <div class="absolute -right-4 top-20 flex flex-col gap-2 z-30">
            <div class="w-2 h-2 bg-skena rounded-full animate-ping"></div>
            <div class="w-2 h-2 bg-skena rounded-full"></div>
            <div class="w-2 h-2 bg-skena/50 rounded-full"></div>
          </div>

          <div class="absolute bottom-10 -left-6 bg-black/80 backdrop-blur border border-skena/50 p-3 rounded text-skena font-mono text-xs shadow-[0_0_20px_rgba(57,255,20,0.2)]">
            <p>TARGET_LOCKED: <span class="text-white">USER_01</span></p>
            <p>AESTHETIC: <span class="text-red-500">CALCULATING...</span></p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
@keyframes scan {
  0% {
    top: 0%;
    opacity: 0;
  }
  10% {
    opacity: 1;
  }
  90% {
    opacity: 1;
  }
  100% {
    top: 100%;
    opacity: 0;
  }
}
.animate-scan {
  animation: scan 3s cubic-bezier(0.4, 0, 0.2, 1) infinite;
}

@keyframes marquee {
  0% {
    transform: translateX(0);
  }
  100% {
    transform: translateX(-50%);
  }
}
.animate-marquee {
  animation: marquee 15s linear infinite;
}

/* Masking bawah gambar biar agak fading ke hitam */
.mask-image-gradient {
  mask-image: linear-gradient(to bottom, black 80%, transparent 100%);
  -webkit-mask-image: linear-gradient(to bottom, black 80%, transparent 100%);
}
</style>
