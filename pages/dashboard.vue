<script setup lang="ts">
// Proteksi halaman dashboard (hanya user login yang bisa masuk)
const user = useSupabaseUser();
watchEffect(() => {
  if (!user.value) {
    navigateTo("/login");
  }
});
</script>

<template>
  <div class="min-h-screen bg-black text-white flex items-center justify-center">
    <div class="text-center">
      <h1 class="text-4xl font-bold text-skena mb-4">Welcome to Dashboard 🚀</h1>
      <p class="text-gray-400">Design dashboard user akan kita buat besok ya!</p>

      <button @click="useSupabaseClient().auth.signOut()" class="mt-8 px-6 py-2 bg-red-600 rounded-full text-white font-bold">Logout</button>
    </div>
  </div>
</template>
