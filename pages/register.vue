<script setup lang="ts">
const supabase = useSupabaseClient();
const user = useSupabaseUser();

// State Form
const username = ref("");
const email = ref("");
const password = ref("");
const aesthetic = ref(""); // Buat dropdown
const isLoading = ref(false);
const errorMessage = ref("");

// Pilihan Aesthetic (Buat dropdown)
const aestheticOptions = ["Streetwear 👟", "Y2K ✨", "Cyberpunk 🤖", "Old Money 💸", "Gorpcore ⛰️", "Vintage 📼"];

// Redirect kalau user ternyata udah login
watchEffect(() => {
  if (user.value) {
    navigateTo("/dashboard");
  }
});

// Fungsi Register
const handleRegister = async () => {
  if (!aesthetic.value) {
    errorMessage.value = "Please select your aesthetic vibe! 🤨";
    return;
  }

  isLoading.value = true;
  errorMessage.value = "";

  try {
    // 1. Daftar ke Supabase Auth
    const { error } = await supabase.auth.signUp({
      email: email.value,
      password: password.value,
      options: {
        // Kita simpan Username & Aesthetic di Metadata user
        data: {
          username: username.value,
          aesthetic: aesthetic.value,
        },
      },
    });

    if (error) throw error;

    // Kalau sukses, biasanya Supabase kirim email konfirmasi (tergantung settingan)
    alert("Registrasi Berhasil! Cek email kamu buat konfirmasi 📧");
    navigateTo("/login");
  } catch (error: any) {
    errorMessage.value = error.message;
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <div class="min-h-screen flex items-center justify-center p-4 relative overflow-hidden font-sans">
    <BackgroundSkena />

    <div class="w-full max-w-5xl grid grid-cols-1 lg:grid-cols-2 gap-0 relative z-10">
      <div class="hidden lg:flex relative items-center justify-center p-10">
        <div class="relative w-95 h-125 border-4 border-white rounded-4xl bg-zinc-800 -rotate-3 hover:rotate-0 transition-transform duration-500 shadow-[0_0_30px_rgba(57,255,20,0.3)] overflow-hidden z-20">
          <img src="https://images.unsplash.com/photo-1617387682382-353285750339?q=80&w=1965&auto=format&fit=crop" alt="Cyber Aesthetic" class="w-full h-full object-cover" />
        </div>

        <div class="absolute top-20 right-20 bg-white p-2 pb-8 shadow-lg rotate-12 z-10 animate-float-delayed">
          <div class="w-24 h-24 bg-orange-200 overflow-hidden">
            <img src="https://images.unsplash.com/photo-1511556532299-8f662fc26c06?q=80&w=100" class="w-full h-full object-cover" />
          </div>
        </div>

        <div class="absolute bottom-20 left-10 text-skena animate-spin-slow opacity-80 z-30">
          <UIcon name="i-heroicons-star-solid" class="w-16 h-16" />
        </div>
      </div>

      <div class="flex items-center justify-center">
        <div class="w-full max-w-md bg-foreground/90 backdrop-blur-xl border border-skena rounded-4xl p-8 shadow-[0_0_20px_rgba(57,255,20,0.15)] relative">
          <div class="absolute -top-1 -right-1 w-6 h-6 border-t-2 border-r-2 border-skena rounded-tr-lg"></div>
          <div class="absolute -bottom-1 -left-1 w-6 h-6 border-b-2 border-l-2 border-skena rounded-bl-lg"></div>

          <div class="mb-6">
            <h1 class="text-4xl font-black text-white mb-2">Join the <span class="text-skena">VibeZone</span> ⚡</h1>
            <p class="text-zinc-400 text-sm">Create your AI stylist profile & start dripping.</p>
          </div>

          <form @submit.prevent="handleRegister" class="space-y-4">
            <div class="space-y-1">
              <label class="text-[10px] font-bold text-zinc-500 uppercase tracking-widest ml-1">Nama Panggung (Username)</label>
              <div class="relative group">
                <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                  <span class="text-zinc-500 font-bold group-focus-within:text-skena">@</span>
                </div>
                <input
                  v-model="username"
                  type="text"
                  placeholder="e.g. CyberBae_99"
                  class="w-full bg-[#121212] border border-zinc-700 text-white rounded-xl py-3 pl-10 pr-4 focus:border-skena focus:ring-1 focus:ring-skena outline-none transition-all"
                  required
                />
              </div>
            </div>

            <div class="space-y-1">
              <label class="text-[10px] font-bold text-zinc-500 uppercase tracking-widest ml-1">Email Drop</label>
              <div class="relative group">
                <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                  <UIcon name="i-heroicons-envelope" class="text-zinc-500 group-focus-within:text-skena" />
                </div>
                <input
                  v-model="email"
                  type="email"
                  placeholder="you@universe.com"
                  class="w-full bg-[#121212] border border-zinc-700 text-white rounded-xl py-3 pl-10 pr-4 focus:border-skena focus:ring-1 focus:ring-skena outline-none transition-all"
                  required
                />
              </div>
            </div>

            <div class="space-y-1">
              <label class="text-[10px] font-bold text-zinc-500 uppercase tracking-widest ml-1">Your Main Aesthetic</label>
              <div class="relative">
                <select v-model="aesthetic" class="w-full bg-[#121212] border border-zinc-700 text-white rounded-xl py-3 pl-4 pr-10 focus:border-skena focus:ring-1 focus:ring-skena outline-none appearance-none cursor-pointer">
                  <option value="" disabled>Select your vibe...</option>
                  <option v-for="opt in aestheticOptions" :key="opt" :value="opt">{{ opt }}</option>
                </select>
                <div class="absolute inset-y-0 right-0 pr-4 flex items-center pointer-events-none">
                  <UIcon name="i-heroicons-chevron-down" class="text-zinc-500" />
                </div>
              </div>
            </div>

            <div class="space-y-1">
              <label class="text-[10px] font-bold text-zinc-500 uppercase tracking-widest ml-1">Password</label>
              <div class="relative group">
                <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                  <UIcon name="i-heroicons-lock-closed" class="text-zinc-500 group-focus-within:text-skena" />
                </div>
                <input
                  v-model="password"
                  type="password"
                  placeholder="Make it secret..."
                  class="w-full bg-[#121212] border border-zinc-700 text-white rounded-xl py-3 pl-10 pr-4 focus:border-skena focus:ring-1 focus:ring-skena outline-none transition-all"
                  required
                />
              </div>
            </div>

            <div v-if="errorMessage" class="p-3 bg-red-500/10 border border-red-500/50 rounded-lg text-red-500 text-xs text-center">
              {{ errorMessage }}
            </div>

            <button
              type="submit"
              :disabled="isLoading"
              class="w-full bg-skena hover:bg-[#2bff05] text-black font-black text-lg py-4 rounded-full mt-2 shadow-[0_0_20px_rgba(57,255,20,0.4)] hover:shadow-[0_0_30px_rgba(57,255,20,0.6)] hover:scale-[1.02] active:scale-95 transition-all flex items-center justify-center gap-2 group"
            >
              <span v-if="isLoading">CREATING PROFILE...</span>
              <span v-else class="flex items-center gap-2">
                GENERATE MY PROFILE
                <UIcon name="i-heroicons-rocket-launch-solid" class="group-hover:translate-x-1 group-hover:-translate-y-1 transition-transform" />
              </span>
            </button>
          </form>

          <div class="relative flex py-5 items-center">
            <div class="grow border-t border-zinc-800"></div>
            <span class="shrink mx-4 text-[10px] text-zinc-600 font-bold uppercase tracking-widest">Or Vibe With</span>
            <div class="grow border-t border-zinc-800"></div>
          </div>

          <div class="flex justify-center gap-4 mb-4">
            <button class="w-12 h-12 rounded-full bg-zinc-900 border border-zinc-800 flex items-center justify-center hover:border-skena hover:text-skena transition-all">
              <UIcon name="i-logos-google-icon" class="w-5 h-5" />
            </button>
            <button class="w-12 h-12 rounded-full bg-zinc-900 border border-zinc-800 flex items-center justify-center hover:border-[#ff0050] hover:text-[#ff0050] transition-all">
              <UIcon name="i-logos-tiktok-icon" class="w-5 h-5" />
            </button>
          </div>

          <div class="text-center">
            <p class="text-zinc-500 text-xs">
              Already vibing?
              <NuxtLink to="/login" class="text-white font-bold underline decoration-skena decoration-2 hover:text-skena transition-colors"> Log in here </NuxtLink>
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
