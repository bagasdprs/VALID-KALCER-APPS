<script setup lang="ts">
// definePageMeta({
//   // layout: "empty",
//   layout: "default",
// });

const supabase = useSupabaseClient();
const user = useSupabaseUser();

// State form
const email = ref("");
const password = ref("");
const isLoading = ref(false);
const errorMessage = ref("");

// Redirect kalau udah login
watchEffect(() => {
  if (user.value) {
    navigateTo("/dashboard");
  }
});

// Fungsi Login
const handleLogin = async () => {
  isLoading.value = true;
  errorMessage.value = "";

  try {
    const { error } = await supabase.auth.signInWithPassword({
      email: email.value,
      password: password.value,
    });
    if (error) throw error;
  } catch (error: any) {
    errorMessage.value = error.message;
  } finally {
    isLoading.value = false;
  }
};

// Fungsi Social Login (Google/TikTok)
const handleSocialLogin = async (provider: "google" | "tiktok") => {
  // Logic nanti, pasang UI dulu
  alert(`Login with ${provider} coming soon!`);
};
</script>

<template>
  <div class="min-h-screen flex items-center justify-center p-4 relative overflow-hidden font-sans">
    <BackgroundSkena />

    <div class="absolute top-1/2 right-[10%] text-skena text-6xl font-bold opacity-80 animate-float-delayed">+</div>
    <div class="absolute bottom-[20%] left-[15%] text-skena-dim text-6xl opacity-50 animate-spin-slow">*</div>

    <div class="w-full max-w-112.5 relative z-10">
      <div class="absolute -top-3 -right-2 bg-skena text-black text-[10px] font-black px-3 py-1 rounded-sm transform rotate-3 z-20 shadow-neon border border-black uppercase tracking-wider">Gen Z Ready</div>

      <div class="bg-[#0a0a0a]/90 backdrop-blur-md border-2 border-skena/50 rounded-4xl p-8 shadow-neon relative">
        <div class="text-center mb-8">
          <div class="w-16 h-16 bg-zinc-900 border border-zinc-800 rounded-full mx-auto flex items-center justify-center mb-4 shadow-inner group hover:scale-110 transition-transform">
            <UIcon name="i-heroicons-fire-solid" class="w-8 h-8 text-skena group-hover:animate-pulse" />
          </div>
          <h1 class="text-3xl font-bold text-white mb-1 tracking-tight">VibeZone</h1>
          <p class="text-zinc-500 text-sm font-medium">Unlock your AI fashion era.</p>
        </div>

        <div class="space-y-3 mb-6">
          <button
            @click="handleSocialLogin('google')"
            class="w-full bg-zinc-900 hover:bg-zinc-800 hover:border-skena text-white text-sm font-semibold py-3 px-4 rounded-xl flex items-center justify-center gap-3 border border-zinc-800 transition-all group"
          >
            <UIcon name="i-logos-google-icon" class="w-5 h-5" />
            Continue with Google
          </button>

          <button
            @click="handleSocialLogin('tiktok')"
            class="w-full bg-zinc-900 hover:bg-zinc-800 hover:border-[#ff0050] text-white text-sm font-semibold py-3 px-4 rounded-xl flex items-center justify-center gap-3 border border-zinc-800 transition-all"
          >
            <UIcon name="i-logos-tiktok-icon" class="w-5 h-5" />
            Continue with TikTok
          </button>
        </div>

        <div class="flex items-center gap-4 mb-6">
          <div class="h-px bg-zinc-800 flex-1"></div>
          <span class="text-xs text-zinc-600 font-bold uppercase tracking-widest">OR</span>
          <div class="h-px bg-zinc-800 flex-1"></div>
        </div>

        <form @submit.prevent="handleLogin" class="space-y-5">
          <div class="group">
            <label class="block text-[10px] font-bold text-skena tracking-wider uppercase mb-1 ml-1">Email</label>
            <input
              v-model="email"
              type="email"
              placeholder="ur-email@vibe.com"
              class="w-full bg-[#121212] border border-zinc-800 text-white rounded-xl px-4 py-3 text-sm focus:border-skena focus:ring-1 focus:ring-skena focus:shadow-neon outline-none transition-all placeholder:text-zinc-700"
              required
            />
          </div>

          <div class="group">
            <label class="block text-[10px] font-bold text-skena-dim tracking-wider uppercase mb-1 ml-1">Password</label>
            <input
              v-model="password"
              type="password"
              placeholder="********"
              class="w-full bg-[#121212] border border-zinc-800 text-white rounded-xl px-4 py-3 text-sm focus:border-skena focus:ring-1 focus:ring-skena focus:shadow-neon outline-none transition-all placeholder:text-zinc-700"
              required
            />
          </div>

          <div v-if="errorMessage" class="text-red-500 text-xs text-center bg-red-500/10 p-2 rounded-lg border border-red-500/20 animate-shake">
            {{ errorMessage }}
          </div>

          <button
            type="submit"
            :disabled="isLoading"
            class="w-full bg-skena text-black font-black text-lg py-3 rounded-full mt-2 shadow-neon hover:scale-[1.02] hover:bg-white active:scale-95 transition-all flex items-center justify-center gap-2 relative overflow-hidden group"
          >
            <span v-if="isLoading" class="animate-pulse">LOADING...</span>
            <span v-else class="flex items-center gap-2">
              LET'S GO
              <UIcon name="i-heroicons-arrow-right-20-solid" class="w-5 h-5 group-hover:translate-x-1 transition-transform" />
            </span>
          </button>
        </form>

        <div class="text-center mt-6 space-y-2">
          <p class="text-zinc-500 text-xs">
            No account?
            <NuxtLink to="/register" class="text-white underline decoration-skena decoration-2 underline-offset-4 hover:text-skena transition-colors font-bold">Join the wave</NuxtLink>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<!-- <template>
  <div class="min-h-screen bg-black flex items-center justify-center p-4 relative overflow-hidden font-sans">
    <BackgroundSkena />
    <div class="hidden md:block absolute top-[15%] left-[20%] animate-pulse">
      <div class="w-24 h-24 rounded-full overflow-hidden border-4 border-[#db2777] shadow-[0_0_20px_rgba(219,39,119,0.5)]">
        <img src="https://images.unsplash.com/photo-1515378960530-7c0da6231fb1?w=400&q=80" alt="Vibe" class="w-full h-full object-cover" />
      </div>
    </div>

    <div class="hidden md:block absolute bottom-[15%] right-[20%] rotate-6 hover:rotate-0 transition-transform duration-500">
      <div class="w-32 h-32 rounded-2xl overflow-hidden border-b-4 border-r-4 border-primary p-1 bg-zinc-800 shadow-2xl">
        <img src="https://images.unsplash.com/photo-1552374196-c4e7ffc6e126?w=400&q=80" alt="Aesthetic" class="w-full h-full object-cover rounded-xl" />
      </div>
    </div>

    <div class="absolute top-1/2 right-[10%] text-[#db2777] text-6xl font-bold opacity-80">+</div>
    <div class="absolute bottom-[20%] left-[15%] text-primary text-6xl opacity-50 animate-spin-slow">*</div>

    <div class="w-full max-w-105 relative z-10">
      <div class="absolute -top-3 -right-2 bg-[#db2777] text-black text-[10px] font-black px-3 py-1 rounded-sm transform rotate-3 z-20 shadow-lg border border-black uppercase tracking-wider">Gen Z Ready</div>

      <div class="bg-[#0a0a0a] border-2 border-primary rounded-4xl p-8 shadow-[0_0_40px_rgba(34,197,94,0.15)] relative">
        <div class="text-center mb-8">
          <div class="w-12 h-12 bg-zinc-900 border border-zinc-800 rounded-full mx-auto flex items-center justify-center mb-4 shadow-inner">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 text-primary" viewBox="0 0 24 24" fill="currentColor" stroke="none">
              <path
                d="M8.5,14.5A4.49,4.49,0,0,0,13,19h.5a4.49,4.49,0,0,0,3.95-2.58A4.45,4.45,0,0,1,16,17a2.9,2.9,0,0,1-2.9,2.9A2.9,2.9,0,0,1,10.2,17a2.92,2.92,0,0,1,2.53-2.9H13a.58.58,0,0,0,.58-.58.58.58,0,0,0-.58-.58H12a2.4,2.4,0,0,1-2.4-2.4A2.4,2.4,0,0,1,12,8.12a.58.58,0,0,0,.58-.58.58.58,0,0,0-.58-.58H11.5A4.49,4.49,0,0,0,7,11.5a4.49,4.49,0,0,0,4.5,4.5h.5A4.49,4.49,0,0,0,16.45,13.42A4.45,4.45,0,0,1,15,14a2.9,2.9,0,0,1-2.9-2.9A2.9,2.9,0,0,1,15,8.2a2.92,2.92,0,0,1-2.53,2.9H12a.58.58,0,0,0-.58.58.58.58,0,0,0,.58.58h.5a2.4,2.4,0,0,1,2.4,2.4A2.4,2.4,0,0,1,12.5,17.08a.58.58,0,0,0-.58.58.58.58,0,0,0,.58.58h.5A4.49,4.49,0,0,0,17.5,13.5a4.49,4.49,0,0,0-4.5-4.5Z"
              />
              <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8z" opacity=".3" />
            </svg>
          </div>
          <h1 class="text-3xl font-bold text-white mb-1 tracking-tight">Masuk ke VibeZone</h1>
          <p class="text-zinc-500 text-sm font-medium">Unlock your AI fashion era.</p>
        </div>

        <div class="space-y-3 mb-6">
          <button @click="handleSocialLogin('google')" class="w-full bg-zinc-900 hover:bg-zinc-800 text-white text-sm font-semibold py-3 px-4 rounded-xl flex items-center justify-center gap-3 border border-zinc-800 transition-all group">
            <img src="https://www.svgrepo.com/show/475656/google-color.svg" class="w-5 h-5" alt="Google" />
            Continue with Google
          </button>

          <button @click="handleSocialLogin('tiktok')" class="w-full bg-zinc-900 hover:bg-zinc-800 text-white text-sm font-semibold py-3 px-4 rounded-xl flex items-center justify-center gap-3 border border-zinc-800 transition-all">
            <svg class="w-5 h-5 fill-white" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path
                d="M19.589 6.686a4.793 4.793 0 0 1-3.77-4.245V2h-3.445v13.672a2.896 2.896 0 0 1-5.201 1.743l-.002-.001.002.001a2.895 2.895 0 0 1 3.183-4.51v-3.5a6.373 6.373 0 0 0-5.394 9.365 6.373 6.373 0 0 0 8.982-3.28V7.649a8.337 8.337 0 0 0 5.644 1.242v-3.48a4.802 4.802 0 0 1-3.77-4.245V2z"
              />
            </svg>
            Continue with TikTok
          </button>
        </div>

        <div class="flex items-center gap-4 mb-6">
          <div class="h-px bg-zinc-800 flex-1"></div>
          <span class="text-xs text-zinc-600 font-bold uppercase tracking-widest">OR</span>
          <div class="h-px bg-zinc-800 flex-1"></div>
        </div>

        <form @submit.prevent="handleLogin" class="space-y-5">
          <div class="group">
            <label class="block text-[10px] font-bold text-primary tracking-wider uppercase mb-1 ml-1 group-focus-within:text-[#db2777] transition-colors">Email</label>
            <input
              v-model="email"
              type="email"
              placeholder="ur-email@vibe.com"
              class="w-full bg-[#121212] border border-zinc-800 text-white rounded-xl px-4 py-3 text-sm focus:border-primary focus:ring-1 focus:ring-primary outline-none transition-all placeholder:text-zinc-700"
              required
            />
          </div>

          <div class="group">
            <label class="block text-[10px] font-bold text-[#db2777] tracking-wider uppercase mb-1 ml-1 group-focus-within:text-primary transition-colors">Password</label>
            <input
              v-model="password"
              type="password"
              placeholder="********"
              class="w-full bg-[#121212] border border-zinc-800 text-white rounded-xl px-4 py-3 text-sm focus:border-[#db2777] focus:ring-1 focus:ring-[#db2777] outline-none transition-all placeholder:text-zinc-700"
              required
            />
          </div>

          <div v-if="errorMessage" class="text-red-500 text-xs text-center bg-red-500/10 p-2 rounded-lg border border-red-500/20">
            {{ errorMessage }}
          </div>

          <button
            type="submit"
            :disabled="isLoading"
            class="w-full bg-primary text-black font-black text-lg py-3 rounded-full mt-2 hover:scale-[1.02] hover:shadow-[0_0_25px_rgba(34,197,94,0.6)] active:scale-95 transition-all flex items-center justify-center gap-2 relative overflow-hidden group"
          >
            <span v-if="isLoading">LOADING...</span>
            <span v-else class="flex items-center gap-2">
              LET'S GO
              <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5 group-hover:translate-x-1 transition-transform" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="3">
                <path stroke-linecap="round" stroke-linejoin="round" d="M14 5l7 7m0 0l-7 7m7-7H3" />
              </svg>
            </span>
          </button>
        </form>

        <div class="text-center mt-6 space-y-2">
          <p class="text-zinc-500 text-xs">
            No account?
            <NuxtLink to="/register" class="text-white underline decoration-primary decoration-2 underline-offset-4 hover:text-primary transition-colors font-bold">Join the wave</NuxtLink>
          </p>
          <button class="text-zinc-600 text-[10px] hover:text-zinc-400 transition-colors">Forgot password?</button>
        </div>
      </div>
    </div>
  </div>
</template> -->

<style scoped>
/* Animasi Muter pelan buat dekorasi */
@keyframes spin-slow {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
.animate-spin-slow {
  animation: spin-slow 12s linear infinite;
}

.animate-shake {
  animation: shake 0.5s cubic-bezier(0.36, 0.07, 0.19, 0.97) both;
}

@keyframes shake {
  10%,
  90% {
    transform: translate3d(-1px, 0, 0);
  }
  20%,
  80% {
    transform: translate3d(2px, 0, 0);
  }
  30%,
  50%,
  70% {
    transform: translate3d(-4px, 0, 0);
  }
  40%,
  60% {
    transform: translate3d(4px, 0, 0);
  }
}
</style>
