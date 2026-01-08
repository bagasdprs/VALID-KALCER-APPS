<script setup lang="ts">
import { ref } from "vue";

const openIndex = ref<number | null>(null);

const faqs = [
  {
    question: "AI-NYA ERROR GAK SIH? MASA OUTFIT MAHAL GUE DIBILANG JELEK?",
    answer: "Harga gak beli selera, Bestie. Algoritma kita dilatih pake visual composition & color theory, bukan baca struk belanja. Kalau outfit lo mahal tapi nabrak, ya tetep jelek di mata sistem. Cry about it.",
  },
  {
    question: "FOTO GUE AMAN GAK? TAKUT DISEBAR NIH.",
    answer: "Aman, chill. Kita gak butuh nyimpen arsip outfit basic lo, menuh-menuhin server doang. Habis di-roasting & dapet skor, data foto lo langsung di-shredder (dihapus dari memori sementara). Privasi lo valid.",
  },
  {
    question: "GIMANA CARA BIAR DAPET SKOR 10/10?",
    answer: "Ganti gaya. Atau coba lagi tahun depan pas selera lo udah upgrade. Hint: AI kita suka siluet yang unik, layering yang rapi, dan color palette yang gak bikin sakit mata. Usaha dikit lah.",
  },
  {
    question: "INI BENERAN AI ATAU ADMIN GABUT YANG NGETIK?",
    answer: "100% Artificial Intelligence, 0% Empathy. Kita pake teknologi LLM & Vision tercanggih. Gak ada admin yang punya waktu buat nge-roast lo satu-satu manual.",
  },
  {
    question: "KOK MULUT AI-NYA JAHAT BANGET? SIAPA YANG NGAJARIN?",
    answer: "Itu bukan jahat, itu jujur. Dunia udah kebanyakan sugarcoating. Kita di sini buat ngasih reality check biar lo gak keluar rumah kayak badut. Anggap aja ini tough love.",
  },
  {
    question: "INI BAYAR GAK SIH?",
    answer: "Gratis. Tis. Tis. Tapi ada *Rate Limit*. Jangan spam tombol scan mentang-mentang gratis. API AI mahal, Bestie. Kalau lo spam, IP lo bakal kena *timeout* otomatis.",
  },
  {
    question: "APA AJA YANG DINILAI SAMA AI-NYA?",
    answer: "4 Pilar Utama: Silhouette (Bentuk baju di badan lo), Color Coordination (Nabrak atau senada), Vibe Context (Nyambung gak sama lokasi), dan Originality. Brand mahal gak masuk hitungan.",
  },
  {
    question: "ADA APLIKASINYA DI PLAYSTORE/APPSTORE?",
    answer: "Gak perlu install-install club. Web ini udah PWA (*Progressive Web App*). Tinggal 'Add to Home Screen', rasanya sama kayak app native tapi gak bikin memori HP kentang lo penuh.",
  },
  {
    question: "HASILNYA BISA DI-REVISI GAK? GUE GAK TERIMA.",
    answer: "Keputusan AI bersifat mutlak dan tidak dapat diganggu gugat. Kalau lo gak terima dibilang *mid*, berarti emang lo *mid*. Coba foto dengan *angle* atau *lighting* yang lebih proper, siapa tau nambah 0.5 poin.",
  },
  {
    question: "SPILL TECH STACK DONG BANG?",
    answer: "Frontend: Nuxt 3 (Vue). Backend: Golang (Fiber). Database: Supabase. AI: Gemini Pro Vision. Stack kita lebih valid dan *performant* daripada OOTD lo.",
  },
  {
    question: "HARUS BAYAR GAK? DUIT GUE ABIS BUAT THRIFT.",
    answer:
      "Tenang, ada 3 jalur buat lo:\n1. TIER NPC (Free): Gratis seumur hidup. Dapet skor & roastingan kasar doang. Iklan numpuk.\n2. TIER MAIN CHARACTER (29k/bln): No ads, dapet 'Patch Notes' (saran perbaikan outfit), & akses leaderboard.\n3. TIER FINAL BOSS (99k/bln): Unlimited scans, deep analytic (color hex, brand detect), & badge 'Verified Skena' di profil.",
  },
  {
    question: "GUE UDAH BAYAR TAPI SAKIT HATI DI-ROAST, BISA REFUND?",
    answer: "Kocak lo. Ini 'Valid Kalcer', bukan sesi terapi. Kita jual kebenaran pahit. No Refund policy berlaku mutlak. Kalau gak kuat mental, stick to Free Tier aja atau balik ke IG story biar dipuji temen fake lo.",
  },
  {
    question: "HP GUE KENTANG KAMERANYA BUREM, NGARUH GAK?",
    answer: "Ngaruh dikit. Kalau resolusi lo setara CCTV Indomaret, AI bakal susah detect tekstur kain lo. Minimal pencahayaan terang lah, biar outfit lo yang dinilai, bukan noise kamera lo.",
  },
  {
    question: "DATANYA DARI MANA? SOTOY BANGET AI-NYA.",
    answer: "Kita scraping data dari ribuan foto street style Paris, Tokyo, sampai Blok M. Dataset kita divalidasi sama fashion expert beneran sebelum diajatin ke AI. Jadi kalau dia bilang celana lo gak nyambung, trust the data.",
  },
  {
    question: "GIMANA CARA MASUK GLOBAL LEADERBOARD?",
    answer: "Konsistensi. Upload fit beda-beda tiap hari. Kalau rata-rata skor lo di atas 8.5 selama seminggu, lo bakal otomatis naik ke 'Hall of Validation'. Ingat: curang pake foto orang lain = auto ban permanen.",
  },
  {
    question: "BRAND GUE MAU COLLAB ATAU PASANG IKLAN, BISA?",
    answer: "Bisa, asal brand lo gak cringe. Kita kurasi ketat partner yang masuk ekosistem Valid Kalcer. Kalau lo yakin produk lo valid, drop email di footer. Kalau 'mid', jangan harap dibales.",
  },

  // {
  //   question: "GUE PAKE BRANDED HEAD-TO-TOE, KOK SKOR CUMA 3?",
  //   answer: "Logo segede gaban gak bikin lo otomatis keren. Money can't buy taste. Kalau styling lo kayak papan iklan berjalan atau jemuran numpuk, ya tetep dapet nilai jeblok. Less logo, more vibe.",
  // },
  // {
  //   question: "BISA PROTES GAK? GUE GAK TERIMA DIKATAIN JAMET!",
  //   answer: "Tombol komplain udah kita hapus demi kesehatan mental admin. Keputusan AI bersifat final, mengikat, dan menyakitkan. Terima nasib atau belajar styling lagi.",
  // },
  // {
  //   question: "APA GUNANYA APLIKASI INI SELAIN BIKIN SAKIT HATI?",
  //   answer: "Menyelamatkan kesehatan mata publik. Kita adalah garda terdepan melawan polusi visual di jalanan. Kalau lo tersinggung, berarti lo bagian dari polusinya. You're welcome.",
  // },
  // {
  //   question: "BISA BAYAR BIAR DAPET NILAI BAGUS? (PREMIUM USER)",
  //   answer: "Gak. Simpen duit lo buat beli wardrobe yang beneran valid, atau minimal beli kaca baru. Integritas algoritma kita lebih mahal daripada sogokan lo.",
  // },
  // {
  //   question: "ADA VERSI APLIKASI HP GAK? BIAR GAMPANG AKSES.",
  //   answer: "Gak usah nambah-nambahin sampah di storage HP lo yang udah penuh sama foto selfie burik itu. Pake web aja udah cukup canggih. Gak perlu install, langsung gas.",
  // },
];

