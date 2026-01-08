<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>

  <!-- <div class="fixed bottom-4 right-4 z-[9999] flex flex-col gap-2 opacity-50 hover:opacity-100 transition-opacity">
    <button @click="setRole('admin')" class="bg-red-600 text-white px-3 py-1 text-xs font-bold rounded shadow">👮‍♂️ Jadi Admin</button>
    <button @click="setRole('user')" class="bg-blue-600 text-white px-3 py-1 text-xs font-bold rounded shadow">👤 Jadi User</button>
    <button @click="logout()" class="bg-gray-600 text-white px-3 py-1 text-xs font-bold rounded shadow">🚪 Logout</button>
  </div> -->
</template>

<script setup lang="ts">
const tokenCookie = useCookie("token");
const roleCookie = useCookie("role");

function setRole(role: string) {
  tokenCookie.value = "dummy-token-123";
  roleCookie.value = role;
  alert(`Mode berubah jadi: ${role.toUpperCase()}`);
  window.location.reload();
}

function logout() {
  tokenCookie.value = null;
  roleCookie.value = null;
  alert("Sudah Logout!");
  window.location.reload();
}
</script>
