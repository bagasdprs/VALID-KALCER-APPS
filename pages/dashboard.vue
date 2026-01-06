<script setup lang="ts">
const user = useSupabaseUser();
watchEffect(() => {
  if (!user.value) {
    navigateTo("/login");
  }
});
</script>

<template>
  <div class="min-h-screen bg-[#050505] text-white font-sans selection:bg-skena selection:text-black">
    <Navbar />

    <main class="pt-24 px-4 pb-10 max-w-7xl mx-auto">
      <div class="text-center py-20">
        <h1 class="text-3xl font-bold text-gray-500">Area Konten Dashboard...</h1>
        <p class="text-gray-600">Nanti Kotak Scan & Result Card masuk sini.</p>
      </div>
    </main>
  </div>
</template>