const toggle = (index: number) => {
  openIndex.value = openIndex.value === index ? null : index;
};
</script>

<template>
  <section class="relative py-24 px-4 md:px-6 bg-black z-10 font-sans border-t border-zinc-900/50">
    <div class="max-w-4xl mx-auto">
      <div class="text-center mb-16 relative">
        <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full border border-green-900 bg-green-900/10 mb-6">
          <span class="w-2 h-2 rounded-full bg-skena animate-pulse"></span>
          <span class="text-[10px] font-mono text-skena tracking-widest uppercase">SYSTEM STATUS: ONLINE</span>
        </div>

        <h2 class="text-5xl md:text-7xl font-black text-white tracking-tighter mb-4">FAQ // <span class="text-transparent bg-clip-text bg-gradient-to-r from-skena to-green-600">BANYAK TANYA LO</span></h2>

        <div class="text-zinc-500 font-mono text-xs md:text-sm space-y-1">
          <p>> Initiating validation protocol...</p>
          <p>> Decoding common user queries regarding style metrics.</p>
        </div>
      </div>

      <div class="space-y-4">
        <div v-for="(item, index) in faqs" :key="index" class="group border border-zinc-800 bg-zinc-900/20 hover:border-skena/30 transition-all duration-300">
          <button @click="toggle(index)" class="w-full flex items-center justify-between p-6 text-left">
            <div class="flex items-center gap-6">
              <span class="font-mono text-skena/50 text-sm md:text-base"> {{ String(index + 1).padStart(2, "0") }} // </span>
              <h3 class="text-white font-bold text-lg md:text-xl uppercase tracking-tight group-hover:text-skena transition-colors">
                {{ item.question }}
              </h3>
            </div>

            <div class="w-8 h-8 md:w-10 md:h-10 border border-zinc-700 flex items-center justify-center group-hover:bg-skena group-hover:text-black group-hover:border-skena transition-all duration-300">
              <span class="text-xl leading-none transition-transform duration-500" :class="{ 'rotate-45': openIndex === index }"> + </span>
            </div>
          </button>

          <div class="grid transition-all duration-500 ease-in-out" :class="openIndex === index ? 'grid-rows-[1fr] opacity-100' : 'grid-rows-[0fr] opacity-0'">
            <div class="overflow-hidden">
              <div class="px-6 pb-8 pl-16 md:pl-20">
                <p class="text-zinc-400 font-mono text-sm md:text-base leading-relaxed border-l-2 border-skena/20 pl-4"><span class="text-skena">>_</span> {{ item.answer }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- <div class="mt-16 flex justify-center opacity-20">
        <span class="text-skena text-2xl animate-spin-slow">⚙</span>
      </div> -->
    </div>
  </section>
</template>

<style scoped>
.animate-spin-slow {
  animation: spin 8s linear infinite;
}
@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>
